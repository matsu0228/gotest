package repository

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // raw connector
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Database is struct of database models
type Database struct {
	db         *sqlx.DB
	timeFormat string
}

func generateDatabaseQuery(user, pass, host, port, dbname, option string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", user, pass, host, port, dbname, option)
}

// NewDatabase is constructor
func NewDatabase(user, pass, host, port, dbname, option string) (Database, error) {

	const createdFormat = `2006-01-02 15:04:05` // +0700 JST`

	connectQuery := generateDatabaseQuery(user, pass, host, port, dbname, option)
	databaseConn, err := sqlx.Connect("mysql", connectQuery)
	if err != nil {
		return Database{}, err
	}
	if err := databaseConn.Ping(); err != nil {
		return Database{}, err
	}
	return Database{
		db:         databaseConn,
		timeFormat: createdFormat,
	}, err
}

// Get is getter
func (r Database) Get(title string) ([]SomeData, error) {

	var dataset []SomeData

	// sClmName := "id, title, body, created_at, updated_at "
	sClmName := " * "
	err := r.db.Select(&dataset, "SELECT "+sClmName+" FROM todo.tasks where title = ?", title)
	if err != nil {
		return []SomeData{}, err
	}

	return dataset, nil
}

// Save is updater or inserter
func (r Database) Save(title, body string) error {

	dataset, err := r.Get(title)
	if err != nil {
		errors.Wrapf(err, "Save() with %v", title)
	}

	now := time.Now().Format(r.timeFormat)
	log.Printf("[INFO] time: %v", now)

	// insert
	if len(dataset) == 0 {
		tsql := `INSERT INTO tasks(title, body, updated_at, created_at) VALUES(?, ?, ?, ?);`
		res := r.db.MustExec(tsql, title, body, now, now)
		added, err := res.RowsAffected()
		log.Printf("[INFO] insert data %#v, %#v", added, err)
		return nil
	}

	// update
	tsql := `UPDATE tasks SET title= ?, body = ?, updated_at = ? WHERE title = ?;`
	res := r.db.MustExec(tsql, title, body, now, title)
	added, err := res.RowsAffected()
	log.Printf("[INFO] update data %#v, %#v", added, err)

	return nil
}
