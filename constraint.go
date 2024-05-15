package go_migrator

type Constraint struct {
	name  string
	cType ConstraintType
}

type ConstraintType int32

const (
	NotNullConstraint    ConstraintType = 1
	UniqueConstraint     ConstraintType = 2
	PrimaryKeyConstraint ConstraintType = 3
	ForeignKeyConstraint ConstraintType = 4
)

func (c ConstraintType) toString() string {
	switch c {
	case NotNullConstraint:
		return "NOT NULL"
	case UniqueConstraint:
		return "UNIQUE"
	case PrimaryKeyConstraint:
		return "PRIMARY KEY"
	}

	return ""
}

func (c ConstraintType) ToLower() string {
	switch c {
	case UniqueConstraint:
		return "unique"
	case PrimaryKeyConstraint:
		return "primary_key"
	}

	return ""
}
