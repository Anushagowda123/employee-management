package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"employee-management/service/models"
	"employee-management/service/repository"
)

var DB *sqlx.DB

// EmployeeDB implements repository.EmployeeRepository
type EmployeeDB struct {
	DB *sqlx.DB
}

func Connect() {
	var err error
	// default XAMPP credentials: user root, empty password
	dsn := "root:@tcp(127.0.0.1:3306)/employee_db?parseTime=true"

	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB open error:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}

	fmt.Println("Connected to MySQL database!")
}

// NewEmployeeRepo returns an implementation of repository.EmployeeRepository
func NewEmployeeRepo() repository.EmployeeRepository {
	return &EmployeeDB{DB: DB}
}

func (e *EmployeeDB) CreateEmployee(emp models.Employee) error {
    _, err := e.DB.Exec(CreateEmployeeQuery, emp.Name, emp.Position, emp.Salary)
    if err != nil {
        fmt.Println("SQL ERROR:", err)  
        return err
    }

    fmt.Println("INSERT SUCCESS")   
    return nil
}

func (e *EmployeeDB) GetAllEmployees() ([]models.Employee, error) {
	var emps []models.Employee
	err := e.DB.Select(&emps, GetAllEmployeesQuery)
	return emps, err
}

func (e *EmployeeDB) GetEmployeeByID(id int) (*models.Employee, error) {
	var emp models.Employee
	err := e.DB.Get(&emp, GetEmployeeByIDQuery, id)
	if err == sql.ErrNoRows {
		return nil, err
	}
	return &emp, err
}

func (e *EmployeeDB) DeleteEmployee(id int) error {
	_, err := e.DB.Exec(DeleteEmployeeQuery, id)
	return err
}
