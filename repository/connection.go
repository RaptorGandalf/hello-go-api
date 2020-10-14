package repository

import (
	"github.com/jinzhu/gorm"
	// for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Database - Provides interaction with a database
type Database struct {
	Connection *gorm.DB
	SampleRepo SampleRepository
}

// GetDatabase - Returns a database struct used for interacting with the database
func GetDatabase() (Database, error) {
	connection, err := GetConnection()
	if err != nil {
		return Database{}, err
	}

	database := buildDatabase(connection)

	return database, nil
}

// GetDatabaseForConnection - Returns a database struct created with the specified connection
func GetDatabaseForConnection(connection *gorm.DB) Database {
	return buildDatabase(connection)
}

// GetConnection - Returns a database connection
func GetConnection() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "postgres://postgres:postgres@localhost:5432/sample?sslmode=disable")
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(50)

	return db, nil
}

func buildDatabase(connection *gorm.DB) Database {
	return Database{
		Connection: connection,
		SampleRepo: GetSampleRepository(connection),
	}
}
