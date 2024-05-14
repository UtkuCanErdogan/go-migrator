package go_migrator

type ColumnBuilder struct {
	Name           string
	ColumnType     ColumnType
	Constraints    []Constraint
	Length         int32
	DefaultValue   any
	ForeignBuilder *ForeignBuilder
}

type AlterColumnBuilder struct {
	Name           string
	NewName        *string
	ColumnType     *ColumnType
	Operation      Operation
	Constraint     []Constraint
	Length         int32
	DefaultValue   any
	ForeignBuilder *ForeignBuilder
}

type ForeignBuilder struct {
	ReferenceColumn *string
	ReferenceTable  *string
}

func (f *ForeignBuilder) Reference(referenceColumn string) *ForeignBuilder {
	f.ReferenceColumn = &referenceColumn
	return f
}

func (f *ForeignBuilder) Table(referenceTable string) *ForeignBuilder {
	f.ReferenceTable = &referenceTable
	return f
}

func (c *ColumnBuilder) NotNull() *ColumnBuilder {
	c.Constraints = append(c.Constraints, NotNullConstraint)
	return c
}

func (c *ColumnBuilder) Unique() *ColumnBuilder {
	c.Constraints = append(c.Constraints, UniqueConstraint)
	return c
}

func (c *ColumnBuilder) Primary() *ColumnBuilder {
	c.Constraints = append(c.Constraints, PrimaryKeyConstraint)
	return c
}

func (c *ColumnBuilder) Default(value any) *ColumnBuilder {
	c.DefaultValue = value
	return c
}

func (a *AlterColumnBuilder) Primary() *AlterColumnBuilder {
	a.Constraint = append(a.Constraint, PrimaryKeyConstraint)
	return a
}

func (a *AlterColumnBuilder) Default(value any) *AlterColumnBuilder {
	a.DefaultValue = value
	return a
}
