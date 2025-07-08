package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"sync"
)

type Config struct {
	GRPCConf GRPCConfig `yaml:"grpc"`
	DbConf   DbConfig   `yaml:"db"`
}

type GRPCConfig struct {
	Port int `yaml:"port"`
}

type DbConfig struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"db_name"`
}

var instance *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		path := FetchConfigPath()
		instance = LoadConfigByPath(path)
	})
	return instance
}

func LoadConfigByPath(path string) *Config {
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(err)
	}
	return &cfg
}

func FetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config-path", "", "load config from path")
	flag.Parse()
	if res == "" {
		os.Getenv("CONFIG_PATH")
	}
	return res
}
