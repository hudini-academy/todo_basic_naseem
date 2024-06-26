package mysql

import (
	"database/sql"
)

type UserModel struct {
	DB *sql.DB
}

// We'll use the Insert method to add a new record to the users table.
func (m *UserModel) Insert(name, email, password string) error {

	stmt := `INSERT INTO users(name,email,hashed_password,created)
             VALUES(?, ?,?,UTC_TIMESTAMP())`

	_, err := m.DB.Exec(stmt, name, email, password)
	if err != nil {
		return err
	}
	return nil
}

// We'll use the Authenticate method to verify whether a user exists with
// the provided email address and password. This will return the relevant
// user ID if they do.
func (m *UserModel) Authenticate(name, password string) (bool, error) {
	stmt := `select id from  users where name=? and hashed_password=?`
	rows, err := m.DB.Query(stmt, name, password)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return rows.Next(), nil
}
