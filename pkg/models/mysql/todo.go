package mysql

import (
	"database/sql"
	"naseem/pkg/models"
)

// Define a SnippetModel type which wraps a sql.DB connection pool.
type TodoModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *TodoModel) Insert(name string, tags string) (int, error) {
	stmt := `INSERT INTO todos(name,tags,created, expires)
             VALUES(?,?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	expiry := 365 // Placeholder value of 1 year for expiry

	_, err := m.DB.Exec(stmt, name, tags, expiry)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

// This will fetch the all items in database
func (m *TodoModel) Latest() ([]*models.Todo, error) {

	stmt := `SELECT id, name, tags,created, expires FROM todos`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize an empty slice to hold the models.Todo objects.
	Todo := []*models.Todo{}

	// database connection.
	for rows.Next() {
		// Create a pointer to a new zeroed Todo struct.
		s := &models.Todo{}

		err = rows.Scan(&s.ID, &s.Name, &s.Tags, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		Todo = append(Todo, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return Todo, nil
}

// this will delete an item from database using id
func (m *TodoModel) Delete(id int) (int, error) {
	stmt := `DELETE FROM todos WHERE id=?`
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

// this will uodate an item from database
func (m *TodoModel) Update(id int, name string) (int, error) {
	stmt := `UPDATE todos SET Name =? WHERE id =?`
	_, err := m.DB.Exec(stmt, name, id)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (m *TodoModel) Special() ([]*models.Todo, error) {

	stmt := `SELECT id, name, created, expires FROM todos where name like 'special%'`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize an empty slice to hold the models.Todo objects.
	Todo := []*models.Todo{}

	// database connection.
	for rows.Next() {
		// Create a pointer to a new zeroed Todo struct.
		s := &models.Todo{}

		err = rows.Scan(&s.ID, &s.Name, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		Todo = append(Todo, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return Todo, nil
}
