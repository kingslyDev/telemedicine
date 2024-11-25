package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kingslyDev/telemedicine/server/config"
	"github.com/kingslyDev/telemedicine/server/middleware"
	"github.com/kingslyDev/telemedicine/server/routes"
)

func main() {
	// Muat file .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or could not be loaded")
	}

	// Pastikan JWT_SECRET tersedia
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not configured. Please set it in .env file.")
	}

	// Inisialisasi database
	config.InitDB()

	// Inisialisasi router Gin
	router := gin.Default()

	// Middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // URL frontend Anda
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Register public routes (tanpa autentikasi)
	routes.RegisterAuthRoutes(router)

	// ROUTE PROFILE
	routes.RegisterProfileRoutes(router)

	// Group routes dengan middleware JWT
	protected := router.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware()) // Middleware JWT
	{
		// Tambahkan protected routes
		protected.GET("/profile", func(c *gin.Context) {
			userID := c.GetString("user_id")
			email := c.GetString("email")
			role := c.GetString("role")

			c.JSON(200, gin.H{
				"user_id": userID,
				"email":   email,
				"role":    role,
			})
		})

		// Contoh route lain yang dilindungi
		protected.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome to the dashboard!"})
		})
	}

	// Jalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port jika tidak diatur di .env
	}
	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}
