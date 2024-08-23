package queries

import (
	"mail-sender/db"
	"mail-sender/db/model"
	"mail-sender/db/table"

	"github.com/go-jet/jet/v2/mysql"
)

type User struct {
	model.Users
	model.Accounts
}

func GetUser(userID uint32) (user User, err error) {
	err = mysql.SELECT(table.Users.AllColumns, table.Accounts.AllColumns).FROM(
		table.Users,
	).WHERE(
		table.Users.ID.EQ(mysql.Int(int64(userID))),
	).Query(db.Handle, &user)

	return user, err
}
