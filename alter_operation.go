package go_migrator

type Operation int32

const (
	OperationCreateColumn     Operation = 1
	OperationRenameColumn     Operation = 2
	OperationCreateConstraint Operation = 3
	OperationChangeType       Operation = 4
	OperationDropConstraint   Operation = 5
	OperationDropColumn       Operation = 6
	OperationSetDefault       Operation = 7
)
