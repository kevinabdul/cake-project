package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"privy/internal/data"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type config struct {
	dsn string
}

type application struct {
	config config
	logger *log.Logger
	repos  data.Repos
}

// $ curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
func main() {
	cfg := config{dsn: fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE")),
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := connectToDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Printf("database connection is established")

	app := &application{
		config: cfg,
		repos:  data.NewRepos(db),
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("API_SERVER_PORT")),
		Handler: app.routes(),
	}

	logger.Printf("starting server on %s", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}

// wait for 10 minutes for database initialization
// feel free to change the waiting time for database initialization
// based on your computer's capacity and speed
func connectToDB(cfg config) (*sql.DB, error) {
	var counts int
	for {
		connection, err := openDB(cfg)
		if err != nil {
			log.Println("MySQL not yet ready ...")
			counts++
		} else {
			log.Println("Connected to MySQL")
			return connection, nil
		}

		if counts > 200 {
			log.Println(err)
			return nil, err
		}

		time.Sleep(3 * time.Second)
		continue
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
