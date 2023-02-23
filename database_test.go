package golang_mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(host:8889)/db_golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
