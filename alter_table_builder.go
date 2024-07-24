package go_migrator

type AlterTableBuilder struct {
	tableName string
	columns   []*AlterColumnBuilder
}

func (m *Migrator) AlterTable(tableName string) *AlterTableBuilder {
	builder := &AlterTableBuilder{tableName: tableName}
	return builder
}

func (a *AlterTableBuilder) DropColumn(columnName string) *AlterTableBuilder {
	column := &AlterColumnBuilder{name: columnName, operation: OperationDropColumn, constraints: nil, foreignBuilder: nil, defaultValue: nil}
	a.columns = append(a.columns, column)
	return a
}

func (a *AlterTableBuilder) RenameColumn(columnName, newColumnName string) *AlterTableBuilder {
	column := &AlterColumnBuilder{name: columnName, newName: &newColumnName, operation: OperationRenameColumn, constraints: nil, foreignBuilder: nil, defaultValue: nil}
	a.columns = append(a.columns, column)
	return a
}

func (a *AlterTableBuilder) CreateConstraint(columnName string, constraint ConstraintType) *AlterTableBuilder {
	column := &AlterColumnBuilder{name: columnName, operation: OperationCreateConstraint, constraints: []Constraint{{cType: constraint}}, foreignBuilder: nil, defaultValue: nil}
	a.columns = append(a.columns, column)
	return a
}

func (a *AlterTableBuilder) DropConstraint(columnName, constraintName string, constraintType ConstraintType) *AlterTableBuilder {
	column := &AlterColumnBuilder{name: columnName, operation: OperationDropConstraint, constraints: []Constraint{{cType: constraintType, name: constraintName}}, foreignBuilder: nil, defaultValue: nil}
	a.columns = append(a.columns, column)
	return a
}

func (a *AlterTableBuilder) Increments(columnName string) *AlterColumnBuilder {
	cType := serial
	column := &AlterColumnBuilder{name: columnName, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) UUID(columnName string) *AlterColumnBuilder {
	cType := uuid
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) String(columnName string, length int32) *AlterColumnBuilder {
	cType := varChar
	column := &AlterColumnBuilder{name: columnName, constraints: nil, columnType: &cType, operation: OperationCreateColumn, length: length}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) SmallInt(columnName string) *AlterColumnBuilder {
	cType := smallInt
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Integer(columnName string) *AlterColumnBuilder {
	cType := integer
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) BigInt(columnName string) *AlterColumnBuilder {
	cType := bigInt
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Serial(columnName string) *AlterColumnBuilder {
	cType := serial
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) BigSerial(columnName string) *AlterColumnBuilder {
	cType := bigSerial
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Numeric(columnName string) *AlterColumnBuilder {
	cType := numeric
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Real(columnName string) *AlterColumnBuilder {
	cType := real
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Double(columnName string) *AlterColumnBuilder {
	cType := double
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Char(columnName string, length int32) *AlterColumnBuilder {
	cType := char
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType, length: length}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Text(columnName string) *AlterColumnBuilder {
	cType := text
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Date(columnName string) *AlterColumnBuilder {
	cType := date
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) DateTime(columnName string) *AlterColumnBuilder {
	cType := dateTime
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Timestamp(columnName string) *AlterColumnBuilder {
	cType := timestamp
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Timestampz(columnName string) *AlterColumnBuilder {
	cType := timestampz
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Interval(columnName string) *AlterColumnBuilder {
	cType := interval
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Money(columnName string) *AlterColumnBuilder {
	cType := money
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Json(columnName string) *AlterColumnBuilder {
	cType := json
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Jsonb(columnName string) *AlterColumnBuilder {
	cType := jsonb
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Point(columnName string) *AlterColumnBuilder {
	cType := point
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Line(columnName string) *AlterColumnBuilder {
	cType := line
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Lseg(columnName string) *AlterColumnBuilder {
	cType := lseg
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Box(columnName string) *AlterColumnBuilder {
	cType := box
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Path(columnName string) *AlterColumnBuilder {
	cType := path
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Polygon(columnName string) *AlterColumnBuilder {
	cType := polygon
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Circle(columnName string) *AlterColumnBuilder {
	cType := circle
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Cidr(columnName string) *AlterColumnBuilder {
	cType := cidr
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Inet(columnName string) *AlterColumnBuilder {
	cType := inet
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) MacAddr(columnName string) *AlterColumnBuilder {
	cType := macAddr
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) MacAddr8(columnName string) *AlterColumnBuilder {
	cType := macAddr8
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Bit(columnName string, length int32) *AlterColumnBuilder {
	cType := bit
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType, length: length}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) BitVar(columnName string, length int32) *AlterColumnBuilder {
	cType := bitVar
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType, length: length}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Xml(columnName string) *AlterColumnBuilder {
	cType := xml
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Hstore(columnName string) *AlterColumnBuilder {
	cType := hStore
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Int4Range(columnName string) *AlterColumnBuilder {
	cType := int4range
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Int8Range(columnName string) *AlterColumnBuilder {
	cType := int8range
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) NumRange(columnName string) *AlterColumnBuilder {
	cType := int8range
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) TsRange(columnName string) *AlterColumnBuilder {
	cType := tsRange
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) TstzRange(columnName string) *AlterColumnBuilder {
	cType := tstzRange
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) DateRange(columnName string) *AlterColumnBuilder {
	cType := dateRange
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Boolean(columnName string) *AlterColumnBuilder {
	cType := boolean
	column := &AlterColumnBuilder{name: columnName, constraints: nil, columnType: &cType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Array(columnName string) *ArrayBuilder {
	cType := array
	builder := &ArrayBuilder{referenceColumn: nil, referenceTable: nil}
	column := &AlterColumnBuilder{name: columnName, columnType: &cType, arrayBuilder: builder}
	a.columns = append(a.columns, column)

	return builder
}

func (a *AlterColumnBuilder) Foreign() *ForeignBuilder {
	foreign := &ForeignBuilder{referenceColumn: nil, referenceTable: nil}
	return foreign
}

func (a *AlterTableBuilder) Build() *TableBuilder {
	tableBuilder := &TableBuilder{createBuilder: nil, alterBuilder: a}
	return tableBuilder
}
