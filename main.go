package main

import (
	"fmt"
	"mail-sender/db"
	"mail-sender/db/enum"
	"mail-sender/db/model"
	"mail-sender/db/table"
	"mail-sender/tools"

	. "github.com/go-jet/jet/v2/mysql"
	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

func main() {
	tools.LoadEnv()
	db.Connect()

	var mailQueue []model.MailQueue

	// Build a sample SELECT query using Jet
	jet := SELECT(table.MailQueue.AllColumns).FROM(
		table.MailQueue,
	).WHERE(
		table.MailQueue.Status.EQ(enum.MailQueueStatus.Pending),
	)

	// Execute the query and scan the results into a struct
	err := jet.Query(db.Handle, &mailQueue)
	if err != nil {
		panic(err)
	}

	// Process the results
	for _, entry := range mailQueue {
		fmt.Printf("ID: %d, Template: %d, Status: %s\n", entry.ID, entry.TemplateID, entry.Status)
	}
}
