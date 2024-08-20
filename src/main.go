package main

import (
	"fmt"

	"github.com/awaisamjad/db/DataType"
	"github.com/awaisamjad/db/column"
	"github.com/awaisamjad/db/table"
	// "github.com/awaisamjad/db/db"
	// "github.com/awaisamjad/db/table"
)

func main() {
	fmt.Println()

	column1, _ := column.New[int](
		"Column 1",
		1,
		DataType.INTEGER,
		[]int{1, 2, 4, 5},
	)
	column2, _ := column.New[string](
		"Column 2",
		2,
		DataType.STRING,
		[]string{"First", "Second", "Third", "Fourth"},
	)
	column3, _ := column.New[rune](
		"Column 3",
		3,
		DataType.CHAR,
		[]rune{'a', 'b', 'c', 'd'},
	)
	column4, _ := column.New[float64](
		"Column 4",
		4,
		DataType.FLOAT,
		[]float64{1.1, 2.2, 4.3, 5.7},
	)

	columns := [DataType.AllowedTypes]interface{}{
		column1,
		column2,
		column3,
		column4,
	}

	table1, err := table.New("Table 1", 1, columns)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	fmt.Println("Table created:", table1.ToString())

}
