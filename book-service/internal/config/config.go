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
	Port int `yaml:"port" env:"BOOK_SERVICE_PORT"`
}

type DbConfig struct {
	Port     int    `yaml:"port" env:"POSTGRES_DB_PORT"`
	Host     string `yaml:"host" env:"POSTGRES_DB_HOST"`
	Username string `yaml:"username" env:"POSTGRES_DB_USERNAME"`
	Password string `yaml:"password" env:"POSTGRES_DB_PASSWORD"`
	Dbname   string `yaml:"db_name" env:"POSTGRES_DB_NAME"`
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

func FetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config-path", "", "load config from path")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
