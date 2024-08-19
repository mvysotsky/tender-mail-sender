package queries

import "mail-sender/db/model"

type MailQueueEntry model.MailQueue

type MailQueue []MailQueueEntry
