package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func InitializeMySQLConnection(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	// Set additional configuration options if needed

	return db, nil
}

func ExecuteQuery(db *sql.DB, query string) ([]interface{}, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []interface{}
	for rows.Next() {
		// Process and append results
	}

	return results, nil
}
