package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectToDatabase(uri string) (*sql.DB, error) {
	db, err := psql.Connect(psql.WithConnectionWaiting())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Successfully connected!")
}
