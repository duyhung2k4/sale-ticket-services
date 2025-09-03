package connection

import "github.com/spf13/viper"

type ConnectionConfigModel struct {
	App    AppConnection `mapstructure:"app"`
	JwtKey string        `mapstructure:"jwt-key"`
}

type AppConnection struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

var connectionConfig *ConnectionConfigModel

func loadConfig() error {
	viper.SetConfigFile("config/dev.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&connectionConfig); err != nil {
		return err
	}

	return nil
}

func GetConfig() *ConnectionConfigModel {
	return connectionConfig
}
