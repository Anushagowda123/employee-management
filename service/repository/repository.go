package repository

import "employee-management/service/models"

type EmployeeRepository interface {
	CreateEmployee(emp models.Employee) error
	GetAllEmployees() ([]models.Employee, error)
	GetEmployeeByID(id int) (*models.Employee, error)
	DeleteEmployee(id int) error
}
