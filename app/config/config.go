package config

import (
	"fmt"
	"github.com/kpango/glg"
	"github.com/spf13/viper"
)

type Config struct {
	Psql_connection string
	Addr            string
	Port            int
	Debug           bool
	//SideServerUrl   string
}

func NewConfig(pth string) (*Config, error) {
	glg.Debug("Starting config init")
	viper.AutomaticEnv()
	viper.AddConfigPath(pth)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		glg.Warnf("Error reading config file: %v\n", err)
	}
	Psql_connection := fmt.Sprintf("postgres://%s:%s@%s/%s",
		viper.GetString("POSTGRES_USER"),
		viper.GetString("POSTGRES_PASSWORD"),
		viper.GetString("POSTGRES_ADDR"),
		viper.GetString("POSTGRES_DB"))

	addr := viper.GetString("APP_ADDR")
	port := viper.GetInt("APP_PORT")
	logDebug := viper.GetBool("LOG_DEBUG")
	//SideServerUrl := viper.GetString("FETCH_SERVER_URL")

	if !logDebug {
		glg.Get().SetLevel(glg.INFO)
	}

	c := Config{
		Psql_connection: Psql_connection,
		Addr:            addr,
		Port:            port,
		Debug:           logDebug,
		//SideServerUrl:   SideServerUrl,
	}

	return &c, nil
}
