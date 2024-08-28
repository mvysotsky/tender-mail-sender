package main

import (
	"mail-sender/db"
	"mail-sender/service"
	"mail-sender/tools"
)

func main() {
	tools.LoadEnv()
	db.Connect()

	mailSender := service.NewMailSender()
	if err := mailSender.SendEmailNotifications(); err != nil {
		panic(err)
	}
}
