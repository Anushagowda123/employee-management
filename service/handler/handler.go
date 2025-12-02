package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"employee-management/service/models"
	"employee-management/service/repository"
	"employee-management/service/repository/db"
)

type EmployeeHandler struct {
	Repo repository.EmployeeRepository
}

func NewEmployeeHandler() *EmployeeHandler {
	return &EmployeeHandler{Repo: db.NewEmployeeRepo()}
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var emp models.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Repo.CreateEmployee(emp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create failed"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}

func (h *EmployeeHandler) GetAllEmployees(c *gin.Context) {
	list, err := h.Repo.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fetch failed"})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *EmployeeHandler) GetEmployeeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	emp, err := h.Repo.GetEmployeeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, emp)
}

func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Repo.DeleteEmployee(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
