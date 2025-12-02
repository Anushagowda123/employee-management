**EMPLOYEE MANAGEMENT SYSTEM (Go + Gin + MySQL)

->A simple Employee Management REST API built using Golang, Gin Web Framework and MySQL. 
->This project implements a clean folder structure using Separation of Concerns the Repository Pattern, and proper layered architecture.

**FEATURES IMPLEMENTED

- Create a new employee (POST)
- Get a list of all employees (GET)
- Get a single employee by ID (GET)
- Delete an employee (DELETE)
- Modular folder structure using handlers, models, repository, and DB layers
- Fully tested using Postman

 ** PROJECT FOLDER STRUCTURE

employee-management/
├── go.mod
├── go.sum
├── main.go
└── service/
├── handler/
│ └── handler.go       (Handles API logic)
├── libhttp/
│ └── routes.go        (API route definitions)
├── models/
│ └── models.go        (Employee struct model)
└── repository/
├── repository.go      (Repository interface)
├── queries.go         (SQL queries for CRUD)
└── db/
└── db.go              (MySQL connection & DB operations)

** TECHNOLOGIES USED

->Go 1.20+
->Gin Web Framework
->MySQL (XAMPP)
->Postman for API testing

 **DATABASE SETUP

 1. Start MySQL
->Open XAMPP
->Start MySQL

2. Create Database in phpMyAdmin
Go to → http://localhost/phpmyadmin

Run this:

CREATE DATABASE employee_db;
USE employee_db;

CREATE TABLE employees (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  position VARCHAR(255),
  salary INT
);

**RUNNING THE APPLICATION

->Open VS Code terminal inside the project folder:
cd employee-management
->RUN :go run main.go
->IF SUCCESS:Connected to MySQL database!
[GIN-debug] Listening and serving HTTP on :8081
Your API is live at:http://localhost:8081

**API ENDPOINTS (Use Postman)

➤ Create Employee
->POST
http://localhost:8081/employees/
Body (raw → JSON):{
  "name": "Anu",
  "position": "Developer",
  "salary": 70000
}
->Get All Employees
->GET :http://localhost:8081/employees/

->Get Employee By ID
GET:http://localhost:8081/employees/1

->Delete Employee
DELETE:http://localhost:8081/employees/1

**TROUBLESHOOTING

-> MySQL connection refused

Start MySQL from XAMPP
Check DB username/password in db.go

-> Duplicate ID / insert errors

Ensure column id is:
PRIMARY KEY
AUTO_INCREMENT

-> Port already in use
Change port in main.go:router.Run(":8081")


**SUBMISSION

My public GitHub repository link:

https://github.com/<Anushagowda123>/employee-management


This README serves as the documentation for my assignment.


Anusha C
 As a beginner in Go — I completed the project successfully with proper architecture and API testing.

