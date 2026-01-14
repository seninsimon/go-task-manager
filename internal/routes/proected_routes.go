package routes

import (
	"net/http"
	"task-manager/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterProtectedRoutes(r *gin.Engine) {

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/me" , func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK , gin.H{
				"user_id" : ctx.GetUint("user_id"),
				"email" : ctx.GetString("email"),
			})
		})
	}

}