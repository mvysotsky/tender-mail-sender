package queries

import (
	"mail-sender/db"
	"mail-sender/db/model"
	"mail-sender/db/table"

	"github.com/go-jet/jet/v2/mysql"
)

func GetTemplateByID(templateID uint32) (template model.Templates, err error) {
	err = mysql.SELECT(table.Templates.AllColumns).FROM(
		table.Templates,
	).WHERE(
		table.Templates.ID.EQ(mysql.Int(int64(templateID))),
	).Query(db.Handle, &template)

	return template, err
}
