package main

import (
	"fmt"
	"github.com/awaisamjad/db/Type"
	"github.com/awaisamjad/db/column"
	"github.com/awaisamjad/db/db"
	"github.com/awaisamjad/db/table"
)

func main() {
	table1 := table.Table{
		Name: "Table 1",
		Id:   1,
		Columns: []column.Column{
			{Name: "Date", Id: 1, Type: Type.INTEGER, Values: []int{1, 2, 34324321432, 4, 5, 6}},
			{Name: "Time", Id: 2, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
			{Name: "£", Id: 3, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
		},
	}

	table2 := table.Table{
		Name: "Table 2",
		Id:   1,
		Columns: []column.Column{
			{Name: "Date", Id: 1, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
			{Name: "Age", Id: 1, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
			{Name: "£", Id: 1, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
		},
	}

	db := db.DB{
		Name:   "Database",
		Tables: []table.Table{table1, table2},
	}

	// fmt.Println(table1.toString())
	fmt.Println(db.toString())
}
