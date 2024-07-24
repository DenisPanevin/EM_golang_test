package main

import (
	_ "EM-Api-testTask/docs" // Import generated swagger docs
	"EM-Api-testTask/internal/server"
	"flag"
	"github.com/kpango/glg"
)

func main() {

	// @title Task tracker api
	// @version 1.0
	// @description Test task for Effective Mobile
	// @host localhost:8080
	// @BasePath /

	var configPath string
	flag.StringVar(&configPath, "config-path", "./", "path to config file")
	flag.Parse()
	glg.Get().SetLineTraceMode(glg.TraceLineShort)

	app, err := server.NewApp(configPath)
	if err != nil {
		glg.Fatal(err)
	}

	if err := app.Run(); err != nil {
		glg.Fatal(err)
	}

}
