package pgxhub

import (
	"context"
	"crypto/tls"
	_ "embed"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"google.golang.org/grpc/status"

	"github.com/vanti-dev/sc-bos/internal/util/pki"
	"github.com/vanti-dev/sc-bos/internal/util/rpcutil"
	"github.com/vanti-dev/sc-bos/pkg/system/hub/enroll"

	"github.com/vanti-dev/sc-bos/pkg/gen"
)

//go:embed schema.sql
var schemaSql string

func SetupDB(ctx context.Context, pool *pgxpool.Pool) error {
	return pool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, schemaSql)
		return err
	})
}

func NewServer(ctx context.Context, connStr string) (*Server, error) {
	pool, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("connect %w", err)
	}

	return NewServerFromPool(ctx, pool)
}

func NewServerFromPool(ctx context.Context, pool *pgxpool.Pool, opts ...Option) (*Server, error) {
	err := SetupDB(ctx, pool)
	if err != nil {
		return nil, fmt.Errorf("setup %w", err)
	}

	s := &Server{
		pool: pool,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s, nil
}

type Server struct {
	gen.UnimplementedHubApiServer
	logger *zap.Logger
	pool   *pgxpool.Pool

	ManagerName   string
	ManagerAddr   string
	Authority     pki.Source  // trust authority for the cohort of smart core nodes
	TestTLSConfig *tls.Config // TLS config used when initiating test connections with a node
}

func (n *Server) GetHubNode(ctx context.Context, request *gen.GetHubNodeRequest) (*gen.HubNode, error) {
	logger := rpcutil.ServerLogger(ctx, n.logger)
	var dbEnrollment Enrollment
	err := n.pool.BeginFunc(ctx, func(tx pgx.Tx) (err error) {
		dbEnrollment, err = GetEnrollment(ctx, tx, request.GetAddress())
		return
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, status.Error(codes.NotFound, "no node registration with specified address")
	} else if err != nil {
		logger.Error("GetEnrollment failed", zap.Error(err), zap.String("address", request.GetAddress()))
		return nil, status.Error(codes.Internal, "failed to retrieve enrollment")
	}

	return &gen.HubNode{
		Name:        dbEnrollment.Name,
		Address:     dbEnrollment.Address,
		Description: dbEnrollment.Description,
	}, nil
}

func (n *Server) EnrollHubNode(ctx context.Context, request *gen.EnrollHubNodeRequest) (*gen.HubNode, error) {
	logger := rpcutil.ServerLogger(ctx, n.logger)
	nodeReg := request.GetNode()
	if nodeReg == nil {
		return nil, status.Error(codes.InvalidArgument, "node must be supplied")
	}
	if nodeReg.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "node.address must be supplied")
	}

	// check if the node is already enrolled
	err := n.pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		_, err := GetEnrollment(ctx, tx, nodeReg.Address)
		return err
	})
	if !errors.Is(err, pgx.ErrNoRows) {
		return nil, status.Errorf(codes.AlreadyExists, "%s already enrolled", nodeReg.Address)
	}

	en, err := enroll.Controller(ctx, &gen.Enrollment{
		TargetName:     nodeReg.Name,
		TargetAddress:  nodeReg.Address,
		ManagerName:    n.ManagerName,
		ManagerAddress: n.ManagerAddr,
	}, n.Authority)
	if err != nil {
		logger.Error("failed to enroll area controller", zap.Error(err),
			zap.String("target_address", nodeReg.Address))
		return nil, status.Error(codes.Unknown, "target refused registration")
	}

	err = n.pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		return CreateEnrollment(ctx, tx, Enrollment{
			Name:        en.TargetName,
			Description: nodeReg.Description,
			Address:     en.TargetAddress,
			Cert:        en.Certificate,
		})
	})
	if err != nil {
		delErr := n.deleteHubNode(ctx, nodeReg)
		if delErr != nil {
			// failed to rollback!
			logger.Error("pool.CreateEnrollment failed, failed to rollback",
				zap.NamedError("enroll", err), zap.NamedError("rollback", delErr))
			return nil, status.Errorf(codes.DataLoss, "enrollment failed, rollback failed - the system is in a corrupt state, manual intervention required")
		}
		logger.Warn("pool.CreateEnrollment failed", zap.Error(err))
		return nil, status.Error(codes.Aborted, "failed to save the enrollment, no changes have been made")
	}
	return nodeReg, nil
}

func (n *Server) deleteHubNode(ctx context.Context, reg *gen.HubNode) error {
	conn, err := grpc.DialContext(ctx, reg.Address, grpc.WithTransportCredentials(credentials.NewTLS(n.TestTLSConfig)))
	if err != nil {
		n.logger.Error("failed to connect to node", zap.Error(err))
		return status.Error(codes.Unavailable, "unable to connect to target node")
	}
	client := gen.NewEnrollmentApiClient(conn)
	_, err = client.DeleteEnrollment(ctx, &gen.DeleteEnrollmentRequest{})
	if err != nil {
		return err
	}
	return nil
}

func (n *Server) ListHubNodes(ctx context.Context, request *gen.ListHubNodesRequest) (*gen.ListHubNodesResponse, error) {
	logger := rpcutil.ServerLogger(ctx, n.logger)
	var dbEnrollments []Enrollment
	err := n.pool.BeginFunc(ctx, func(tx pgx.Tx) (err error) {
		dbEnrollments, err = ListEnrollments(ctx, tx)
		return
	})
	if err != nil {
		logger.Error("pool.ListEnrollments failed", zap.Error(err))
		return nil, status.Error(codes.Unavailable, "unable to retrieve enrollments")
	}

	var registrations []*gen.HubNode
	for _, en := range dbEnrollments {
		registrations = append(registrations, &gen.HubNode{
			Name:        en.Name,
			Address:     en.Address,
			Description: en.Description,
		})
	}

	return &gen.ListHubNodesResponse{Nodes: registrations}, nil
}

func (n *Server) TestHubNode(ctx context.Context, request *gen.TestHubNodeRequest) (*gen.TestHubNodeResponse, error) {
	reg, err := n.GetHubNode(ctx, &gen.GetHubNodeRequest{Address: request.GetAddress()})
	if err != nil {
		return nil, err
	}

	logger := n.logger.With(zap.String("node_address", reg.Address))

	conn, err := grpc.DialContext(ctx, reg.Address, grpc.WithTransportCredentials(credentials.NewTLS(n.TestTLSConfig)))
	if err != nil {
		logger.Error("failed to connect to node", zap.Error(err))
		return nil, status.Error(codes.Unavailable, "unable to connect to target node")
	}

	client := grpc_reflection_v1alpha.NewServerReflectionClient(conn)
	stream, err := client.ServerReflectionInfo(ctx)
	if err != nil {
		logger.Error("failed to reflect target node", zap.Error(err))
		return nil, status.Error(codes.Unavailable, "target node reflection failed")
	}

	err = stream.Send(&grpc_reflection_v1alpha.ServerReflectionRequest{
		Host:           reg.Address,
		MessageRequest: &grpc_reflection_v1alpha.ServerReflectionRequest_ListServices{ListServices: ""},
	})
	if err != nil {
		logger.Error("node reflection: send request", zap.Error(err))
		return nil, status.Error(codes.Unavailable, "target node reflection failed")
	}

	res, err := stream.Recv()
	if err != nil {
		logger.Error("node reflection: receive response", zap.Error(err))
		return nil, status.Error(codes.Unavailable, "target node reflection failed")
	}

	var serviceNames []string
	for _, service := range res.GetListServicesResponse().GetService() {
		serviceNames = append(serviceNames, service.Name)
	}

	return &gen.TestHubNodeResponse{}, nil
}
