package main

import (
	"fmt"
	"mail-sender/database/enum"
	"mail-sender/database/model"
	"mail-sender/database/table"

	. "github.com/go-jet/jet/v2/mysql"
	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
	"github.com/jmoiron/sqlx"
)

func main() {
	// Database connection
	db, err := sqlx.Open("mysql", "newtendex:devdbaccess@tcp(127.0.0.1:3307)/newtendex?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var mailQueue []model.MailQueue

	// Build a sample SELECT query using Jet
	jet := SELECT(table.MailQueue.AllColumns).FROM(
		table.MailQueue,
	).WHERE(
		table.MailQueue.Status.EQ(enum.MailQueueStatus.Pending),
	)

	// Execute the query and scan the results into a struct
	err = jet.Query(db, &mailQueue)
	if err != nil {
		panic(err)
	}

	// Process the results
	for _, entry := range mailQueue {
		fmt.Printf("ID: %d, Template: %d, Status: %s\n", entry.ID, entry.TemplateID, entry.Status)
	}
}
