package go_migrator

import (
	"fmt"
	"log"
)

type columnType int32

const (
	smallInt   columnType = 1
	integer    columnType = 2
	bigInt     columnType = 3
	serial     columnType = 4
	bigSerial  columnType = 5
	numeric    columnType = 6
	real       columnType = 7
	double     columnType = 8
	varChar    columnType = 9
	char       columnType = 10
	text       columnType = 11
	date       columnType = 12
	dateTime   columnType = 13
	timestamp  columnType = 14
	timestampz columnType = 15
	interval   columnType = 16
	money      columnType = 17
	uuid       columnType = 18
	json       columnType = 19
	jsonb      columnType = 20
	array      columnType = 21
	point      columnType = 22
	line       columnType = 23
	lseg       columnType = 24
	box        columnType = 25
	path       columnType = 26
	polygon    columnType = 27
	circle     columnType = 28
	cidr       columnType = 29
	inet       columnType = 30
	macAddr    columnType = 31
	macAddr8   columnType = 32
	bit        columnType = 33
	bitVar     columnType = 34
	xml        columnType = 35
	hStore     columnType = 36
	int4range  columnType = 37
	int8range  columnType = 38
	numRange   columnType = 39
	tsRange    columnType = 40
	tstzRange  columnType = 41
	dateRange  columnType = 42
	boolean    columnType = 43
)

const (
	ArrayTypeInteger   = 1
	ArrayTypeVarchar   = 2
	ArrayTypeText      = 3
	ArrayTypeJson      = 4
	ArrayTypeNumeric   = 5
	ArrayTypeDate      = 6
	ArrayTypeTimestamp = 7
	ArrayTypeBoolean   = 8
)

func (c columnType) toString(length int32, arrayType *columnType) string {
	switch c {
	case smallInt:
		return "SMALLINT"
	case integer:
		return "INTEGER"
	case bigInt:
		return "BIGINT"
	case serial:
		return "SERIAL"
	case bigSerial:
		return "BIGSERIAL"
	case numeric:
		return "NUMERIC"
	case real:
		return "REAL"
	case double:
		return "DOUBLE"
	case varChar:
		return "VARCHAR(" + fmt.Sprintf("%d", length) + ")"
	case char:
		return "CHAR(" + fmt.Sprintf("%d", length) + ")"
	case text:
		return "TEXT"
	case date:
		return "DATE"
	case dateTime:
		return "DATETIME"
	case timestamp:
		return "TIMESTAMPT"
	case timestampz:
		return "TIMESTAMPTZ"
	case interval:
		return "INTERVAL"
	case money:
		return "MONEY"
	case uuid:
		return "UUID"
	case json:
		return "JSON"
	case jsonb:
		return "JSONB"
	case array:
		if arrayType == nil {
			log.Fatal("Invalid operation")
		}

		switch *arrayType {
		case varChar:
			return fmt.Sprintf("VARCHAR(%d)[]", length)
		case integer:
			return "INTEGER[]"
		case text:
			return "TEXT[]"
		case timestamp:
			return "TIMESTAMP[]"
		case date:
			return "DATE[]"
		case boolean:
			return "BOOLEAN[]"
		case json:
			return "JSON[]"
		case numeric:
			return "NUMERIC[]"
		}
	case point:
		return "POINT"
	case line:
		return "LINE"
	case lseg:
		return "LSEG"
	case box:
		return "BOX"
	case path:
		return "PATH"
	case polygon:
		return "POLYGON"
	case circle:
		return "CIRCLE"
	case cidr:
		return "CIDR"
	case inet:
		return "INET"
	case macAddr:
		return "MACADDR"
	case macAddr8:
		return "MACADDR8"
	case bit:
		return "BIT(" + fmt.Sprintf("%d", length) + ")"
	case bitVar:
		return "BIT VARYING(" + fmt.Sprintf("%d", length) + ")"
	case xml:
		return "XML"
	case hStore:
		return "HSTORE"
	case int4range:
		return "int4range"
	case int8range:
		return "int8range"
	case numRange:
		return "numrange"
	case tsRange:
		return "tsrange"
	case tstzRange:
		return "tstzrange"
	case dateRange:
		return "daterange"
	case boolean:
		return "BOOLEAN"
	}

	return ""
}
