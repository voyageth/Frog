package controllers

import (
	"database/sql"
	"frog/server/app/models"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	r "github.com/revel/revel"
	"github.com/revel/modules/db/app"
	"log"
)

var (
	Dbm *gorp.DbMap
)

func InitDB() {
	log.Println("InitDB start")

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
	t.ColMap("Name").Unique = true
	setColumnSizes(t, map[string]int{
		"Id":   200,
		"Name": 200,
	})

	t = Dbm.AddTable(models.Company{}).SetKeys(true, "No")
	t.ColMap("Name").Unique = true
	setColumnSizes(t, map[string]int{
		"Name": 200,
	})

	t = Dbm.AddTable(models.EmailDomain{}).SetKeys(true, "No")
	t.ColMap("Domain").Unique = true
	t.ColMap("Company").Transient = true
	setColumnSizes(t, map[string]int{
		"Domain": 200,
	})

	Dbm.TraceOn("[gorp]", r.INFO)
	log.Println(Dbm.CreateTablesIfNotExists())

	log.Println("InitDB end")
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
