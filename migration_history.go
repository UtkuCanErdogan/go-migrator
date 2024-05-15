package go_migrator

import (
	"fmt"
	"time"
)

type MigrationHistory struct {
	Id            int32     `json:"id"`
	Name          string    `json:"name"`
	MigrationTime time.Time `json:"migration_time"`
}

func (m *Migrator) saveHistory(migrationHistory []MigrationHistory) error {
	db := m.getDbInstance()
	query := "INSERT INTO _migration_history (name,migration_time) VALUES"
	for index, history := range migrationHistory {
		if index > 0 {
			query = query + ","
		}

		query = query + fmt.Sprintf("('%s','%s')", history.Name, history.MigrationTime.Format(time.RFC3339))
	}

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (m *Migrator) historyList() ([]MigrationHistory, error) {
	db := m.getDbInstance()
	migrationName := "_migration_history"
	if m.Config.Migration != nil && m.Config.Migration.TableName != nil {
		migrationName = *m.Config.Migration.TableName
	}

	query := fmt.Sprintf("SELECT * FROM %s", migrationName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var migrationHistory []MigrationHistory

	for rows.Next() {
		var (
			id            int32
			name          string
			migrationTime time.Time
		)
		if err = rows.Scan(&id, &name, &migrationTime); err != nil {
			return nil, err
		}

		history := MigrationHistory{Id: id, Name: name, MigrationTime: migrationTime}
		migrationHistory = append(migrationHistory, history)
	}

	return migrationHistory, nil
}
