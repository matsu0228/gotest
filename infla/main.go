package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type task struct {
	ID        int       `db:"id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Repository is struct of database models
type Repository struct {
	db *sqlx.DB
}

// NewRepository is constructor
func NewRepository(user, pass, host, port, dbname, option string) (Repository, error) {
	// connectQuery := fmt.Sprintf("%s:%s@%s:%s/%s", user, pass, host, port, dbname)
	connectQuery := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", user, pass, host, port, dbname, option)
	log.Print("[INFO] NewRepository() with ", connectQuery)
	// databaseConn, err := sql.Open("mysql", connectQuery)
	databaseConn, err := sqlx.Connect("mysql", connectQuery)
	if err != nil {
		return Repository{}, err
	}
	if err := databaseConn.Ping(); err != nil {
		return Repository{}, err
	}
	return Repository{db: databaseConn}, err
}

func (r Repository) selectData() ([]task, error) {

	var tasks []task

	// sClmName := "id, title, body, created_at, updated_at "
	sClmName := " * "
	// rows, err := r.db.Query("SELECT " + sClmName + " FROM todo.tasks") //
	err := r.db.Select(&tasks, "SELECT "+sClmName+" FROM todo.tasks")

	if err != nil {
		return []task{}, err
	}

	// columns, err := rows.Columns() // カラム名を取得
	// if err != nil {
	// 	return []task{}, err
	// }
	// log.Print("[INFO] columns", columns)
	// return nil

	// for rows.Next() {
	// 	t := task{}
	// 	if err := rows.StructScan(&t); err != nil {
	// 		// rows.Scan(t); err != nil {
	// 		return []task{}, err
	// 	}
	// 	tasks = append(tasks, t)
	// }
	return tasks, nil
}

func (r Repository) addData(title, body string) error {
	tsql := `INSERT INTO tasks(title, body) VALUES(?, ?);`
	res := r.db.MustExec(tsql, title, body)

	added, err := res.RowsAffected()
	log.Printf("[INFO] addData() %#v, %#v", added, err)
	return nil
}

func main() {
	// defer db.Close() // 関数がリターンする直前に呼び出される

	// @mac hosts設定しておく
	//   $ cat /etc/hosts | grep docker-mysql
	//   -> 127.0.0.1       docker-mysql
	repo, err := NewRepository("root", "mysql", "docker-mysql", "3306", "todo", "?parseTime=true&loc=Japan")
	if err != nil {
		log.Fatal("error=", err)
	}

	tasks, err := repo.selectData()
	if err != nil {
		log.Fatal("selectData() error=", err)
	}

	err = repo.addData("testTitle", "--body-----")
	if err != nil {
		log.Fatal("addData() error=", err)
	}

	format := "2006-01-02 15:04:05"
	cAt := tasks[0].CreatedAt.Format(format)
	uAt := tasks[0].UpdatedAt.Format(format)
	log.Printf("[INFO] result: %#v c:%v u:%v", tasks, cAt, uAt)
}
