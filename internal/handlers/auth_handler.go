package handlers

import (
	"net/http"
	"task-manager/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	services *services.UserService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		services: services.NewUserService(),
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := h.services.Register(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"name":  user.Name,
			"email": user.Email,
		},
	})

}


func (h *AuthHandler) Login(c *gin.Context){
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

    if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest , gin.H{
			"error": "invalid request body",
		})
		return
	}

	token , err := h.services.Login(req.Email , req.Password)
	if err!= nil {
		c.JSON(http.StatusUnauthorized , gin.H{
			"error" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK , gin.H{
		"token" : token,
	})

}
