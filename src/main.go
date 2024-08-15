package main

import (
	"fmt"
	"github.com/awaisamjad/db/Type"
	"github.com/awaisamjad/db/column"
	// "github.com/awaisamjad/db/db"
	"github.com/awaisamjad/db/table"
)

func main() {
	fmt.Println()

	table1, err := table.New("Table1", 1, []column.Column{
		{
			Header: "Column 1",
			Id:     1,
			Type:   Type.INTEGER,
			Values: []int{1, 2, 3, 4, 5},
		},
		{
			Header: "Column 1",
			Id:     1,
			Type:   Type.CHAR,
			Values: []rune{"a"},
		}, {
			Header: "Column 1",
			Id:     1,
			Type:   Type.INTEGER,
			Values: []int{1, 2, 3, 4, 5},
		}, {
			Header: "Column 1",
			Id:     1,
			Type:   Type.INTEGER,
			Values: []int{1, 2, 3, 4, 5},
		}})
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	fmt.Println("Table created:", table1.ToString())

}
