package repos

import (
	"database/sql"
	"mail-sender/db"
	"mail-sender/db/model"
	"mail-sender/db/table"

	"github.com/go-jet/jet/v2/mysql"
)

type DBTemplates interface {
	GetTemplateByID(templateID uint32) (template model.Templates, err error)
}

type dbTemplates struct {
	dbHandle *sql.DB
}

func Templates() DBTemplates {
	return dbTemplates{
		dbHandle: db.Handle,
	}
}

func (r dbTemplates) GetTemplateByID(templateID uint32) (template model.Templates, err error) {
	err = mysql.
		SELECT(table.Templates.AllColumns).
		FROM(table.Templates).
		WHERE(table.Templates.ID.EQ(mysql.Int(int64(templateID)))).
		Query(r.dbHandle, &template)

	return template, err
}
