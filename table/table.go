package table

import (
	"strconv"

	"github.com/awaisamjad/db/column"
	"github.com/awaisamjad/db/utils"
)

type Table struct {
	Name    string
	Id      int
	Columns []column.Column
}

func (t *Table) toString() string {
	table := ""
	max_lengths_of_each_column := make([]int, len(t.Columns))

	// Gets the max len of all values in the column inc name
	for i := 0; i < len(t.Columns); i++ {
		// Make the default max length the name of column. if any values len is greater change variable value
		max_length_value := t.Columns[i].Name
		for _, v := range t.Columns[i].Values {
			value_string := strconv.Itoa(v)
			if len(value_string) > len(max_length_value) {
				max_length_value = value_string
			}
		}
		max_lengths_of_each_column[i] = len(max_length_value)
	}

	// Add column names to the table string
	for i := 0; i < len(t.Columns); i++ {
		table += t.Columns[i].Name + utils.gap(uint32(max_lengths_of_each_column[i]-len(t.Columns[i].Name)+3))
	}
	table += "\n" // New line after column names

	// Add values of each column to the table string
	width := 3
	for i := 0; i < len(t.Columns[0].Values); i++ {
		for j := 0; j < len(t.Columns); j++ {
			value := strconv.Itoa(t.Columns[j].Values[i])
			table += value + gap(uint32(max_lengths_of_each_column[j]-len(value)+width))
		}
		table += "\n" // New line after each row
	}

	return t.Name + "\n\n" + table
}