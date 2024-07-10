package main

import (
	"EM-Api-testTask/dummyServer"
	"EM-Api-testTask/internal/server"
	"github.com/kpango/glg"
)

func main() {
	glg.Get().SetLineTraceMode(glg.TraceLineShort)
	dummyServer.StartDummy()

	app, err := server.NewApp("./config")
	if err != nil {
		glg.Fatal(err)
	}
	if err := app.Run(); err != nil {
		glg.Fatal(err)
	}

}
