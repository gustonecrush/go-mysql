package golang_mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO customer(id, name) VALUES ('arum', 'Arum')"

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert to database")
}
