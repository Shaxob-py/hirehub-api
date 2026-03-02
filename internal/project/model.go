package project

import "time"

type ModelProject struct {
	ID          string    `db:"id"`
	UserID      string    `db:"name"`
	Description string    `db:"description"`
	Price       float64   `db:"price"`
	Skill       string    `db:"skill"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type ModelCreateProject struct {
	ID          string    `db:"id"`
	Description string    `db:"description"`
	Price       float64   `db:"price"`
	Skill       string    `db:"skill"`
	UserID      string    `db:"name"`
	CreatedAt   time.Time `db:"created_at"`
}

type ModelUpdateProject struct {
	ID          string  `db:"id"`
	Description string  `db:"description"`
	Price       float64 `db:"price"`
	Skill       string  `db:"skill"`
}
