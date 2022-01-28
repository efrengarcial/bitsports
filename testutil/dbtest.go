package testutil

import (
	"bitsports/config"
	"bitsports/ent"
	"bitsports/ent/enttest"
	"bitsports/pkg/database"
	"bitsports/pkg/datasource"
	"bitsports/pkg/docker"
	"context"
	"entgo.io/ent/dialect"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
	"time"
)

var logger = logrus.New()

// StartDB starts a database instance.
func StartDB() (*docker.Container, error) {
	image := "postgres:14-alpine"
	port := "5432"
	args := []string{"-e", "POSTGRES_PASSWORD=postgres"}

	return docker.StartContainer(image, port, args...)
}

// StopDB stops a running database instance.
func StopDB(c *docker.Container) {
	docker.StopContainer(c.ID)
}

// NewUnit creates a test database inside a Docker container. It creates the
// required table structure but the database is otherwise empty. It returns
// the database to use as well as a function to call at the end of the test.
func NewUnit(t *testing.T, c *docker.Container, dbName string) (*logrus.Logger, *ent.Client, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbM, err := database.Open(database.Config{
		User:       "postgres",
		Password:   "postgres",
		Host:       c.Host,
		Name:       "postgres",
		DisableTLS: true,
	})
	if err != nil {
		t.Fatalf("Opening database connection: %v", err)
	}

	t.Log("Waiting for database to be ready ...")

	if err := database.StatusCheck(ctx, dbM); err != nil {
		t.Fatalf("status check database: %v", err)
	}

	t.Log("Database ready")

	if _, err := dbM.ExecContext(context.Background(), "CREATE DATABASE "+dbName); err != nil {
		t.Fatalf("creating database %s: %v", dbName, err)
	}
	dbM.Close()

	ReadConfig()
	config.C.Database.Addr = c.Host
	config.C.Database.DBName = dbName
	client := NewDBClient(t)

	t.Log("Ready for testing ...")

	logger.Out = os.Stdout
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true, // Seems like automatic color detection doesn't work on windows terminals
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})
	logger.Level = logrus.DebugLevel

	// teardown is the function that should be invoked when the caller is done
	// with the database.
	teardown := func() {
		t.Helper()
		client.Close()
	}

	return logger, client, teardown
}

// Test owns state for running and shutting down tests.
type Test struct {
	Client   *ent.Client
	Log      *logrus.Logger
	Teardown func()

	t *testing.T
}

// NewDBClient loads database for test.
func NewDBClient(t *testing.T) *ent.Client {
	d := datasource.New()
	return enttest.Open(t, dialect.Postgres, d)
}

// NewIntegration creates a database, seeds it.
func NewIntegration(t *testing.T, c *docker.Container, dbName string) *Test {
	log, db, teardown := NewUnit(t, c, dbName)

	test := Test{
		Client:   db,
		Log:      log,
		t:        t,
		Teardown: teardown,
	}

	return &test
}
