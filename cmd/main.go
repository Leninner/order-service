package main

import (
	"log"
	"log/slog"
	"os"
	"sync"

	"github.com/leninner/order-service/internal/application/routes"
	"github.com/leninner/shared/config"
	"github.com/leninner/shared/server"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &config.Application{
		Logger: logger,
		WG:     sync.WaitGroup{},
		Config: config.Config{
			Port: 8080,
			Env:  "development",
		},
	}

	handler := routes.Routes(app)

	err := server.Serve(app, handler)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
