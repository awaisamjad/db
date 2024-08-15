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

func New(name string, id int, columns []column.Column) (*Table, error) {
	for _, col := range columns {
		if err := col.Validate(); err != nil {
			return nil, fmt.Errorf("invalid column %s: %v", col.Header, err)
		}
	}
	return &Table{
		Name:    name,
		Id:      id,
		Columns: columns,
	}, nil
}

// This is done as we cant fill empty data values with nil so these are the nil replacements
// ! This means these values cant really be used as they may mean nil or 0/0.0/""
// TODO Fix
func getZeroValue(t Type.Type) interface{} {
	switch t {
	case Type.INTEGER:
		return 0
	case Type.FLOAT:
		return 0.0
	case Type.STRING, Type.CHAR:
		return ""
	default:
		return nil
	}
}

// ~ Performs necessary checks on a table
func (t *Table) Validate() error {
	//? Makes sure none of the fields are empty
	if t.Name == "" || strconv.Itoa(t.Id) == "" || len(t.Columns) == 0 {
		return fmt.Errorf("fields arent correct")
	}

	//? Makes sure the Values are the correct type i.e. they match the Type field
	// for i := 0; i < len(t.Columns); i++ {
	// 	for j := 0; j < len(t.Columns[i].Values); j++ {
	// 		value := t.Columns[i].Values[j]
	// 		if utils.CheckValueType(value) != t.Columns[i].Type {
	// 			return fmt.Errorf("The values arent all the same type")
	// 		}
	// 	}
	// }

	//? Makes sure all the columns have same number of values in the column
	//~ If not it fills them with nil values

	// maxLength := len(t.Columns[0].Values)
	// for _, col := range t.Columns {
	// 	if len(col.Values) > maxLength {
	// 		maxLength = len(col.Values)
	// 	}
	// }

	// Pad columns with nil values if they are shorter than maxLength
	//TODO finish this. problem with zeroValue
	// for i := range t.Columns {
	// 	zeroValue := getZeroValue(t.Columns[i].Type)
	// 	for len(t.Columns[i].Values) < maxLength {
	// 		t.Columns[i].Values = append(t.Columns[i].Values, zeroValue)
	// 	}
	// }

	return nil
}

func (t *Table) AddColumn(c column.Column) {
	t.Columns = append(t.Columns, c)
}

func (t *Table) AddColumns(c []column.Column) {
	for i := 0; i < len(c); i++ {
		t.Columns = append(t.Columns, c[i])
	}
}

func (t *Table) AddRow(row []string) error {
	if len(t.Columns) != len(row) {
		return fmt.Errorf("Length of rows doesnt match number of columns")
	}

	// for i := 0; i < len(t.Columns); i++ {
	// 	t.Columns[i].Values = append(t.Columns[i].Values, row[i].toValueType())
	// }

	return nil
}

func (t *Table) ToString() string {
	table := ""
	max_lengths_of_each_column := make([]int, len(t.Columns))

	//? Gets the max len of all values in the column inc name
	for i := 0; i < len(t.Columns); i++ {
		// Make the default max length the name of column. if any values len is greater change variable value
		max_length_value := t.Columns[i].Header
	
		switch t.Columns[i].Type {
		case Type.INTEGER:
			values, ok := t.Columns[i].Values.([]int)
			if !ok {
				//TODO Handle the error case where Values is not of type []int
				continue
			}
			for _, v := range values {
				value_string := strconv.Itoa(v)
				if len(value_string) > len(max_length_value) {
					max_length_value = value_string
				}
				fmt.Println(value_string)
			}
		case Type.STRING:
			values, ok := t.Columns[i].Values.([]string)
			if !ok {
				// Handle the error case where Values is not of type []string
				continue
			}
			for _, v := range values {
				if len(v) > len(max_length_value) {
					max_length_value = v
				}
				fmt.Println(v)
			}
		//? As CHAR is one character it doesnt need to be handled
		//! Assumption
		//TODO Handle just in case
		// case Type.CHAR:
		// 	values, ok := t.Columns[i].Values.([]rune)
		// 	if !ok {
		// 		// Handle the error case where Values is not of type []string
		// 		continue
		// 	}
		// 	for _, v := range values {
		// 		if len(v) > len(max_length_value) {
		// 			max_length_value = v
		// 		}
		// 		fmt.Println(v)
		// 	}
		case Type.FLOAT:
			values, ok := t.Columns[i].Values.([]float64)
			if !ok {
				// Handle the error case where Values is not of type []string
				continue
			}
			for _, v := range values {
				//! No idea how the FormatFloat function works
				value_string := strconv.FormatFloat(v, 'f', -1, 64)
				if len(value_string) > len(max_length_value) {
					max_length_value = value_string
				}
				fmt.Println(value_string)
			}
		default:
			fmt.Println("Unsupported column type")
		}
	
		max_lengths_of_each_column[i] = len(max_length_value)
	}

	//? Add column names to the table string
	for i := 0; i < len(t.Columns); i++ {
		table += t.Columns[i].Header + utils.Gap(uint32(max_lengths_of_each_column[i]-len(t.Columns[i].Header)+3))
	}
	table += "\n" // New line after column names

	// Add values of each column to the table string
	width := 3
	//* this should work as the length of all columns should be the same
	//! Check if this check does occur before
	len_of_values := len(t.Columns[0].Values.([]float64))

	for i := 0; i < len_of_values; i++ {
		for j := 0; j < len(t.Columns); j++ {
			// Get the length of each value. As it is an interface it requires a case switch on its type
			values, ok := t.Columns[j].Values.([]interface{})
			if !ok {
				// Handle the error case where Values is not of type []interface{}
				continue
			}
			value := values[i]
			var value_conv string
			switch v := value.(type) {
			case int:
				value_conv = strconv.Itoa(v)
			case float64:
				value_conv = strconv.FormatFloat(v, 'f', -1, 64)
			case string:
				value_conv = v
			default:
				value_conv = ""
			}
			table += value_conv + utils.Gap(uint32(max_lengths_of_each_column[j]-len(value_conv)+width))
		}
		table += "\n" // New line after each row
	}

	return t.Name + "\n\n" + table
	// return fmt.Sprintf("Table Name: %s, Id: %d, Columns: %v", t.Name, t.Id, t.Columns)
}



func (t *Table) GetHeaders() []string {
	headers := make([]string, len(t.Columns))
	for i := 0; i < len(t.Columns); i++ {
		headers = append(headers, t.Columns[i].GetHeader())
	}
	return headers
}

// * Join functions
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
