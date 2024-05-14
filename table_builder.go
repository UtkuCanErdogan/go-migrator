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
	return &CreateTableBuilder{tableName: tableName, columns: nil}
}

func (m *Migrator) AlterTable(tableName string) *AlterTableBuilder {
	builder := AlterTableBuilder{tableName: tableName}
	return &builder
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
	column := &ColumnBuilder{Name: columnName, Constraints: []Constraint{AutoIncrementConstraint}, ColumnType: ColumnTypeInteger}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) UUID(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{Name: columnName, Constraints: nil, ColumnType: ColumnTypeUUID}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) String(columnName string, length int32) *ColumnBuilder {
	column := &ColumnBuilder{Name: columnName, Constraints: nil, ColumnType: ColumnTypeString, Length: length}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Integer(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{Name: columnName, Constraints: nil, ColumnType: ColumnTypeInteger}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Boolean(columnName string) *ColumnBuilder {
	column := &ColumnBuilder{Name: columnName, Constraints: nil, ColumnType: ColumnTypeBoolean}
	t.columns = append(t.columns, column)
	return column
}

func (t *CreateTableBuilder) Foreign(columnName string) *ForeignBuilder {
	foreign := &ForeignBuilder{ReferenceColumn: nil, ReferenceTable: nil}
	column := &ColumnBuilder{Name: columnName, Constraints: []Constraint{ForeignKeyConstraint}, ColumnType: ColumnTypeBoolean, ForeignBuilder: foreign}
	t.columns = append(t.columns, column)

	return foreign
}

func (a *AlterTableBuilder) DropColumn(columnName string) *AlterTableBuilder {
	column := &AlterColumnBuilder{Name: columnName, Operation: OperationDropColumn, Constraint: nil, ForeignBuilder: nil, DefaultValue: nil}
	a.columns = append(a.columns, column)
	return a
}

func (a *AlterTableBuilder) RenameColumn(columnName, newColumnName string) *AlterTableBuilder {
	column := &AlterColumnBuilder{Name: columnName, NewName: &newColumnName, Operation: OperationRenameColumn, Constraint: nil, ForeignBuilder: nil, DefaultValue: nil}
	a.columns = append(a.columns, column)
	return a
}

func (a *AlterTableBuilder) CreateConstraint(columnName string, constraint Constraint) *AlterTableBuilder {
	column := &AlterColumnBuilder{Name: columnName, Operation: OperationCreateConstraint, Constraint: []Constraint{constraint}, ForeignBuilder: nil, DefaultValue: nil}
	a.columns = append(a.columns, column)
	return a
}

func (a *AlterTableBuilder) DropConstraint(columnName string, constraint Constraint) *AlterTableBuilder {
	column := &AlterColumnBuilder{Name: columnName, Operation: OperationDropConstraint, Constraint: []Constraint{constraint}, ForeignBuilder: nil, DefaultValue: nil}
	a.columns = append(a.columns, column)
	return a
}

func (a *AlterTableBuilder) Increments(columnName string) *AlterColumnBuilder {
	columnType := ColumnTypeInteger
	column := &AlterColumnBuilder{Name: columnName, Constraint: []Constraint{AutoIncrementConstraint}, Operation: OperationCreateColumn, ColumnType: &columnType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) UUID(columnName string) *AlterColumnBuilder {
	columnType := ColumnTypeUUID
	column := &AlterColumnBuilder{Name: columnName, Constraint: nil, Operation: OperationCreateColumn, ColumnType: &columnType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) String(columnName string, length int32) *AlterColumnBuilder {
	columnType := ColumnTypeString
	column := &AlterColumnBuilder{Name: columnName, Constraint: nil, ColumnType: &columnType, Operation: OperationCreateColumn, Length: length}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Integer(columnName string) *AlterColumnBuilder {
	columnType := ColumnTypeInteger
	column := &AlterColumnBuilder{Name: columnName, Constraint: nil, Operation: OperationCreateColumn, ColumnType: &columnType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Boolean(columnName string) *AlterColumnBuilder {
	columnType := ColumnTypeBoolean
	column := &AlterColumnBuilder{Name: columnName, Constraint: nil, ColumnType: &columnType}
	a.columns = append(a.columns, column)
	return column
}

func (a *AlterTableBuilder) Foreign(columnName string) *ForeignBuilder {
	foreign := &ForeignBuilder{ReferenceColumn: nil, ReferenceTable: nil}
	column := &AlterColumnBuilder{Name: columnName, Constraint: []Constraint{ForeignKeyConstraint}, Operation: OperationCreateColumn, ForeignBuilder: foreign}
	a.columns = append(a.columns, column)

	return foreign
}
