package repos

import (
	"database/sql"
	"mail-sender/db"
	"mail-sender/db/model"
	"mail-sender/db/table"
	"sync"

	"github.com/go-jet/jet/v2/mysql"
)

type Tender struct {
	model.Tenders
	model.Categories
}

type DBTenders interface {
	GetTender(tenderID uint32) (tender Tender, err error)
}

type dbTenders struct {
	dbHandle *sql.DB
	cache    map[uint32]Tender
	mutex    sync.RWMutex
}

func Tenders() DBTenders {
	return &dbTenders{
		dbHandle: db.Handle,
	}
}

func (r *dbTenders) GetTender(tenderID uint32) (tender Tender, err error) {
	r.mutex.RLock()

	if tender, ok := r.cache[tenderID]; ok {
		r.mutex.RUnlock()
		return tender, nil
	}

	r.mutex.RUnlock()
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if err = mysql.
		SELECT(table.Tenders.AllColumns, table.Categories.AllColumns).
		FROM(table.Tenders.
			LEFT_JOIN(table.Categories, table.Tenders.CategoryID.EQ(table.Categories.ID))).
		WHERE(table.Tenders.ID.EQ(mysql.Int(int64(tenderID)))).
		Query(db.Handle, &tender); err != nil {
		return Tender{}, err
	}

	r.cache[tenderID] = tender

	return tender, nil
}
