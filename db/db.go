package db

import "github.com/awaisamjad/db/table"

type DB struct {
	Name   string
	Tables []table.Table
}

func (db *DB) toString() string {
	databases := ""
	for i := 0; i < len(db.Tables); i++ {
		databases += db.Tables[i].toString() + "\n"
	}
	return db.Name + "\n\n" + databases
}