package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kingslyDev/telemedicine/server/handlers"
	"github.com/kingslyDev/telemedicine/server/middleware"
)

func RegisterProfileRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// Patient routes
			protected.GET("/patient/profile", handlers.GetPatientProfileHandler)
			protected.PUT("/patient/profile", handlers.UpdatePatientProfileHandler)

			// Doctor routes
			protected.GET("/doctor/profile", handlers.GetDoctorProfileHandler)
			protected.PUT("/doctor/profile", handlers.UpdateDoctorProfileHandler)

			// Admin routes
			protected.GET("/admin/profile", handlers.GetAdminProfileHandler)
			protected.PUT("/admin/profile", handlers.UpdateAdminProfileHandler)
		}
	}
}
