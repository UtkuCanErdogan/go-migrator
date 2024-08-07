package go_migrator

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func (m *Migrator) Migrate() error {
	db := m.getDbInstance()
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS _migration_history (\n    id SERIAL PRIMARY KEY,\n    name varchar(255) NOT NULL,\n    migration_time timestamptz NOT NULL\n)")
	if err != nil {
		return err
	}

	pastHistoryList, err := m.historyList()
	if err != nil {
		return err
	}

	var historyList []MigrationHistory
	var query []string

	for _, migrate := range m.migrate {
		var migrationRan bool
		for _, history := range pastHistoryList {
			if history.Name == migrate.key {
				migrationRan = true
				break
			}
		}

		if migrationRan {
			continue
		}

		tableBuilder := migrate.fn(m)

		if tableBuilder != nil && tableBuilder.createBuilder != nil {
			builder := *tableBuilder.createBuilder

			createQuery := "CREATE TABLE " + *m.Config.Schema + "." + builder.tableName + "(\n"
			for index, column := range builder.columns {
				var arrayType *columnType
				var referenceQuery string
				if column.arrayBuilder != nil {
					aType := column.arrayBuilder.columnType
					arrayType = &aType

					if column.arrayBuilder.referenceColumn != nil && column.arrayBuilder.referenceTable != nil {
						referenceQuery = fmt.Sprintf("REFERENCES %s (%s)", *column.arrayBuilder.referenceTable, *column.arrayBuilder.referenceColumn)
					}
				}

				createQuery = createQuery + column.name + " " + column.columnType.toString(column.length, arrayType) + referenceQuery
				if len(column.constraints) > 0 {
					for _, cons := range column.constraints {
						createQuery = createQuery + " " + cons.cType.toString()
					}
				}

				if column.defaultValue != nil {
					createQuery = createQuery + " DEFAULT " + fmt.Sprintf("%v", column.defaultValue)
				}

				if column.foreignBuilder != nil {
					if column.foreignBuilder.referenceColumn == nil || column.foreignBuilder.referenceTable == nil {
						return errors.New("reference column and table required for foreign operation")
					}

					createQuery = createQuery + "CONSTRAINT fk_" + column.name + " FOREIGN KEY(" + *column.foreignBuilder.referenceColumn + ") " + "REFERENCES " + *m.Config.Schema + "." + *column.foreignBuilder.referenceTable
				}

				if !(index == len(builder.columns)-1) {
					createQuery = createQuery + " ," + "\n"
				}
			}

			createQuery = createQuery + ")"
			query = append(query, createQuery)
		}

		if tableBuilder != nil && tableBuilder.alterBuilder != nil {
			builder := *tableBuilder.alterBuilder

			for _, column := range builder.columns {
				alterQuery := "ALTER TABLE " + *m.Config.Schema + "." + builder.tableName

				if column.rawSql != nil {
					alterQuery = *column.rawSql
				} else {
					switch column.operation {
					case OperationCreateColumn:
						var arrayType *columnType
						var referenceQuery string
						if column.arrayBuilder != nil {
							aType := column.arrayBuilder.columnType
							arrayType = &aType

							if column.arrayBuilder.referenceColumn != nil && column.arrayBuilder.referenceTable != nil {
								referenceQuery = fmt.Sprintf("REFERENCES %s (%s)", *column.arrayBuilder.referenceTable, *column.arrayBuilder.referenceColumn)
							}
						}

						alterQuery = alterQuery + " ADD COLUMN " + column.name + " " + column.columnType.toString(column.length, arrayType) + referenceQuery
						if column.defaultValue != nil {
							alterQuery = alterQuery + " DEFAULT " + fmt.Sprintf("%v", column.defaultValue)
						}

						if column.foreignBuilder != nil {
							if column.foreignBuilder.referenceColumn == nil || column.foreignBuilder.referenceTable == nil {
								return errors.New("reference column and table required for foreign operation")
							}

							alterQuery = "ALTER TABLE " + builder.tableName + " ADD CONSTRAINT fk_" + column.name + " FOREIGN KEY(" + *column.foreignBuilder.referenceColumn + ") " + "REFERENCES " + *m.Config.Schema + "." + *column.foreignBuilder.referenceTable
						}

						if len(column.constraints) > 0 {
							for _, cons := range column.constraints {
								alterQuery = alterQuery + " " + cons.cType.toString()
							}
						}
					case OperationRenameColumn:
						if column.newName == nil {
							return errors.New("column new name required")
						}

						alterQuery = alterQuery + " RENAME COLUMN " + column.name + " TO " + *column.newName
					case OperationCreateConstraint:
						cons := column.constraints[0]
						if column.foreignBuilder != nil {
							if column.foreignBuilder.referenceColumn == nil || column.foreignBuilder.referenceTable == nil {
								return errors.New("reference column and table required for foreign operation")
							}

							alterQuery = alterQuery + " ADD CONSTRAINT fk_" + column.name + " FOREIGN KEY(" + *column.foreignBuilder.referenceColumn + ") " + "REFERENCES " + *m.Config.Schema + *column.foreignBuilder.referenceTable
						} else {
							alterQuery = alterQuery + " ADD CONSTRAINT " + strings.ToLower(column.name) + "_" + cons.cType.ToLower() + " " + cons.cType.toString() + "("
							alterQuery = alterQuery + column.name + ")"
						}
					case OperationChangeType:
						if column.columnType == nil {
							return errors.New("column type required")
						}

						alterQuery = alterQuery + " ALTER COLUMN " + column.name + " TYPE " + column.columnType.toString(column.length, nil)
					case OperationDropConstraint:
						cons := column.constraints[0]
						alterQuery = alterQuery + " DROP CONSTRAINT " + cons.name
					case OperationDropColumn:
						alterQuery = alterQuery + " DROP COLUMN " + column.name
					case OperationSetDefault:
						if column.defaultValue == nil {
							return errors.New("default value cannot be empty")
						}

						alterQuery = alterQuery + " ALTER COLUMN " + column.name + " SET DEFAULT " + fmt.Sprintf("%v", column.defaultValue)
					default:
						return errors.New("un supported operation")
					}
				}

				query = append(query, alterQuery)
			}
		}

		var hasError bool
		for _, q := range query {
			tx, err := db.Begin()
			if err != nil {
				return err
			}

			_, err = tx.Exec(q)
			if err != nil {
				hasError = true
				err = tx.Rollback()
				if err != nil {
					return err
				}

				break
			}

			err = tx.Commit()
			if err != nil {
				err = tx.Rollback()
				if err != nil {
					return err
				}

				return err
			}
		}

		if !hasError {
			history := MigrationHistory{
				Name:          migrate.key,
				MigrationTime: time.Now(),
			}

			historyList = append(historyList, history)
		}

		query = nil
	}

	if len(historyList) > 0 {
		for {
			if len(historyList) > 500 {
				willSaveList := historyList[:500]
				err = m.saveHistory(willSaveList)
				if err != nil {
					return err
				}

				historyList = historyList[500:]
				if len(historyList) == 0 {
					break
				}
			} else {
				err = m.saveHistory(historyList)
				if err != nil {
					return err
				}

				break
			}
		}

	}

	return nil
}
