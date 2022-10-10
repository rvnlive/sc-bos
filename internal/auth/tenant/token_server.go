package tenant

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/vanti-dev/bsp-ew/internal/auth"
	"github.com/vanti-dev/bsp-ew/internal/util/rpcutil"
	"go.uber.org/zap"
)

type TokenServer struct {
	tokens *TokenSource
	logger *zap.Logger

	clientCredentialVerifier Verifier
	clientCredentialValidity time.Duration

	passwordVerifier Verifier
	passwordValidity time.Duration
}

func NewTokenServer(name string, opts ...TokenServerOption) (*TokenServer, error) {
	key, err := generateKey()
	if err != nil {
		return nil, err
	}
	tokens := &TokenSource{
		Key:    key,
		Issuer: name,
		Now:    time.Now,
	}

	s := &TokenServer{tokens: tokens}
	for _, opt := range opts {
		opt(s)
	}

	return s, nil
}

type TokenServerOption func(ts *TokenServer)

func WithLogger(logger *zap.Logger) TokenServerOption {
	return func(ts *TokenServer) {
		ts.logger = logger
	}
}

func WithClientCredentialFlow(v Verifier, validity time.Duration) TokenServerOption {
	return func(ts *TokenServer) {
		ts.clientCredentialVerifier = v
		ts.clientCredentialValidity = validity
	}
}

func WithPasswordFlow(v Verifier, validity time.Duration) TokenServerOption {
	return func(ts *TokenServer) {
		ts.passwordVerifier = v
		ts.passwordValidity = validity
	}
}

func (s *TokenServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// time out operations after one minute
	ctx, cancel := context.WithTimeout(request.Context(), time.Minute)
	defer cancel()
	logger := rpcutil.HTTPLogger(request, s.logger)

	writer.Header().Set("Cache-Control", "no-store")

	// check and extract request parameters
	if request.Method != http.MethodPost {
		writeTokenError(writer, errInvalidRequest, logger)
		return
	}
	if err := request.ParseForm(); err != nil {
		writeTokenError(writer, errInvalidRequest, logger)
		logger.Error("form parse error", zap.Error(err))
		return
	}

	var err error
	switch request.PostForm.Get("grant_type") {
	case "client_credentials":
		err = s.clientCredentialsFlow(ctx, writer, request)
	case "password":
		err = s.passwordFlow(ctx, writer, request)
	default:
		err = errUnsupportedGrantType
	}

	if err != nil {
		writeTokenError(writer, err, logger)
	}
}

func (s *TokenServer) clientCredentialsFlow(ctx context.Context, writer http.ResponseWriter, request *http.Request) error {
	if s.clientCredentialVerifier == nil {
		return errUnsupportedGrantType
	}
	clientId, clientSecret, err := s.clientCreds(request)
	if err != nil {
		return err
	}

	// lookup secret, and ensure it's for the matching client
	secretData, err := s.clientCredentialVerifier.Verify(ctx, clientId, clientSecret)
	if err != nil {
		return errInvalidClient
	}

	// generate an access token for the client
	token, err := s.tokens.GenerateAccessToken(secretData, s.clientCredentialValidity)
	if err != nil {
		return errors.New("failed to generate token")
	}

	// send response to the client
	response := tokenSuccessResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   int(s.clientCredentialValidity.Seconds()),
	}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		return errors.New("failed to marshal response")
	}

	_, err = writer.Write(responseBytes)
	if err != nil {
		return errors.New("failed to write response body")
	}
	return nil
}

func (s *TokenServer) passwordFlow(ctx context.Context, writer http.ResponseWriter, request *http.Request) error {
	if s.passwordVerifier == nil {
		return errUnsupportedGrantType
	}
	username, password, err := s.userCreds(request)
	if err != nil {
		return err
	}

	// lookup secret, and ensure it's for the matching client
	secretData, err := s.passwordVerifier.Verify(ctx, username, password)
	if err != nil {
		return errInvalidClient
	}

	// generate an access token for the client
	token, err := s.tokens.GenerateAccessToken(secretData, s.passwordValidity)
	if err != nil {
		return errors.New("failed to generate token")
	}

	// send response to the client
	response := tokenSuccessResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   int(s.passwordValidity.Seconds()),
	}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		return errors.New("failed to marshal response")
	}

	_, err = writer.Write(responseBytes)
	if err != nil {
		return errors.New("failed to write response body")
	}
	return nil
}

func (s *TokenServer) clientCreds(request *http.Request) (clientID string, clientSecret string, err error) {
	if !request.PostForm.Has("client_id") || !request.PostForm.Has("client_secret") {
		return "", "", errInvalidRequest
	}
	clientID = request.PostForm.Get("client_id")
	clientSecret = request.PostForm.Get("client_secret")
	return clientID, clientSecret, nil
}

func (s *TokenServer) userCreds(request *http.Request) (username string, password string, err error) {
	if !request.PostForm.Has("username") || !request.PostForm.Has("password") {
		return "", "", errInvalidRequest
	}
	username = request.PostForm.Get("username")
	password = request.PostForm.Get("password")
	return username, password, nil
}

func (s *TokenServer) TokenValidator() auth.TokenValidator {
	return s.tokens
}

func writeTokenError(writer http.ResponseWriter, err error, logger *zap.Logger) {
	tokErr, ok := err.(tokenError)
	if !ok {
		tokErr = tokenError{
			Code:             500,
			ErrorName:        "internal",
			ErrorDescription: err.Error(),
		}
	}

	body, marshalErr := json.Marshal(tokErr)
	if marshalErr != nil {
		logger.Error("failed to marshal error response", zap.Error(err))
		body = []byte(`{"error": "internal"}`)
	}

	writer.WriteHeader(tokErr.Code)
	_, writeErr := writer.Write(body)
	if writeErr != nil {
		logger.Error("failed to write error body", zap.Error(err))
	}
}

type tokenSuccessResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"` // in seconds
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
}

type tokenError struct {
	Code             int    `json:"-"` // http status code to use
	ErrorName        string `json:"error"`
	ErrorDescription string `json:"error_description,omitempty"`
	ErrorURI         string `json:"error_uri,omitempty"`
}

func (te tokenError) Error() string {
	if te.ErrorDescription != "" {
		return fmt.Sprintf("%s: %s", te.ErrorName, te.ErrorDescription)
	} else {
		return te.ErrorName
	}
}

var (
	errInvalidRequest       = tokenError{Code: 400, ErrorName: "invalid_request"}
	errInvalidClient        = tokenError{Code: 401, ErrorName: "invalid_client"}
	errUnsupportedGrantType = tokenError{Code: 400, ErrorName: "unsupported_grant_type"}
)
