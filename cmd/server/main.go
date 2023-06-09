package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/zookeeper"
	v1 "go-micro-base-template/api/provider/v1"
	"go-micro-base-template/configs"
	"go-micro-base-template/internal"
	"os"
	"time"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = ""
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	configs.GetConfig()
}
func main() {
	address := configs.GetConfig().RegistrarCfg.ZookeeperCfg.Addr
	sessionTimeout := configs.GetConfig().RegistrarCfg.ZookeeperCfg.SessionTimeout
	register := zookeeper.NewRegistry(
		registry.Addrs(address...),
		registry.Timeout(time.Duration(sessionTimeout)*time.Second),
	)

	grpcLicenseAddr := configs.GetConfig().Server.GRPC.Addr
	version := configs.GetConfig().Server.GRPC.Version
	service := micro.NewService(
		micro.Name(configs.GetConfig().Server.ServiceName),
		micro.Registry(register),
		micro.Address(grpcLicenseAddr),
		micro.Version(version),
	)

	v1.RegisterProviderHandler(service.Server(), &internal.ProviderServiceImpl{})

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
