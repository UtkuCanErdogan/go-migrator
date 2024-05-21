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
		return "smallint"
	case integer:
		return "integer"
	case bigInt:
		return "bigint"
	case serial:
		return "serial"
	case bigSerial:
		return "bigserial"
	case numeric:
		return "numeric"
	case real:
		return "real"
	case double:
		return "double"
	case varChar:
		return "varchar(" + fmt.Sprintf("%d", length) + ")"
	case char:
		return "char(" + fmt.Sprintf("%d", length) + ")"
	case text:
		return "text"
	case date:
		return "date"
	case dateTime:
		return "datetime"
	case timestamp:
		return "timestamp"
	case timestampz:
		return "timestampz"
	case interval:
		return "interval"
	case money:
		return "money"
	case uuid:
		return "uuid"
	case json:
		return "json"
	case jsonb:
		return "jsonb"
	case array:
		if arrayType == nil {
			log.Fatal("Invalid operation")
		}

		switch *arrayType {
		case varChar:
			return fmt.Sprintf("varchar(%d)[]", length)
		case integer:
			return "integer[]"
		case text:
			return "text[]"
		case timestamp:
			return "timestamp[]"
		case date:
			return "date[]"
		case boolean:
			return "boolean[]"
		case json:
			return "json[]"
		case numeric:
			return "numeric[]"
		}
	case point:
		return "point"
	case line:
		return "line"
	case lseg:
		return "lseg"
	case box:
		return "box"
	case path:
		return "path"
	case polygon:
		return "polygon"
	case circle:
		return "circle"
	case cidr:
		return "cidr"
	case inet:
		return "inet"
	case macAddr:
		return "macaddr"
	case macAddr8:
		return "macaddr8"
	case bit:
		return "bit(" + fmt.Sprintf("%d", length) + ")"
	case bitVar:
		return "bit varying(" + fmt.Sprintf("%d", length) + ")"
	case xml:
		return "xml"
	case hStore:
		return "hstore"
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
		return "boolean"
	}

	return ""
}
