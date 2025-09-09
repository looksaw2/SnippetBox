package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql" // New import
	"log"
	"net/http"
	"os"
)

type application struct {
	errLog  *log.Logger
	infoLog *log.Logger
}

func main() {
	//解析配置
	addr := flag.String("addr", ":8080", "This is the go server run on port")
	//解析DSN
	dsn := flag.String("dsn", "root@rootpassword/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	//
	db, err := openDB(*dsn)
	if err != nil {
		errlog.Fatal(err)
	}
	defer db.Close()
	app := &application{
		errLog:  errlog,
		infoLog: infolog,
	}
	srv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errlog,
	}
	infolog.Printf("Starting server on : %s", *addr)
	err = srv.ListenAndServe()
	errlog.Fatal(err)
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
