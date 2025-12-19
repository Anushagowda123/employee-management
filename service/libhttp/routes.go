package libhttp

import (
	"github.com/Anushagowda123/employee-management/service/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *handler.Handler) {

	api := r.Group("/api")
	{
		api.POST("/employees", h.CreateEmployee)
		api.GET("/employees", h.GetEmployees)
		api.GET("/employees/:id", h.GetEmployeeByID)
		api.PUT("/employees/:id", h.UpdateEmployee)
		api.DELETE("/employees/:id", h.DeleteEmployee)

		api.POST("/onboarding", h.CreateOnboarding)
		api.PUT("/onboarding/:id/approve", h.ApproveOnboarding)
		api.PUT("/onboarding/:id/reject", h.RejectOnboarding)
		api.GET("/onboarding", h.ListOnboardings)
	}
}
