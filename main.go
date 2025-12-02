package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"employee-management/service/libhttp"
	"employee-management/service/repository/db"
)

func main() {
	// connect to DB (will fatal log on error)
	db.Connect()

	// create gin router and register routes
	router := gin.Default()
	libhttp.RegisterRoutes(router)

	log.Println("Server running on :8081")
	router.Run(":8081")
}
