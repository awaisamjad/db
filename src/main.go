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
			Header: "Column 2",
			Id:     2,
			Type:   Type.CHAR,
			Values: []rune{'a'},
		},
		{
			Header: "Column 3",
			Id:     3,
			Type:   Type.INTEGER,
			Values: []int{6, 7, 8, 9, 10},
		},
		{
			Header: "Column 4",
			Id:     4,
			Type:   Type.FLOAT,
			Values: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
		},
	})
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	fmt.Println("Table created:", table1.ToString())

}
