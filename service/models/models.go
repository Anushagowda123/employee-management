package models

type Employee struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Position string `db:"position" json:"position"`
	Salary   int    `db:"salary" json:"salary"`
}
