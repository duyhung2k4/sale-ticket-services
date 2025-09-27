package connection

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig `mapstructure:"app"`
}

type AppConfig struct {
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	GrpcPort string `mapstructure:"grpc_port"`
	HttpPort string `mapstructure:"http_port"`
}

var (
	ConfigInfo *Config = &Config{}
)

func (s *Config) loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error load config: ", err.Error())
		return
	}

	if err := viper.Unmarshal(&s); err != nil {
		fmt.Println("error map config to struct: ", err.Error())
		return
	}
}
