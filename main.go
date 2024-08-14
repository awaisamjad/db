// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hello is a simple hello, world demonstration web server.
//
// It serves version information on /version and answers
// any other request like /name by saying "Hello, name!".
//
// See golang.org/x/example/outyet for a more sophisticated server.
package main

import (
	"fmt"
	"strconv"
)

type Type string

const (
	INTEGER Type = "INTEGER_TYPE"
	FLOAT   Type = "FLOAT_TYPE"
	STRING  Type = "STRING_TYPE"
	CHAR    Type = "CHAR_TYPE"
)

type Column struct {
	Name   string
	Id     int
	Type   Type
	Values []int
}

func (c *Column) toString() string {
	values := ""
	for i := 0; i < len(c.Values); i++ {
		// Converts int to string
		values += strconv.Itoa(c.Values[i])
		if i < len(c.Values)-1 {
			values += "\n"
		}
	}
	return c.Name + "\n" + values
}

func (c *Column) getName() string {
	return c.Name
}

type Table struct {
	Name    string
	Id      int
	Columns []Column
}

func gap(width uint32) string {
	gap := ""
	for i := 0; i < int(width); i++ {
		gap += " "
	}
	return gap
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
		table += t.Columns[i].Name + gap(uint32(max_lengths_of_each_column[i]-len(t.Columns[i].Name)+3))
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

type DB struct {
	Name   string
	Tables []Table
}

func (db *DB) toString() string {
	databases := ""
	for i := 0; i < len(db.Tables); i++ {
		databases += db.Tables[i].toString() + "\n"
	}
	return db.Name + "\n\n" + databases
}

func main() {
	table1 := Table{
		Name: "Table 1",
		Id:   1,
		Columns: []Column{
			{"Date", 1, INTEGER, []int{1, 2, 34324321432, 4, 5, 6}},
			{"Time", 2, INTEGER, []int{1, 2, 3, 4, 5, 6}},
			{"£", 3, INTEGER, []int{1, 2, 3, 4, 5, 6}},
		},
	}

	table2 := Table{
		Name: "Table 2",
		Id:   1,
		Columns: []Column{
			{"Date", 1, INTEGER, []int{1, 2, 3, 4, 5, 6}},
			{"Age", 1, INTEGER, []int{1, 2, 3, 4, 5, 6}},
			{"£", 1, INTEGER, []int{1, 2, 3, 4, 5, 6}},
		},
	}

	db := DB{
		Name:   "Database",
		Tables: []Table{table1, table2},
	}

	// fmt.Println(table1.toString())
	fmt.Println(db.toString())
}
