package go_migrator

type ColumnBuilder struct {
	name           string
	columnType     ColumnType
	constraints    []Constraint
	length         int32
	defaultValue   any
	foreignBuilder *ForeignBuilder
}

type AlterColumnBuilder struct {
	name           string
	newName        *string
	columnType     *ColumnType
	operation      Operation
	constraints    []Constraint
	length         int32
	defaultValue   any
	foreignBuilder *ForeignBuilder
}

type ForeignBuilder struct {
	referenceColumn *string
	referenceTable  *string
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
	c.constraints = append(c.constraints, Constraint{cType: NotNullConstraint})
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

func (a *AlterColumnBuilder) Primary() *AlterColumnBuilder {
	a.constraints = append(a.constraints, Constraint{cType: NotNullConstraint})
	return a
}

func (a *AlterColumnBuilder) Default(value any) *AlterColumnBuilder {
	a.defaultValue = value
	return a
}

func (a *AlterColumnBuilder) NotNull() *AlterColumnBuilder {
	a.constraints = append(a.constraints, Constraint{cType: NotNullConstraint})
	return a
}

func (a *AlterColumnBuilder) Unique() *AlterColumnBuilder {
	a.constraints = append(a.constraints, Constraint{cType: UniqueConstraint})
	return a
}
