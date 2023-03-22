package postgres

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

var ddlConnStr string

func TestMain(m *testing.M) {
	const dbConnEnv = "DATABASE_CONNECTION"
	ddlConnStr = os.Getenv(dbConnEnv)
	if ddlConnStr == "" {
		log.Printf("%s is not set, skipping", dbConnEnv)
		return
	}

	exitCode := m.Run()
	defer os.Exit(exitCode)
}

func getDBConnectionString() string{
	return ddlConnStr
}

func getMigrationDir() string{
	return filepath.Join("..", "..", "migrations")
}