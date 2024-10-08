package main

import (
	// "fmt"

	"github.com/awaisamjad/db/Type"
	"github.com/awaisamjad/db/column"

	// "github.com/awaisamjad/db/db"
	"github.com/awaisamjad/db/table"
)

func main() {
	table1 := table.Table{
		Name: "Table 1",
		Id:   1,
		Columns: []column.Column{
			{Header: "Date", Id: 1, Type: Type.INTEGER, Values: []int{1, 2, 34324321432, 4, 5, 6}},
			{Header: "Time", Id: 2, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
			{Header: "£", Id: 3, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
			{Header: "Weather", Id: 3, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
		},
	}

	table2 := table.Table{
		Name: "Table 2",
		Id:   1,
		Columns: []column.Column{
			{Header: "Date", Id: 1, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
			{Header: "Age", Id: 1, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
			{Header: "£", Id: 1, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
			{Header: "Weather", Id: 3, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
		},
	}

	// db := db.DB{
	// 	Name:   "Database",
	// 	Tables: []table.Table{table1, table2},
	// }

	table1.InnerJoin(table2, "Date")
	// fmt.Println(db.ToString())
}
