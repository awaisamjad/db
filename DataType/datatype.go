package DataType

type DataType string

const (
	INTEGER DataType = "INTEGER"
	FLOAT   DataType = "FLOAT"
	STRING  DataType = "STRING"
	CHAR    DataType = "CHAR"
	UNKNOWN DataType = ""
)

type AllowedTypes interface {
	int | string | rune | float64
}