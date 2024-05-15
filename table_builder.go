package go_migrator

type CreateTableBuilder struct {
	tableName string
	columns   []*ColumnBuilder
}

type AlterTableBuilder struct {
	tableName string
	columns   []*AlterColumnBuilder
}

func (m *Migrator) CreateTable(tableName string) *CreateTableBuilder {
	builder := &CreateTableBuilder{tableName: tableName, columns: nil}
	return builder
}

func (m *Migrator) AlterTable(tableName string) *AlterTableBuilder {
	builder := &AlterTableBuilder{tableName: tableName}
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
	column := &ColumnBuilder{name: columnName, columnType: ColumnTypeSerial}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) UUID(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: ColumnTypeUUID}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) String(columnName string, length int32) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: ColumnTypeString, length: length}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Integer(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: ColumnTypeInteger}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Boolean(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{name: columnName, constraints: nil, columnType: ColumnTypeBoolean}
	t.columns = append(t.columns, column)
	return column
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
	columnType := ColumnTypeSerial
	column := &AlterColumnBuilder{name: columnName, operation: OperationCreateColumn, columnType: &columnType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) UUID(columnName string) *AlterColumnBuilder {
	columnType := ColumnTypeUUID
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &columnType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) String(columnName string, length int32) *AlterColumnBuilder {
	columnType := ColumnTypeString
	column := &AlterColumnBuilder{name: columnName, constraints: nil, columnType: &columnType, operation: OperationCreateColumn, length: length}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Integer(columnName string) *AlterColumnBuilder {
	columnType := ColumnTypeInteger
	column := &AlterColumnBuilder{name: columnName, constraints: nil, operation: OperationCreateColumn, columnType: &columnType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Boolean(columnName string) *AlterColumnBuilder {
	columnType := ColumnTypeBoolean
	column := &AlterColumnBuilder{name: columnName, constraints: nil, columnType: &columnType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Foreign(columnName string) *ForeignBuilder {
	foreign := &ForeignBuilder{referenceColumn: nil, referenceTable: nil}
	column := &AlterColumnBuilder{name: columnName, constraints: []Constraint{{cType: ForeignKeyConstraint}}, operation: OperationCreateColumn, foreignBuilder: foreign}
	a.columns = append(a.columns, column)

	return foreign
}

func (a *AlterTableBuilder) Build() *TableBuilder {
	tableBuilder := &TableBuilder{createBuilder: nil, alterBuilder: a}
	return tableBuilder
}
