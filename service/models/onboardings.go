package models

import "time"

type Onboarding struct {
	ID         int       `json:"id"`
	EmployeeID int       `json:"employee_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
