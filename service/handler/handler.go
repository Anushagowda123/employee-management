package handler

import (
	"net/http"
	"strconv"

	"github.com/Anushagowda123/employee-management/service/models"
	"github.com/Anushagowda123/employee-management/service/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repo repository.EmployeeRepository
}

func NewHandler(repo repository.EmployeeRepository) *Handler {
	return &Handler{Repo: repo}
}


//EMPLOYEE APIs 


// POST /api/employees
func (h *Handler) CreateEmployee(c *gin.Context) {
	var emp models.Employee

	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Repo.CreateEmployee(&emp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "employee created"})
}

// GET /api/employees
func (h *Handler) GetEmployees(c *gin.Context) {
	employees, err := h.Repo.GetEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}

// GET /api/employees/:id
func (h *Handler) GetEmployeeByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	emp, err := h.Repo.GetEmployeeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}

	c.JSON(http.StatusOK, emp)
}

// PUT /api/employees/:id
func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var emp models.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Repo.UpdateEmployee(id, &emp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "employee updated"})
}

// DELETE /api/employees/:id
func (h *Handler) DeleteEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.Repo.DeleteEmployee(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "employee deleted"})
}


//ONBOARDING APIs 


// POST /api/onboarding
func (h *Handler) CreateOnboarding(c *gin.Context) {
	var ob models.Onboarding

	if err := c.ShouldBindJSON(&ob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Repo.CreateOnboarding(&ob); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "onboarding created"})
}

// PUT /api/onboarding/:id/approve
func (h *Handler) ApproveOnboarding(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.Repo.ApproveOnboarding(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "onboarding approved"})
}

// PUT /api/onboarding/:id/reject
func (h *Handler) RejectOnboarding(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.Repo.RejectOnboarding(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "onboarding rejected"})
}

// GET /api/onboarding
func (h *Handler) ListOnboardings(c *gin.Context) {
	onboardings, err := h.Repo.GetOnboardings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, onboardings)
}