package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	userRepo    UserRepository
	templateDir string
	tp          *TemplateRenderer
}

// routing - mux
// routing -> handlers -> controllers -> handler
// GET / - homepage
// POST /users - CreateUser

func main() {

	db, err := connectToDatabase("users_database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog:    log.New(os.Stderr, "ERROR\t", log.Ltime|log.LstdFlags|log.Lmicroseconds|log.Lshortfile),
		infoLog:     log.New(os.Stdout, "INFO\t", log.Ltime|log.LstdFlags),
		userRepo:    NewUserRepository(db),
		templateDir: "./13web/templates",
	}

	app.tp = NewTemplateRenderer(app.templateDir, false)

	log.Printf("Listening on port :8080")
	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}

func connectToDatabase(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
