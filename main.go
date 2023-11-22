package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sqlx.Connect("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatalln(err)
	}

	view, err := NewTaskView()
	if err != nil {
		log.Fatalln(err)
	}

	repository := NewTaskRepository(db)

	args := GetArgs()

	if err := NewTaskController(args, view, repository).Serve(); err != nil {
		log.Fatalln(err)
	}

}
