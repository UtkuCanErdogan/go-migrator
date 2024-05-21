package go_migrator

type CreateTableBuilder struct {
	tableName string
	columns   []*ColumnBuilder
}

func (m *Migrator) CreateTable(tableName string) *CreateTableBuilder {
	builder := &CreateTableBuilder{tableName: tableName, columns: nil}
	return builder
}

func (m *Migrator) DropTable(tableName string) error {
	db := m.getDbInstance()

	_, err := db.Exec("DROP TABLE " + tableName)
	if err != nil {
		return err
	}

	return nil
}

func (t *CreateTableBuilder) Increments(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, columnType: serial}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) UUID(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: uuid}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) String(columnName string, length int32) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: varChar, length: length}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) SmallInt(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: smallInt}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Integer(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: integer}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) BigInt(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: bigInt}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Serial(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: serial}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) BigSerial(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: bigSerial}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Numeric(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: numeric}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Real(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: real}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Double(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: double}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Char(columnName string, length int32) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: char, length: length}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Text(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: text}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Date(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: date}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) DateTime(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: dateTime}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Timestamp(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: timestamp}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Timestampz(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: timestampz}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Interval(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: interval}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Money(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: money}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Json(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: json}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Jsonb(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: jsonb}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Point(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: point}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Line(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: line}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Lseg(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: lseg}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Box(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: box}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Path(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: path}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Polygon(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: polygon}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Circle(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: circle}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Cidr(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: cidr}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Inet(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: inet}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) MacAddr(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: macAddr}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) MacAddr8(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: macAddr8}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Bit(columnName string, length int32) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: bit, length: length}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) BitVar(columnName string, length int32) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: bitVar, length: length}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Xml(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: xml}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Hstore(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: hStore}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Int4Range(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: int4range}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Int8Range(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: int8range}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) NumRange(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: numRange}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) TsRange(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: tsRange}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) TstzRange(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: tstzRange}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) DateRange(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: dateRange}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Boolean(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: boolean}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Array(columnName string) *ArrayBuilder {
	builder := &ArrayBuilder{}
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: boolean, arrayBuilder: builder}
	t.columns = append(t.columns, column)
	return builder
}

func (t *CreateTableBuilder) Foreign(columnName string) *ForeignBuilder {
	foreign := &ForeignBuilder{referenceColumn: nil, referenceTable: nil}
	column := &ColumnBuilder{name: columnName, constraints: []Constraint{{cType: ForeignKeyConstraint}}, foreignBuilder: foreign}
	t.columns = append(t.columns, column)

	return foreign
}

func (t *CreateTableBuilder) Build() *TableBuilder {
	tableBuilder := &TableBuilder{createBuilder: t, alterBuilder: nil}
	return tableBuilder
}
