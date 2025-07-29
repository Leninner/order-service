package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/leninner/infrastructure/kafka"
	"github.com/leninner/order-service/internal/application"
	"github.com/leninner/order-service/internal/application/routes"
	appConfig "github.com/leninner/order-service/internal/config"
	sharedConfig "github.com/leninner/shared/config"
	sharedDI "github.com/leninner/shared/di"
	"github.com/leninner/shared/server"

	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	config := appConfig.LoadConfig()
	
	// Create logger
	loggerInstance := slog.New(slog.NewTextHandler(os.Stdout, nil))
	
	// Create shared container
	sharedContainer := sharedDI.NewSharedContainer(config.Config, loggerInstance)
	
	// Initialize database
	db, err := server.OpenDB(config.Config)
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		panic(err)
	}
	sharedContainer.SetDatabase(db)
	
	// Initialize Kafka
	kafkaModule := kafka.NewKafkaModule()
	sharedContainer.SetKafka(kafkaModule)

	// Create order service container
	container := application.NewOrderServiceContainer(sharedContainer)

	defer container.Close()

	handler := routes.Routes(container)

	app := &sharedConfig.Application{
		Config: config.Config,
		Logger: container.GetLogger(),
		DataSource: container.GetDatabase(),
	}

	err = server.Serve(app, handler)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
} 