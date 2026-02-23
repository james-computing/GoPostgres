package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	connectionString := getConnectionString()
	fmt.Println(connectionString)

	dbConnection, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Test connection to database
	err = dbConnection.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer dbConnection.Close()
}

func getConnectionString() string {
	bytes, err := os.ReadFile("../secrets/connectionString.txt")
	if err != nil {
		log.Fatal()
	}
	return string(bytes)
}
