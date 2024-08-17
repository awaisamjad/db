package main

import (
	"fmt"
	"github.com/awaisamjad/db/DataType"
	"github.com/awaisamjad/db/column"
	// "github.com/awaisamjad/db/db"
	// "github.com/awaisamjad/db/table"
)

func main() {
	fmt.Println()

	column1, err := column.New(
		"Column 1",
		1,
		DataType.INTEGER,
		[]int{1,2,4,5},
	)

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	fmt.Printf("Created column: %+v\n", column1)

	// table1, err := table.New("Table1", 1, []column.Column{
	// 	{
	// 		Header: "Column 1",
	// 		Id:     1,
	// 		Type:   DataType.INTEGER,
	// 		Values: []int{1, 2, 3, 4, 5},
	// 	},
	// 	{
	// 		Header: "Column 2",
	// 		Id:     2,
	// 		Type:   DataType.INTEGER,
	// 		Values: []int{10,20,30,40,50,60},
	// 	},
	// 	{
	// 		Header: "Column 3",
	// 		Id:     3,
	// 		Type:   DataType.INTEGER,
	// 		Values: []int{6, 7, 8, 9, 10},
	// 	},
	// 	{
	// 		Header: "Column 4",
	// 		Id:     4,
	// 		Type:   DataType.FLOAT,
	// 		Values: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
	// 	},
	// })
	// if err != nil {
	// 	fmt.Println("Error creating table:", err)
	// 	return
	// }

	// fmt.Println("Table created:", table1.ToString())

}
