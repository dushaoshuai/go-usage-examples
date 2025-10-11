package gorm_learn

import (
	"context"
	"testing"
)

func TestPostgreSQLDB(t *testing.T) {
	db := PostgreSQLDB(context.Background())
	t.Log(db.Migrator().CurrentDatabase())

	// Output:
	// 2025/10/10 23:29:28 /tmp/go-usage-examples/gorm_learn/db_test.go:10
	// [3.648ms] [rows:1] SELECT CURRENT_DATABASE()
	//    db_test.go:10: defaultpostgresqldb
	// --- PASS: TestPostgreSQLDB (0.01s)
	// PASS
}
