package queries

import (
	"mail-sender/db"
	"mail-sender/db/model"
	"mail-sender/db/table"

	"github.com/go-jet/jet/v2/mysql"
)

type Tender struct {
	model.Tenders
	Name string
}

func GetTender(tenderID uint32) (tender Tender, err error) {
	err = mysql.SELECT(table.Tenders.AllColumns, table.Categories.Name).FROM(
		table.Tenders.
			LEFT_JOIN(table.Categories, table.Tenders.CategoryID.EQ(table.Categories.ID)),
	).WHERE(
		table.Tenders.ID.EQ(mysql.Int(int64(tenderID))),
	).Query(db.Handle, &tender)

	return tender, err
}
