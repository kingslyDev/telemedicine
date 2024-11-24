package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kingslyDev/telemedicine/server/config"
	"github.com/kingslyDev/telemedicine/server/routes"
)

func main() {
    // Inisialisasi database
    config.InitDB()

    // Inisialisasi router Gin
    router := gin.Default()

    // Middleware CORS
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // Ganti sesuai URL frontend Anda
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    // Register routes
    routes.RegisterAuthRoutes(router)

    // Menjalankan server di port 8080
    router.Run(":8080")
}