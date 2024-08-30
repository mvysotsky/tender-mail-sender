package repos

import (
	"database/sql"
	"mail-sender/db"
	"mail-sender/db/model"
	"mail-sender/db/table"
	"mail-sender/tools"
	"sync"

	"github.com/go-jet/jet/v2/mysql"
)

type DBTemplates interface {
	GetTemplateByID(templateID uint32) (template string, err error)
}

type dbTemplates struct {
	dbHandle *sql.DB
	cache    map[uint32]string
	mutex    sync.RWMutex
}

func Templates() DBTemplates {
	return &dbTemplates{
		dbHandle: db.Handle,
	}
}

func (r *dbTemplates) GetTemplateByID(templateID uint32) (string, error) {
	r.mutex.RLock()

	if template, ok := r.cache[templateID]; ok {
		r.mutex.RUnlock()
		return template, nil
	}

	r.mutex.RUnlock()
	r.mutex.Lock()
	defer r.mutex.Unlock()

	var tplData model.Templates

	if err := mysql.
		SELECT(table.Templates.AllColumns).
		FROM(table.Templates).
		WHERE(table.Templates.ID.EQ(mysql.Int(int64(templateID)))).
		Query(r.dbHandle, &tplData); err != nil {
		return "", err
	}

	r.cache[templateID] = tools.GetMustacheTemplate(*tplData.Template)

	return r.cache[templateID], nil
}
