package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"sync"
)

type GrpcConfig struct {
	Port int `yaml:"port" env:"ORDER_GRPC_PORT"`
}
type DBConfig struct {
	Host     string `yaml:"host" env:"MONGO_DB_HOST"`
	Port     string `yaml:"port" env:"MONGO_DB_PORT"`
	Username string `yaml:"username" env:"MONGO_DB_USERNAME"`
	Password string `yaml:"password" env:"MONGO_DB_PASSWORD"`
}
type BookService struct {
	Host string `yaml:"host" env:"BOOK_SERVICE_HOST"`
	Port string `yaml:"port" env:"BOOK_SERVICE_PORT"`
}

type Config struct {
	GrpcConfig  GrpcConfig  `yaml:"grpc"`
	DBConfig    DBConfig    `yaml:"db"`
	BookService BookService `yaml:"bookService"`
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
