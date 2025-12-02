package libhttp

import (
	"employee-management/service/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	h := handler.NewEmployeeHandler()
	grp := r.Group("/employees")
	{
		grp.POST("/", h.CreateEmployee)
		grp.GET("/", h.GetAllEmployees)
		grp.GET("/:id", h.GetEmployeeByID)
		grp.DELETE("/:id", h.DeleteEmployee)
	}
}
