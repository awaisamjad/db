package db

import "github.com/awaisamjad/db/table"

type DB struct {
	Name   string
	Tables []table.Table
}

func (db *DB) ToString() string {
	databases := ""
	for i := 0; i < len(db.Tables); i++ {
		databases += db.Tables[i].ToString() + "\n"
	}
	return db.Name + "\n\n" + databases
}