package controllers

import (
	"database/sql"
	"frog/server/app/models"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	r "github.com/revel/revel"
	"github.com/revel/modules/db/app"
)

var (
	Dbm *gorp.DbMap
)

func InitDB() {
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}

	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}

	t := Dbm.AddTable(models.User{}).SetKeys(true, "No")
	t.ColMap("Password").Transient = true
	t.ColMap("EmailDomain").Transient = true
	setColumnSizes(t, map[string]int{
		"Id":   15,
		"Name": 100,
	})

	t = Dbm.AddTable(models.Company{}).SetKeys(true, "No")
	setColumnSizes(t, map[string]int{
		"Name": 200,
	})

	t = Dbm.AddTable(models.EmailDomain{}).SetKeys(true, "No")
	t.ColMap("Company").Transient = true
	setColumnSizes(t, map[string]int{
		"Domain": 200,
	})

	Dbm.TraceOn("[gorp]", r.INFO)
	Dbm.CreateTablesIfNotExists()
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
