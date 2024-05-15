package go_migrator

import "fmt"

type ColumnType int32

const (
	ColumnTypeSerial    ColumnType = 1
	ColumnTypeString    ColumnType = 2
	ColumnTypeInteger   ColumnType = 3
	ColumnTypeUUID      ColumnType = 4
	ColumnTypeBoolean   ColumnType = 5
	ColumnTypeFloat     ColumnType = 6
	ColumnTypeDouble    ColumnType = 7
	ColumnTypeDecimal   ColumnType = 8
	ColumnTypeTimestamp ColumnType = 9
)

func (c ColumnType) toString(length int32) string {
	switch c {
	case ColumnTypeSerial:
		return "SERIAL"
	case ColumnTypeString:
		return "VARCHAR(" + fmt.Sprintf("%d", length) + ")"
	case ColumnTypeUUID:
		return "UUID"
	case ColumnTypeInteger:
		return "INT"
	case ColumnTypeBoolean:
		return "BOOLEAN"
	case ColumnTypeFloat:
		return "FLOAT"
	case ColumnTypeDouble:
		return "DOUBLE"
	case ColumnTypeTimestamp:
		return "TIMESTAMP"
	}

	return ""
}
