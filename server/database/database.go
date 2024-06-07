package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"server/database/queries"
	"server/interfaces"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var err error

func connectDatabase() (*sql.DB, error) {
	log.Println("Creating sqlite-database.db...")
	file, fError := os.Create("database/sqlite-database.db") // Create SQLite file
	if fError != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")
	db, err = sql.Open("sqlite3", "database/sqlite-database.db")
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to SQLite")
	return db, nil
}

func execTableStatement() error {
	statement, err := db.Prepare(queries.ExecCreateTableQuery())
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	return nil
}

func StoreData(data *interfaces.ExchangeRateResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("Persisting data")
	statement, err := db.Prepare(queries.ExecInserIntoQuery(data.Bid))
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = statement.ExecContext(ctx)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	return nil
}

func GetAllRecords() ([]interfaces.ExchangeRecord, error) {
	var records []interfaces.ExchangeRecord
	row, err := db.Query("SELECT * FROM exchange")
	if err != nil {
		return []interfaces.ExchangeRecord{}, err
	}
	defer row.Close()
	for row.Next() {
		item := interfaces.ExchangeRecord{}
		err := row.Scan(&item.Id, &item.CreatedAt, &item.Value)
		if err != nil {
			return []interfaces.ExchangeRecord{}, err
		}
		records = append(records, item)
	}

	return records, nil
}

func SetupDatabase() error {
	_, err := connectDatabase()
	if err != nil {
		panic(err)
	}
	err = execTableStatement()
	if err != nil {
		return err
	}

	return nil
}
