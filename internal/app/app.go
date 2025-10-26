package app

import (
	"code-sharing-be/cmd/config"
	"code-sharing-be/internal/controller"
	"code-sharing-be/internal/repo"
	"code-sharing-be/internal/usecase/code_block"
	"code-sharing-be/pkg/httpserver"
	"code-sharing-be/pkg/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	id       int
	username string
}

func Run(config *config.Config) {
	// Logger
	l := logger.New(config.Log.Level)

	// Repository
	db, err := createMysqlConfig(config)

	if err != nil {
		panic("Could not establish database connection")
	}
	l.Info("Database Connected")

	codeBlockRepo := repo.New(db)

	codeBlockUseCase := code_block.New(*codeBlockRepo)

	server := httpserver.New(config)

	controller.NewRouter(server.Engine, config, codeBlockUseCase)

	server.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case err = <-server.Notify():
		l.Error(fmt.Errorf("app: Run -> httpServer.Notify: %w", err))
	}
}

func createMysqlConfig(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		config.MYSQL.User,
		config.MYSQL.Pass,
		config.MYSQL.URL,
		config.MYSQL.DBName)
	fmt.Printf("Connection string is %s", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if db == nil {
		return nil, err
	}

	return db, nil
}
