package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Psql_connection string
	Addr            string
	Port            int
	Debug           bool
	SideServerUrl   string
}

func NewConfig(pth string) (*Config, error) {

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	PSQL_CONNECTION_STRING := viper.GetString("PSQL_CONNECTION_STRING")
	viper.AddConfigPath(pth)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	addr := viper.GetString("app.addr")
	port := viper.GetInt("app.port")
	logDebug := viper.GetBool("app.log_debug")
	SideServerUrl := viper.GetString("app.SideServerUrl")

	c := Config{
		Psql_connection: PSQL_CONNECTION_STRING,
		Addr:            addr,
		Port:            port,
		Debug:           logDebug,

		SideServerUrl: SideServerUrl,
	}

	// Example usage

	return &c, nil
}
