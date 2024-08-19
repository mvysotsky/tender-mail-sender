package main

import (
	"fmt"
	"log"
	"mail-sender/db"
	"mail-sender/db/queries"
	"mail-sender/tools"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

func main() {
	tools.LoadEnv()
	db.Connect()

	queue, err := queries.GetPendingMail()
	if err != nil {
		log.Fatalf("Error getting pending mail: %v", err)
	}

	// Process the results
	for _, entry := range queue {
		fmt.Printf("ID: %d, Template: %d, Status: %s\n", entry.ID, entry.TemplateID, entry.Status)
	}
}
