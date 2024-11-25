package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kingslyDev/telemedicine/server/handlers/auth"
)

func RegisterAuthRoutes(r *gin.Engine) {	
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", auth.RegisterHandler)
		authRoutes.POST("/login", auth.LoginHandler)
	}
}
