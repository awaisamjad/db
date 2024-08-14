package table

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/awaisamjad/db/Type"
	"github.com/awaisamjad/db/column"
	"github.com/awaisamjad/db/utils"
)

type Table struct {
	Name    string
	Id      int
	Columns []column.Column
}

func (t *Table) ToString() string {
	table := ""
	max_lengths_of_each_column := make([]int, len(t.Columns))

	// Gets the max len of all values in the column inc name
	for i := 0; i < len(t.Columns); i++ {
		// Make the default max length the name of column. if any values len is greater change variable value
		max_length_value := t.Columns[i].Header
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
		table += t.Columns[i].Header + utils.Gap(uint32(max_lengths_of_each_column[i]-len(t.Columns[i].Header)+3))
	}
	table += "\n" // New line after column names

	// Add values of each column to the table string
	width := 3
	for i := 0; i < len(t.Columns[0].Values); i++ {
		for j := 0; j < len(t.Columns); j++ {
			value := strconv.Itoa(t.Columns[j].Values[i])
			table += value + utils.Gap(uint32(max_lengths_of_each_column[j]-len(value)+width))
		}
		table += "\n" // New line after each row
	}

	return t.Name + "\n\n" + table
}

func (t *Table) GetHeaders() []string {
	headers := make([]string, len(t.Columns))
	for i := 0; i < len(t.Columns); i++ {
		headers = append(headers, t.Columns[i].GetHeader())
	}
	return headers
}

func (table1 *Table) InnerJoin(table2 Table, header string) (Table, error) {
	//TODO make sure the size of the tables matches
	if !utils.Contains(table1.GetHeaders(), header) && !utils.Contains(table2.GetHeaders(), header) {
		//TODO fix this return of table
		return Table{
			Name: "Table 2",
			Id:   1,
			Columns: []column.Column{
				{Header: "Date", Id: 1, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
				{Header: "Age", Id: 1, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
				{Header: "£", Id: 1, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
			},
		}, errors.New("headers dont match")
	}

	same_values := utils.SameValues(table1.GetHeaders(), table2.GetHeaders())
	fmt.Println(same_values)

	return Table{
		Name: "Table 2",
		Id:   1,
		Columns: []column.Column{
			{Header: "Date", Id: 1, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
			{Header: "Age", Id: 1, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
			{Header: "£", Id: 1, Type: Type.INTEGER, Values: []int{1, 2, 3, 4, 5, 6}},
		},
	}, nil

}
