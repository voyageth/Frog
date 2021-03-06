package controllers

import (
	"database/sql"
	"github.com/voyageth/frog/server/app/models"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	r "github.com/revel/revel"
	"github.com/revel/modules/db/app"
)

var (
	Dbm *gorp.DbMap
)

func InitDB() {
	r.INFO.Println("InitDB start")

	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}

	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}

	t := Dbm.AddTable(models.User{}).SetKeys(true, "No")
	t.ColMap("Password").Transient = true
	t.ColMap("EmailDomain").Transient = true
	t.ColMap("Email").SetUnique(true)
	t.ColMap("Email").SetNotNull(true)
	t.ColMap("Name").SetUnique(true)
	t.ColMap("Name").SetNotNull(true)
	setColumnSizes(t, map[string]int{
		"Email":   200,
		"Name": 200,
	})

	t = Dbm.AddTable(models.Company{}).SetKeys(true, "No")
	t.ColMap("Name").SetUnique(true)
	t.ColMap("Name").SetNotNull(true)
	setColumnSizes(t, map[string]int{
		"Name": 200,
	})

	t = Dbm.AddTable(models.EmailDomain{}).SetKeys(true, "No")
	t.ColMap("Company").Transient = true
	t.ColMap("Domain").SetUnique(true)
	t.ColMap("Domain").SetNotNull(true)
	setColumnSizes(t, map[string]int{
		"Domain": 200,
	})

	Dbm.TraceOn("[gorp]", r.TRACE)
	r.INFO.Println(Dbm.CreateTablesIfNotExists())

	r.INFO.Println("InitDB end")
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
