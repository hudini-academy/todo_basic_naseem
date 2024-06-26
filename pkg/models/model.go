package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Todo struct {
	ID      int
	Name    string
	Tags    string
	Created time.Time
	Expires time.Time
}
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}
type Special struct {
	ID      int
	Name    string
	Created time.Time
	Expires time.Time
}
