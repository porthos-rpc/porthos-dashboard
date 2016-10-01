package storage

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/porthos-rpc/porthos-dashboard/models"
)

// DBStorage holds a db connection.
type DBStorage struct {
	db *sqlx.DB
}

// SetMaxIdleConns in the db connection
func (s *DBStorage) SetMaxIdleConns(n int) {
	s.db.SetMaxIdleConns(n)
}

// SetMaxOpenConns in the db connection
func (s *DBStorage) SetMaxOpenConns(n int) {
	s.db.SetMaxOpenConns(n)
}

// Ping the db connection
func (s *DBStorage) Ping() bool {
	_, err := s.db.Exec("SELECT 1")

	if err != nil {
		fmt.Errorf("Postgres Ping error: %s", err)
	}

	return err == nil
}

// InsertAggregatedMetric persists the aggregated metric in the db.
func (s *DBStorage) InsertAggregatedMetric(a *models.AggregatedMetric) {
	fmt.Println("Inserting aggregated metric")

	stmt, err := s.db.Prepare(`INSERT INTO metrics(
		service_name, method_name, timestamp, throughput, responseTime, status2XX) VALUES(?, ?, ?, ?, ?)`)
	defer stmt.Close()

	if err != nil {
		fmt.Errorf("Error preparing sql statement: %s", err)
		return
	}

	_, err = stmt.Exec(a.ServiceName, a.MethodName, a.Timestamp, a.Throughput, a.ResponseTime, a.Status2XX)

	if err != nil {
		fmt.Errorf("Error executing sql statement: %s", err)
	}

}

// NewDb creates a new DB connection.
func NewDb(driver, url string) *sqlx.DB {
	db, err := sqlx.Connect(driver, url)

	if err != nil {
		panic(err)
	}

	_, err = createSchemaIfNotExists(db)

	if err != nil {
		panic(err)
	}

	return db
}

func createSchemaIfNotExists(db *sqlx.DB) (sql.Result, error) {
	schema := `CREATE TABLE IF NOT EXISTS metrics (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		service_name varchar NOT NULL,
		method_name varchar NOT NULL,
		timestamp DATETIME NOT NULL,
		throughput BIGINT NOT NULL,
		responseTime BIGINT NOT NULL,
		status2XX BIGINT NOT NULL
	);`

	// execute a query on the server
	return db.Exec(schema)
}

// NewStorage creates a new DB.
func NewStorage(db *sqlx.DB) Storage {
	return &DBStorage{db}
}
