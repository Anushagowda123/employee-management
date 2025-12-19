package db

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Anushagowda123/employee-management/service/models"
)

type EmployeeRepo struct {
	DB *sql.DB
}

/*
CONNECT TO DATABASE
*/
func Connect() *sql.DB {
	db, err := sql.Open(
		"mysql",
		"root:@tcp(127.0.0.1:3306)/employee_management?parseTime=true",
	)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}

/*
CREATE EMPLOYEE
*/
func (r *EmployeeRepo) CreateEmployee(emp *models.Employee) error {
	query := `
		INSERT INTO employees (name, email, department, salary)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.DB.Exec(
		query,
		emp.Name,
		emp.Email,
		emp.Department,
		emp.Salary,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	emp.ID = int(id)
	return nil
}

/*
GET ALL EMPLOYEES
*/
func (r *EmployeeRepo) GetEmployees() ([]models.Employee, error) {
	query := `
		SELECT id, name, email, department, salary, created_at
		FROM employees
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := []models.Employee{}

	for rows.Next() {
		var emp models.Employee
		err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Email,
			&emp.Department,
			&emp.Salary,
			&emp.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		employees = append(employees, emp)
	}

	return employees, nil
}

/*
GET EMPLOYEE BY ID
*/
func (r *EmployeeRepo) GetEmployeeByID(id int) (*models.Employee, error) {
	query := `
		SELECT id, name, email, department, salary, created_at
		FROM employees
		WHERE id = ?
	`

	var emp models.Employee

	err := r.DB.QueryRow(query, id).Scan(
		&emp.ID,
		&emp.Name,
		&emp.Email,
		&emp.Department,
		&emp.Salary,
		&emp.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("employee not found")
	}
	if err != nil {
		return nil, err
	}

	return &emp, nil
}

/*
UPDATE EMPLOYEE
*/
func (r *EmployeeRepo) UpdateEmployee(id int, emp *models.Employee) error {
	query := `
		UPDATE employees
		SET name = ?, email = ?, department = ?, salary = ?
		WHERE id = ?
	`

	result, err := r.DB.Exec(
		query,
		emp.Name,
		emp.Email,
		emp.Department,
		emp.Salary,
		id,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("employee not found")
	}

	return nil
}

/*
DELETE EMPLOYEE
*/
func (r *EmployeeRepo) DeleteEmployee(id int) error {
	query := `
	DELETE FROM employees
	WHERE id = ?
	`

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("employee not found")
	}

	return nil
}

// ONBOARDING 

func (r *EmployeeRepo) CreateOnboarding(ob *models.Onboarding) error {
	query := `
		INSERT INTO onboardings (employee_id, status)
		VALUES (?, ?)
	`
	_, err := r.DB.Exec(query, ob.EmployeeID, "PENDING")
	return err
}

func (r *EmployeeRepo) ApproveOnboarding(id int) error {
	query := `
		UPDATE onboardings
		SET status = 'APPROVED'
		WHERE id = ?
	`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *EmployeeRepo) RejectOnboarding(id int) error {
	query := `
		UPDATE onboardings
		SET status = 'REJECTED'
		WHERE id = ?
	`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *EmployeeRepo) GetOnboardings() ([]models.Onboarding, error) {
	rows, err := r.DB.Query(`
		SELECT id, employee_id, status, created_at
		FROM onboardings
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Onboarding

	for rows.Next() {
		var ob models.Onboarding
		if err := rows.Scan(
			&ob.ID,
			&ob.EmployeeID,
			&ob.Status,
			&ob.CreatedAt,
		); err != nil {
			return nil, err
		}
		list = append(list, ob)
	}

	return list, nil
}
                          




