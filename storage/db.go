package storage

import (
	"database/sql"
	"fmt"
	"time"

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

	stmt, _ := s.db.Prepare(`INSERT INTO metrics(
		serviceName, methodName, timestamp, throughput, responseTime, status2XX) VALUES(?, ?, ?, ?, ?, ?)`)
	defer stmt.Close()

	_, err := stmt.Exec(a.ServiceName, a.MethodName, a.Timestamp, a.Throughput, a.ResponseTime, a.Status2XX)

	if err != nil {
		fmt.Errorf("Error executing sql statement: %s", err)
	}
}

// FindMethodMetrics returns metrics since the given time (grouped by service.method).
func (s *DBStorage) FindMethodMetrics(since time.Time) ([]*models.ServiceMethodMetrics, error) {
	metrics := []*models.ServiceMethodMetrics{}

	err := s.db.Select(&metrics, `SELECT serviceName, methodName,
		MIN(throughput) as minThroughput, MAX(throughput) as maxThroughput, ROUND(AVG(throughput)) as avgThroughput,
		MIN(responseTime) as minResponseTime, MAX(responseTime) as maxResponseTime, ROUND(AVG(responseTime)) as avgResponseTime,
		MIN(status2XX) as minStatus2XX, MAX(status2XX) as maxStatus2XX, ROUND(AVG(status2XX)) as avgStatus2XX
		FROM metrics
		WHERE timestamp >= ?
		GROUP BY serviceName, methodName
	`, since)

	if err != nil {
		return nil, err
	}

	for _, m := range metrics {
		err := s.db.Select(&m.History, `
		SELECT
			serviceName,
			methodName,
			ROUND(AVG(throughput)) as throughput,
			ROUND(AVG(responseTime)) as responseTime 
		FROM metrics 
        WHERE timestamp >= ? AND serviceName = ? AND methodName = ?
		GROUP BY
			serviceName,
			methodName,
			strftime('%Y%m%d%H0', timestamp) + strftime('%M', timestamp)/?
		`, since, m.ServiceName, m.MethodName, 5)

		if err != nil {
			return nil, err
		}
	}

	return metrics, nil
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
		serviceName varchar NOT NULL,
		methodName varchar NOT NULL,
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
