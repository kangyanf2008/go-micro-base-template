package configs

import (
	"flag"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/yaml"
	"github.com/micro/go-micro/config/reader"
	"github.com/micro/go-micro/config/reader/json"
	"github.com/micro/go-micro/config/source/file"
)

var (
	flagConf string
	conf     SysConfig
)

func GetConfig() SysConfig {
	return conf
}

type GrpcCfg struct {
	Addr    string `toml:"addr"`
	Version string `toml:"version"`
	Timeout int64  `toml:"timeout"`
}
type HttpCfg struct {
	Addr    string `toml:"addr"`
	Timeout int64  `toml:"timeout"`
}
type ServerCfg struct {
	GRPC        GrpcCfg `toml:"grpc"`
	Http        HttpCfg `toml:"http"`
	ServiceName string  `toml:"serviceName"`
}

type ZookeeperCfg struct {
	Addr           []string `toml:"addr"`
	Namespace      string   `toml:"namespace"`
	SessionTimeout int64    `toml:"sessionTimeout"`
}

type RegistrarCfg struct {
	ZookeeperCfg ZookeeperCfg `toml:"zookeeper"`
}

type SysConfig struct {
	Server       ServerCfg    `toml:"server"`
	RegistrarCfg RegistrarCfg `toml:"registrar"`
}

func init() {
	flag.StringVar(&flagConf, "conf", "../configs/config.yaml", "config path, eg: -conf config.yaml")
	flag.Parse()
	enc := yaml.NewEncoder()
	c := config.NewConfig(config.WithReader(
		json.NewReader( // json reader for internal config merge
			reader.WithEncoder(enc),
		),
	))

	// load the config from a file source
	if err := c.Load(file.NewSource(
		file.WithPath(flagConf),
	)); err != nil {
		panic(err)
		return
	}

	// define our own host type

	// read a database host
	if err := c.Scan(&conf); err != nil {
		panic(err)
	}
}
