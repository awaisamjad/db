package column

import (
	"fmt"
	// "strconv"

	"github.com/awaisamjad/db/DataType"
	"github.com/awaisamjad/db/utils"

)

type Column[T DataType.AllowedTypes] struct {
	Header string
	Id     int
	Type   DataType.DataType
	Values []T
}

func New [T DataType.AllowedTypes ] (header string, id int, Type DataType.DataType, values []T) (*Column[T], error) {
	column := &Column[T]{
		Header: header,
		Id: id,
		Type: Type,
		Values: values,
	}
	
	err := column.Validate()

	if err != nil {
		return nil, fmt.Errorf("error creating column: %v", err)
	}

	return column, nil
}



func (c *Column[T]) ToString() string {
	values := ""
	for i := 0; i < len(c.Values); i++ {
		values += fmt.Sprintf("%v", c.Values[i])
		if i < len(c.Values)-1 {
			values += "\n"
		}
	}
	return c.Header + "\n" + values
}

func (c *Column[T]) GetHeader() string {
	return c.Header
}

func (c *Column[T]) Validate() error {

	if len(c.Values) < 1 {
		return fmt.Errorf("invalid length of column")
	}

	isTypeValid := utils.IsTypeForValuesValid(c.Type, c.Values)	
	if isTypeValid != nil {
		return fmt.Errorf("DataType Field and type for Values Field do not match")
	}


	return nil
}
