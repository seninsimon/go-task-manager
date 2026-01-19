package handlers

import (
	"net/http"
	"strconv"
	"task-manager/internal/services"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *services.TaskService
}


func NewTaskHandler() *TaskHandler{
	return &TaskHandler{
		service: services.NewTaskService(),
	}
}

func (h *TaskHandler) Create(c *gin.Context) {
	var req struct {
		Title string `json:"title"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userID := c.GetUint("user_id")

	if err := h.service.Create(req.Title, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "task is created"})
}

func (h *TaskHandler) GetAll(c *gin.Context) {

	userId := c.GetUint("user_id")

	tasks, err := h.service.GetAll(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)

}

func (h *TaskHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		Title  string `json:"title"`
		Status string `json:"status"`
	}

	c.ShouldBindJSON(&req)

	userID := c.GetUint("user_id")
	
	if err := h.service.Update(uint(id) ,userID,  req.Title , req.Status ); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task updated"})
}


func (h *TaskHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	userId := c.GetUint("user_id")

	if err := h.service.Delete(uint(id), userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
