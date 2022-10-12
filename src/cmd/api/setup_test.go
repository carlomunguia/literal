package main

import (
	"log"
	"os"
	"testing"

	"literal/internal/data"

	"github.com/DATA-DOG/go-sqlmock"
)

var (
	testApp application
	mockDB  sqlmock.Sqlmock
)

func TestMain(m *testing.M) {
	testDB, mock, _ := sqlmock.New()

	mockDB = mock

	defer testDB.Close()

	testApp = application{
		config:      config{},
		infoLog:     log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog:    log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		models:      data.New(testDB),
		environment: "development",
	}

	os.Exit(m.Run())
}
