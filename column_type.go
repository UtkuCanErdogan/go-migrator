package go_migrator

type ColumnType int32

const (
	ColumnTypeString    ColumnType = 1
	ColumnTypeInteger   ColumnType = 2
	ColumnTypeUUID      ColumnType = 3
	ColumnTypeBoolean   ColumnType = 4
	ColumnTypeFloat     ColumnType = 5
	ColumnTypeDouble    ColumnType = 6
	ColumnTypeDecimal   ColumnType = 7
	ColumnTypeTimestamp ColumnType = 8
)
