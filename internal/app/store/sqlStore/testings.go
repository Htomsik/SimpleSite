package sqlStore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

// TestDb ...
func TestDb(t *testing.T, databaseType, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open(databaseType, databaseURL)

	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}

		db.Close()
	}
}