package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	connectionString := getConnectionString()
	fmt.Println(connectionString)
}

func getConnectionString() string {
	bytes, err := os.ReadFile("../secrets/connectionString.txt")
	if err != nil {
		log.Fatal()
	}
	return string(bytes)
}
