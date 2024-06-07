package main

import (
	"flag"
	"log"
	"naseem/pkg/models/mysql"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
)

// creating an structure of application which contains errorLog and infoLog
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	todo     *mysql.TodoModel
	session  *sessions.Session
	users    *mysql.UserModel
	special  *mysql.SpecialModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	// creating a connection to db and validating password and user name
	dsn := flag.String("dsn", "root:root@/todo?parseTime=true", "mysql database connection")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()

	infoLog, errorLog := initializeLoggers()

	// we are opening databse using OpenDB function
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	// creating an instance of application structure
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		todo:     &mysql.TodoModel{DB: db},
		session:  session,
		users:    &mysql.UserModel{DB: db},
		special:  &mysql.SpecialModel{DB: db},
	}

	// we are creating an instance of http.Server and adding fields to it .
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// we are calling our custom info logger
	infoLog.Printf("Starting server on %s", *addr)
	// we are calling the instance of http.Server and connect to the ListenAndServe
	srv.ListenAndServe()
	// we are calling our custom error logger
	errorLog.Fatal(err)
}
