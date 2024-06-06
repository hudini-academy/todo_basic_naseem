package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

func initializeLoggers() (*log.Logger, *log.Logger) {
	// here we are creating a file of info logger  and appending the message
	f, err := os.OpenFile("./cmd/web/tmp/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// here we are creating a file of error logger  and appending the message
	e, err := os.OpenFile("./cmd/web/tmp/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// we are creating a new custom logger to handle info and error. it will have three parameters.
	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime|log.LUTC)
	errorLog := log.New(e, "EROOR/t", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)

	return infoLog, errorLog
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack()) // stores the error message and stacktrace.
	app.errorLog.Output(2, trace)                              // To get the exact file and line number.
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
