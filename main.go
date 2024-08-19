package main

import (
	"fmt"
	"log"
	"mail-sender/db"
	"mail-sender/db/model"
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

func send(entry model.MailQueue) {
	template, err := queries.GetTemplateByID(entry.TemplateID)
	if err != nil {
		log.Fatalf("Error getting template: %v", err)
	}

	// Send the email
	fmt.Printf("Sending email to %s using template %d\n", entry.Email, template.ID)
}
