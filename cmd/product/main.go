package main

import (
	"bitsports/config"
	"bitsports/ent"
	"bitsports/internal/product/adapter/graphql"
	"bitsports/internal/product/adapter/repository"
	"bitsports/internal/product/application/usecase"
	infraGraphql "bitsports/internal/product/infraestructure/graphql"
	"bitsports/internal/product/infraestructure/router"
	"bitsports/pkg/datasource"
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Create a new instance of the logger.
var logger = logrus.New()

func init() {
	logger.Out = os.Stdout
	//logger.Formatter = &logrus.JSONFormatter{PrettyPrint : true}
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true, // Seems like automatic color detection doesn't work on windows terminals
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})
	logger.Level = logrus.InfoLevel
}

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	if config.C.Debug {
		logger.Level = logrus.DebugLevel
		logger.Debug("Service RUN on DEBUG mode")
	}

	client := newDBClient()

	repoProduct := repository.NewProduct(client)
	ucProduct := usecase.NewProduct(repoProduct)
	repoCategory := repository.NewCategory(client)
	ucCategory := usecase.NewCategory(repoCategory)

	schema :=  graphql.NewSchema(graphql.NewProductResolver(ucProduct),
		graphql.NewCategoryResolver(ucCategory))

	handler, err := infraGraphql.NewServer(schema)

	if err!= nil {
		logger.Fatal(err)
	}

	e := router.New(handler,logger)

	// Start server
	go func() {
		if err := e.Start(":" + config.C.Server.Address); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func newDBClient() *ent.Client {
	client, err := datasource.NewClient()
	if err != nil {
		logger.Fatalf("failed opening mysql client: %v", err)
	}

	return client
}

