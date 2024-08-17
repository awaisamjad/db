package column

import (
	"fmt"
	// "strconv"

	"github.com/awaisamjad/db/Type"
)

type Column struct {
	Header string
	Id     int
	Type   Type.Type
	Values interface{}
}

func New(header string, id int, Type Type.Type, values )

func (c *Column) ToString() string {
	// values := ""
	// for i := 0; i < len(c.Values); i++ {
	// 	// Converts int to string
	// 	values += strconv.Itoa(c.Values[i])
	// 	if i < len(c.Values)-1 {
	// 		values += "\n"
	// 	}
	// }
	// return c.Header + "\n" + values
	return fmt.Sprintf("Header: %s, Id: %d, Type: %s, Values: %v", c.Header, c.Id, c.Type, c.Values)

}

func (c *Column) GetHeader() string {
	return c.Header
}

func (c *Column) Validate() error {
	switch c.Type {
	case Type.INTEGER:
		if _, ok := c.Values.([]int); !ok {
			return fmt.Errorf("values must be of type []int for INTEGER column")
		}
	case Type.FLOAT:
		if _, ok := c.Values.([]float64); !ok {
			return fmt.Errorf("values must be of type []float64 for FLOAT column")
		}
	case Type.STRING:
		if _, ok := c.Values.([]string); !ok {
			return fmt.Errorf("values must be of type []string for STRING column")
		}
	case Type.CHAR:
		if _, ok := c.Values.([]rune); !ok {
			return fmt.Errorf("values must be of type []rune for CHAR column")
		}
	default:
		return fmt.Errorf("unknown column type")
	}
	return nil
}
