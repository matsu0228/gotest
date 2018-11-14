package repository

import (
	"testing"

	"github.com/go-sql-driver/mysql"
)

func TestConnectQuery(t *testing.T) {
	got := generateDatabaseQuery("user", "pass", "port", "host", "dbname", "options")

	_, err := mysql.ParseDSN(got)
	if err != nil {
		t.Fatalf("cannt Parse DSN(DataSourceName). err=%v", err)
	}
	// pp.Print(cfg)
	// if got != want {
	// 	t.Fatalf("want %v, but %v:", want, got)
	// }
}
