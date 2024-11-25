// src/routes/profile.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kingslyDev/telemedicine/server/handlers/admin"
	"github.com/kingslyDev/telemedicine/server/handlers/doctor"
	"github.com/kingslyDev/telemedicine/server/handlers/patient"
	"github.com/kingslyDev/telemedicine/server/middleware"
)

func RegisterProfileRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Rute yang dilindungi dengan middleware JWT
		protected := api.Group("/")
		protected.Use(middleware.JWTAuthMiddleware())
		{
			// Rute Pasien
			protected.GET("/patient/profile", patient.GetProfileHandler)
			protected.PUT("/patient/profile", patient.UpdateProfileHandler)

			// Rute Dokter
			protected.GET("/doctor/profile", doctor.GetProfileHandler)
			protected.PUT("/doctor/profile", doctor.UpdateProfileHandler)

			// Rute Admin
			protected.GET("/admin/profile", admin.GetProfileHandler)
			protected.PUT("/admin/profile", admin.UpdateProfileHandler)
		}
	}
}
