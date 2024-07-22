package go_migrator

type ColumnBuilder struct {
	name           string
	columnType     columnType
	constraints    []Constraint
	length         int32
	defaultValue   any
	foreignBuilder *ForeignBuilder
	arrayBuilder   *ArrayBuilder
}

type AlterColumnBuilder struct {
	name           string
	newName        *string
	columnType     *columnType
	operation      Operation
	constraints    []Constraint
	length         int32
	defaultValue   any
	foreignBuilder *ForeignBuilder
	arrayBuilder   *ArrayBuilder
	rawSql         *string
}

type ForeignBuilder struct {
	referenceColumn *string
	referenceTable  *string
}

type ArrayBuilder struct {
	referenceColumn *string
	referenceTable  *string
	columnType      columnType
	length          int32
	defaultValue    any
	constraints     []Constraint
}

func (f *ForeignBuilder) Reference(referenceColumn string) *ForeignBuilder {
	f.referenceColumn = &referenceColumn
	return f
}

func (f *ForeignBuilder) Table(referenceTable string) *ForeignBuilder {
	f.referenceTable = &referenceTable
	return f
}

func (c *ColumnBuilder) NotNull() *ColumnBuilder {
	c.constraints = append(c.constraints, Constraint{cType: notNullConstraint})
	return c
}

func (c *ColumnBuilder) Unique() *ColumnBuilder {
	c.constraints = append(c.constraints, Constraint{cType: UniqueConstraint})
	return c
}

func (c *ColumnBuilder) Primary() *ColumnBuilder {
	c.constraints = append(c.constraints, Constraint{cType: PrimaryKeyConstraint})
	return c
}

func (c *ColumnBuilder) Default(value any) *ColumnBuilder {
	c.defaultValue = value
	return c
}

func (a *ArrayBuilder) String(length int32) *ArrayBuilder {
	a.columnType = varChar
	a.length = length
	return a
}

func (a *ArrayBuilder) Integer() *ArrayBuilder {
	a.columnType = integer
	return a
}

func (a *ArrayBuilder) Text() *ArrayBuilder {
	a.columnType = text
	return a
}

func (a *ArrayBuilder) Date() *ArrayBuilder {
	a.columnType = date
	return a
}

func (a *ArrayBuilder) Timestamp() *ArrayBuilder {
	a.columnType = timestamp
	return a
}

func (a *ArrayBuilder) Json() *ArrayBuilder {
	a.columnType = json
	return a
}

func (a *ArrayBuilder) Numeric() *ArrayBuilder {
	a.columnType = numeric
	return a
}

func (a *ArrayBuilder) Boolean() *ArrayBuilder {
	a.columnType = boolean
	return a
}

func (a *ArrayBuilder) Reference(referenceColumn string) *ArrayBuilder {
	a.referenceColumn = &referenceColumn
	return a
}

func (a *ArrayBuilder) Table(referenceTable string) *ArrayBuilder {
	a.referenceTable = &referenceTable
	return a
}

func (a *ArrayBuilder) Default(value any) *ArrayBuilder {
	a.defaultValue = value
	return a
}

func (a *ArrayBuilder) NotNull() *ArrayBuilder {
	a.constraints = append(a.constraints, Constraint{cType: notNullConstraint})
	return a
}

func (a *ArrayBuilder) Unique() *ArrayBuilder {
	a.constraints = append(a.constraints, Constraint{cType: UniqueConstraint})
	return a
}

func (a *AlterColumnBuilder) Primary() *AlterColumnBuilder {
	a.constraints = append(a.constraints, Constraint{cType: notNullConstraint})
	return a
}

func (a *AlterColumnBuilder) Default(value any) *AlterColumnBuilder {
	a.defaultValue = value
	return a
}

func (a *AlterColumnBuilder) NotNull() *AlterColumnBuilder {
	a.constraints = append(a.constraints, Constraint{cType: notNullConstraint})
	return a
}

func (a *AlterColumnBuilder) Unique() *AlterColumnBuilder {
	a.constraints = append(a.constraints, Constraint{cType: UniqueConstraint})
	return a
}

func (a *AlterColumnBuilder) RawSql(sql string) {
	a.rawSql = &sql
}
