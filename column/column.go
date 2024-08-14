package column

import (
	"github.com/awaisamjad/db/Type"
	"strconv"
)

type Column struct {
	Header   string
	Id     int
	Type   Type.Type
	Values []int
}

func (c *Column) ToString() string {
	values := ""
	for i := 0; i < len(c.Values); i++ {
		// Converts int to string
		values += strconv.Itoa(c.Values[i])
		if i < len(c.Values)-1 {
			values += "\n"
		}
	}
	return c.Header + "\n" + values
}

func (c *Column) GetHeader() string {
	return c.Header
}