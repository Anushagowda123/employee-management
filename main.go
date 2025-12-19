package main

import (
	"database/sql"
	"log"

	"github.com/Anushagowda123/employee-management/service/handler"
	"github.com/Anushagowda123/employee-management/service/libhttp"
	"github.com/Anushagowda123/employee-management/service/repository/db"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/employee_management?parseTime=true"

	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := dbConn.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB connected successfully")

	repo := &db.EmployeeRepo{DB: dbConn}

	h := handler.NewHandler(repo)

	r := gin.Default()

	libhttp.RegisterRoutes(r, h)

	log.Println("Server running on :8081")
	r.Run(":8081")
}