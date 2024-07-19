package main

import (
	"EM-Api-testTask/internal/server"

	_ "EM-Api-testTask/docs" // Import generated swagger docs
	"github.com/kpango/glg"
)

func main() {

	// @title Task tracker api
	// @version 1.0
	// @description Test task for Effective Mobile
	// @host localhost:8080
	// @BasePath /

	glg.Get().SetLineTraceMode(glg.TraceLineShort)

	app, err := server.NewApp("./config")
	if err != nil {
		glg.Fatal(err)
	}

	if err := app.Run(); err != nil {
		glg.Fatal(err)
	}

}
