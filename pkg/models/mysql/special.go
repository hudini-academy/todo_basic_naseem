package mysql

import (
	"database/sql"
)

type SpecialModel struct {
	DB *sql.DB
}

func (m *SpecialModel) Insert(name string) (int, error) {
	stmt := `INSERT INTO special(name,created, expires)
             VALUES(?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	expiry := 365 // Placeholder value of 1 year for expiry

	_, err := m.DB.Exec(stmt, name, expiry)
	if err != nil {
		return 0, err
	}
	return 0, nil
}
func (m *SpecialModel) Delete(name string) (int, error) {
	stmt := `DELETE FROM special WHERE name=?`
	_, err := m.DB.Exec(stmt, name)
	if err != nil {
		return 0, err
	}
	return 0, nil
}
