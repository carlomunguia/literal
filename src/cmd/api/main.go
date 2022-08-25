package main

import (
	"fmt"
	"literal/internal/driver"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type config struct {
	port int
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	db       *driver.DB
}

func main() {
	var cfg config
	cfg.port = 8081

	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	dsn := os.Getenv("DSN")
	db, err := driver.ConnectPostgres(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.SQL.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		db:       db,
	}

	err = app.serve()
	if err != nil {
		errorLog.Fatal(err)
	}

}

func (app *application) serve() error {
	app.infoLog.Println("Listening on port", app.config.port)
	serve := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	return serve.ListenAndServe()
}