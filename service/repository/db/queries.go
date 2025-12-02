package db

const CreateEmployeeQuery = `
INSERT INTO employees (name, position, salary) VALUES (?, ?, ?);
`

const GetAllEmployeesQuery = `
SELECT id, name, position, salary FROM employees;
`

const GetEmployeeByIDQuery = `
SELECT id, name, position, salary FROM employees WHERE id = ?;
`

const DeleteEmployeeQuery = `
DELETE FROM employees WHERE id = ?;
`
