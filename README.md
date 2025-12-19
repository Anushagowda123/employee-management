Employee Management System – Golang Backend Project

 ->PROJECT OVERVIEW

>>The Employee Management System is a backend RESTful API developed using Go (Golang).
It manages Employees and their Onboarding lifecycle, supporting full CRUD operations and approval workflows.

>>This project is designed to demonstrate real-world backend development skills including clean architecture, interface-driven design, database integration, concurrency readiness, and production-grade API development.

>>All APIs return JSON responses and are tested using Postman.

->TECHNOLOGY STACK

Programming Language: Go (Golang)
Web Framework: Gin
Database: MySQL
Database Access: database/sql
API Style: REST (JSON)
Testing Tool: Postman
Architecture Pattern: Layered Architecture

->FEATURES IMPLEMENTED

Employee Management APIs

1. Create Employees
2. Get All Employees
3. Get Employee by ID
4. Update Employee
5. Delete Employee

 Onboarding Management APIs

6. Create Onboarding
7. Approve Onboarding
8. Reject Onboarding
9. Get All Onboardings

This project intentionally uses core Golang features expected in real backend roles.

1. Go Modules (Dependency Management)

Used go.mod and go.sum for managing dependencies
Ensures version consistency and reproducible builds
Demonstrates modern Go development practices

2. Packages & Project Organization

Each folder represents a logical responsibility:
Package	Responsibility

main	      Application entry point
handler	      HTTP request handling
libhttp	      Route registration
models	      Data structures
repository	  Business & DB logic
worker	      Background processing

3. Structs for Data Modeling

Employees and Onboardings are modeled using Go structs
Strong typing ensures data safety
JSON tags used for API communication
Time handling using Go’s time package

4. Interfaces 

Repository layer is defined using interfaces
Database implementation satisfies these interfaces
Handlers depend on interfaces, not concrete implementation

5. Dependency Injection

Repository implementation is injected into handlers
Prevents tight coupling
Enables future extensibility (e.g., mock repositories)

6. Database Integration (database/sql)

Direct interaction with MySQL using Go’s standard library
Proper handling of:
Connections
Queries
Result scanning
Rows affected
SQL errors

7. RESTful API Design (Gin Framework)

Clear HTTP verbs: GET, POST, PUT, DELETE
Clean endpoint naming
Path parameters and request bodies
JSON-based communication

8.Error Handling 

Explicit error checks at every layer
Meaningful HTTP error responses

-> API TESTING

All APIs tested using Postman

Verified:
Request validation
JSON responses
HTTP status codes
Database persistence
APIs tested end-to-end with real data

-> HOW TO RUN THE APPLICATION

1. Ensure MySQL is running
2. Create required database and tables
3. Update database credentials
4. Run the application:

go run main.go
Server starts on:
http://localhost:8081


API Testing Using Postman

All APIs in this project are tested using Postman.
Each endpoint follows REST principles and returns JSON responses with proper HTTP status codes.

Base URL:
http://localhost:8081

->POSTMAN TESTING GUIDE

 Employee APIs

1. Create Employee

Method: POST
Endpoint:/api/employees
Postman Setup:
Body → raw → JSON
Provide employee  details (name, email, department, salary)

Expected Result:
Status: 201 Created
Employee stored in database

2. Get All Employees

Method: GET
Endpoint:
/api/employees
Postman Setup:
No body required

Expected Result:
Status: 200 OK
JSON array of all employees

3.Get Employee by ID

Method: GET
Endpoint:
/api/employees/{id}

Expected Result:
Status: 200 OK
JSON object of the selected employee

4.Update Employee

Method: PUT
Endpoint:
/api/employees/{id}
Postman Setup:
Body → raw → JSON
Update one or more fields

Expected Result:
Status: 200 OK
Employee data updated in database

5. Delete Employee

Method: DELETE
Endpoint:
/api/employees/{id}
Postman Setup:
No body required

Expected Result:
Status: 200 OK
Employee removed from database

 ONBOARDING APIs 

6. Create Onboarding

Method: POST
Endpoint:
/api/onboarding
Postman Setup:
Body → raw → JSON
Provide employee ID and onboarding details

Expected Result:
Status: 201 Created
Onboarding entry created with default status

7. Approve Onboarding

Method: PUT
Endpoint:
/api/onboarding/{id}/approve
Postman Setup:
Replace {id} with onboarding ID

Expected Result:
Status updated to APPROVED
Status: 200 OK

8. Reject Onboarding

Method: PUT
Endpoint:
/api/onboarding/{id}/reject
Postman Setup:
Replace {id} with onboarding ID

Expected Result:
Status updated to REJECTED
Status: 200 OK

9. Get All Onboardings

Method: GET
Endpoint:
/api/onboarding
Postman Setup:
No body required

Expected Result:
Status: 200 OK
JSON array of all onboarding records


PROJECT FOLDER STRUCTURE:

employee-management/
│
├── main.go
├── go.mod
├── go.sum
│
├── service/
│   ├── handler/
│   │   └── handler.go
│   │
│   ├── libhttp/
│   │   └── routes.go
│   │
│   ├── models/
│   │   ├── models.go
│   │   └── onboardings.go
│   │
│   ├── repository/
│   │   ├── interface.go
│   │   └── db/
│   │       └── db.go
│   │
│   └── worker/
│       └── onboarding_worker.go


FINAL SUMMARY

-> What This Project Demonstrates 

Strong Golang fundamentals
Clean backend architecture
Database-driven API development
Debugging and problem-solving skills
Production-oriented thinking

THANKYOU.


