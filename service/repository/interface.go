package repository

import "github.com/Anushagowda123/employee-management/service/models"

type EmployeeRepository interface {

	// Employee APIs
	CreateEmployee(emp *models.Employee) error
	GetEmployees() ([]models.Employee, error)
	GetEmployeeByID(id int) (*models.Employee, error)
	UpdateEmployee(id int, emp *models.Employee) error
	DeleteEmployee(id int) error

	// Onboarding APIs
	CreateOnboarding(ob *models.Onboarding) error
	ApproveOnboarding(id int) error
	RejectOnboarding(id int) error
	GetOnboardings() ([]models.Onboarding, error)
}
