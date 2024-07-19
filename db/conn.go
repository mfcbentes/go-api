package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	dbHost     = "localhost" //os.Getenv("DB_HOST")
	dbUser     = "postgres"  //os.Getenv("DB_USER")
	dbPort     = "5432"      //os.Getenv("DB_PORT")
	dbPassword = "1234"      //os.Getenv("DB_PASSWORD")
	dbName     = "go_db"     //os.Getenv("DB_NAME")
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbName)

	return db, nil
}
