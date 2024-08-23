package queries

import (
	"mail-sender/db"
	"mail-sender/db/model"
	"mail-sender/db/table"

	"github.com/go-jet/jet/v2/mysql"
)

type Template model.Templates

func GetTemplateByID(templateID uint32) (template Template, err error) {
	err = mysql.SELECT(table.Templates.AllColumns).FROM(
		table.Templates,
	).WHERE(
		table.Templates.ID.EQ(mysql.Int(int64(templateID))),
	).Query(db.Handle, &template)

	return template, err
}
