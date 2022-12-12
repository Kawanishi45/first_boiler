package main

import (
	"context"
	"database/sql"
	"fmt"
	"kawanishi/first_boiler/api_srv/config"
	"kawanishi/first_boiler/api_srv/infrastracture/persistence"
	"kawanishi/first_boiler/api_srv/server"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"os"
)

const (
	exitOK  = 0
	exitErr = 1
)

func main() {
	ctx := context.Background()

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup logger: %s\n", err)
		logger.Debug(fmt.Sprintf("%v", exitErr))
	}
	defer logger.Sync()

	cfg, err := config.LoadEnv(ctx)
	if err != nil {
		logger.Error("failed to load env", zap.Error(err))
		logger.Debug(fmt.Sprintf("%v", exitErr))
	}
	// init mysql
	db, err := sql.Open("mysql", cfg.Dsn)
	if err != nil {
		logger.Error("failed to open mysql", zap.Error(err))
		logger.Debug(fmt.Sprintf("%v", exitErr))
	}
	err = db.PingContext(ctx)
	if err != nil {
		logger.Error("failed to ping db", zap.Error(err))
		logger.Debug(fmt.Sprintf("%v", exitErr))
	}

	repositories, err := persistence.NewRepositories(ctx, db)
	if err != nil {
		logger.Error("failed to init repository", zap.Error(err))
		logger.Debug(fmt.Sprintf("%v", exitErr))
	}

	// init newServer
	newServer, err := server.NewServer(repositories)
	if err != nil {
		logger.Error("failed to init newServer", zap.Error(err))
		logger.Debug(fmt.Sprintf("%v", exitErr))
	}

	fmt.Println(exitOK)

	err = newServer.User(ctx, 1)
	if err != nil {
		logger.Debug(fmt.Sprintf("%v", err))
	}
}
