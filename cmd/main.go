package main

import (
	"flag"
	"log"
	"log/slog"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/leninner/order-service/internal/application/routes"
	"github.com/leninner/shared/config"
	"github.com/leninner/shared/server"
)

func main() {
	var appConfig config.Config

	flag.IntVar(&appConfig.Port, "port", 4000, "API server port")
	flag.StringVar(&appConfig.Env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&appConfig.DB.DSN, "db-dsn", "postgres://order_role:pa55word@localhost/order_db?sslmode=disable", "Database connection string")
	flag.IntVar(&appConfig.DB.MaxOpenConns, "db-max-open-conns", 25, "Maximum number of open connections to the database")
	flag.IntVar(&appConfig.DB.MaxIdleConns, "db-max-idle-conns", 25, "Maximum number of idle connections to the database")
	flag.DurationVar(&appConfig.DB.MaxIdleTime, "db-max-idle-time", 15*time.Minute, "Maximum amount of time a connection may be idle before being closed")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := server.OpenDB(appConfig)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()
	app := &config.Application{
		Logger: logger,
		Config: appConfig,
		DataSource: db,
	}

	handler := routes.Routes(app)

	err = server.Serve(app, handler)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

