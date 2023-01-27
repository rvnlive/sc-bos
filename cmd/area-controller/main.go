package main

import (
	"context"
	"flag"
	"os"

	"github.com/vanti-dev/sc-bos/pkg/app"
	"github.com/vanti-dev/sc-bos/pkg/auto"
	"github.com/vanti-dev/sc-bos/pkg/auto/export"
	"github.com/vanti-dev/sc-bos/pkg/auto/lights"
	"github.com/vanti-dev/sc-bos/pkg/driver"
	"github.com/vanti-dev/sc-bos/pkg/driver/axiomxa"
	"github.com/vanti-dev/sc-bos/pkg/driver/bacnet"
	"github.com/vanti-dev/sc-bos/pkg/driver/mock"
	"github.com/vanti-dev/sc-bos/pkg/driver/xovis"
	"github.com/vanti-dev/sc-bos/pkg/node/alltraits"
	"github.com/vanti-dev/sc-bos/pkg/system"
	"github.com/vanti-dev/sc-bos/pkg/system/alerts"
	"github.com/vanti-dev/sc-bos/pkg/system/publications"
	"github.com/vanti-dev/sc-bos/pkg/system/tenants"
	"github.com/vanti-dev/sc-bos/pkg/testapi"

	"go.uber.org/zap"

	"github.com/vanti-dev/sc-bos/pkg/gen"
)

var systemConfig app.SystemConfig

func init() {
	flag.StringVar(&systemConfig.ListenGRPC, "listen-grpc", ":23557", "address (host:port) to host a Smart Core gRPC server on")
	flag.StringVar(&systemConfig.ListenHTTPS, "listen-https", ":443", "address (host:port) to host a HTTPS server on")
	flag.StringVar(&systemConfig.DataDir, "data-dir", ".data/area-controller-01", "path to local data storage directory")
	flag.StringVar(&systemConfig.StaticDir, "static-dir", "", "(optional) path to directory to host static files over HTTP from")

	flag.BoolVar(&systemConfig.DisablePolicy, "insecure-disable-policy", false, "Insecure! Disable checking requests against the security policy. This option opens up the server to any request.")
	flag.BoolVar(&systemConfig.LocalOAuth, "local-auth", false, "Enable issuing password tokens based on credentials found in users.json")
	flag.BoolVar(&systemConfig.TenantOAuth, "tenant-auth", false, "enable issuing client tokens based on credentials found in tenants.json or verified via the enrollment manager node")
}

func main() {
	os.Exit(app.RunUntilInterrupt(run))
}

func run(ctx context.Context) error {
	flag.Parse()
	systemConfig.LocalConfigFileName = "area-controller.local.json"
	systemConfig.Logger = zap.NewDevelopmentConfig()
	systemConfig.DriverFactories = map[string]driver.Factory{
		axiomxa.DriverName: axiomxa.Factory,
		bacnet.DriverName:  bacnet.Factory,
		mock.DriverName:    mock.Factory,
		xovis.DriverName:   xovis.Factory,
	}
	systemConfig.AutoFactories = map[string]auto.Factory{
		lights.AutoType: lights.Factory,
		"export-mqtt":   export.MQTTFactory,
	}
	systemConfig.SystemFactories = map[string]system.Factory{
		"alerts":       alerts.Factory,
		"tenants":      tenants.Factory,
		"publications": publications.Factory,
	}

	controller, err := app.Bootstrap(ctx, systemConfig)
	if err != nil {
		return err
	}

	alltraits.AddSupport(controller.Node)

	gen.RegisterTestApiServer(controller.GRPC, testapi.NewAPI())

	return controller.Run(ctx)
}
