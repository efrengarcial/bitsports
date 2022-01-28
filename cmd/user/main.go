package main

import (
	"bitsports/config"
	"bitsports/ent"
	"bitsports/internal/user"
	"bitsports/internal/user/adapter"
	"bitsports/internal/user/infraestructure"
	"bitsports/pkg/datasource"
	"context"
	"github.com/labstack/echo/v4"
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

	r := adapter.NewRepository(client)
	uc := user.NewUseCase(r)
	e := echo.New()
	e = infraestructure.NewRouter(e, adapter.NewHandler(uc),logger )

	// Start server
	go func() {
		if err := e.Start(":" + config.C.UserServer.Address); err != nil && err != http.ErrServerClosed {
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
