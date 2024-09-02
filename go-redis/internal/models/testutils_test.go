package models

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func newTestDB(t *testing.T) *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), "postgres://your_user:your_password@localhost:5432/your_db?sslmode=disable")
	if err != nil {
		t.Fatalf("failed to connect to db: %v", err)
	}

	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatalf("failed to read setup.sql: %v", err)
	}

	_, err = db.Exec(context.Background(), string(script))
	if err != nil {
		t.Fatalf("failed to exec setup.sql: %v", err)
	}

	t.Cleanup(func() {
		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatalf("failed to read teardown.sql: %v", err)
		}

		_, err = db.Exec(context.Background(), string(script))
		if err != nil {
			t.Fatalf("failed to exec teardown.sql: %v", err)
		}
		db.Close()
	})
	return db
}
