package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"sync"
)

type GrpcConfig struct {
	Port string `yaml:"port" env:"GRPCPORT"`
}
type DBConfig struct {
	Host string `yaml:"host" env:"DBHOST"`
	Port string `yaml:"port" env:"DBPORT"`
}

type Config struct {
	GrpcConfig GrpcConfig `yaml:"grpc"`
	DBConfig   DBConfig   `yaml:"db"`
}

var instance *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		path := FetchConfigPath()
		instance = GetConfigByPath(path)
	})
	return instance
}

func FetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config-path", "", "config path")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}

func GetConfigByPath(path string) *Config {
	var cfg Config
	if path != "" {
		if err := cleanenv.ReadConfig(path, &cfg); err != nil {
			panic(err)
		}
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
