package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	HttpServer `yaml:"http_server"`
	Postgres   `yaml:"postgres"`
}

type HttpServer struct {
	Address     string        `yaml:"address"`
	Port        string        `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"ssl_mode"`
}

func ConfigLoad() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Config{
		HttpServer{
			Address:     viper.GetString("http_server.address"),
			Port:        viper.GetString("http_server.port"),
			Timeout:     getDurationEnv(viper.GetString("http_server.timeout"), 10*time.Second),
			IdleTimeout: getDurationEnv(viper.GetString("http_server.idle_timeout"), 30*time.Second),
		},
		Postgres{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetString("postgres.port"),
			User:     viper.GetString("postgres.user"),
			Database: viper.GetString("postgres.database"),
			//Password: os.Getenv("POSTGRES_PASSWORD"),
			Password: "postgres",
			SSLMode:  viper.GetString("postgres.ssl_mode"),
		},
	}, nil
}

func getDurationEnv(value string, defaultValue time.Duration) time.Duration {
	duration, err := time.ParseDuration(value)
	if err != nil {
		log.Printf("Invalid duration: %s, using default: %v\n", value, defaultValue)
		return defaultValue
	}
	return duration
}
