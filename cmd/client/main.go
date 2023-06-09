package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/zookeeper"
	providerV1 "go-micro-base-template/api/provider/v1"
	"go-micro-base-template/configs"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"os"
	"strconv"
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
	flag.Parse()
	configs.GetConfig()
}
func main() {
	address := configs.GetConfig().RegistrarCfg.ZookeeperCfg.Addr
	sessionTimeout := configs.GetConfig().RegistrarCfg.ZookeeperCfg.SessionTimeout
	register := zookeeper.NewRegistry(
		registry.Addrs(address...),
		registry.Timeout(time.Duration(sessionTimeout)*time.Second),
	)

	service := micro.NewService(micro.Registry(register))

	client := providerV1.NewProviderService(configs.GetConfig().Server.ServiceName, service.Client())

	var i = 0
	begin := time.Now().Unix()
	for true {
		i++
		body, _ := anypb.New(&wrapperspb.StringValue{Value: "小明#########################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################" +
			"#####################################################################"})

		request := &providerV1.Request{
			ReqId: strconv.Itoa(i),
			Event: providerV1.EVENT_CODE_ADD_USER,
			Body:  body,
		}

		_, err3 := client.BaserService(context.Background(), request)
		if err3 != nil {
			fmt.Printf("erro=%v \n", err3)
		}
		if i%100000 == 0 {
			endTime := time.Now().Unix()
			fmt.Println(strconv.FormatInt(endTime-begin, 10) + "s")
			begin = endTime
		}
		/*		fmt.Printf("response=%v, erro=%v \n", response, err3)
				time.Sleep(time.Second * 1)*/
	}
}
