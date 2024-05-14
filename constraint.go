package go_migrator

type Constraint int32

const (
	AutoIncrementConstraint Constraint = 1
	NotNullConstraint       Constraint = 2
	UniqueConstraint        Constraint = 3
	PrimaryKeyConstraint    Constraint = 4
	ForeignKeyConstraint    Constraint = 5
)
