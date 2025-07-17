package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"sync"
)

type Config struct {
	HttpServer   Server       `yaml:"httpServer"`
	BookService  BookService  `yaml:"bookService"`
	OrderService OrderService `yaml:"orderService"`
}

type Server struct {
	Port int `yaml:"port"`
}

type BookService struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type OrderService struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		path := FetchConfigPath()
		instance = LoadConfigByPath(path)
	})
	return instance
}

func LoadConfigByPath(path string) *Config {
	var cfg Config
	err := cleanenv.ReadConfig(path, &cfg)

	if err != nil {
		panic(err)
	}
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
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
