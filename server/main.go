package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kingslyDev/telemedicine/server/config"
	"github.com/kingslyDev/telemedicine/server/models"
	"github.com/kingslyDev/telemedicine/server/routes"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or could not be loaded")
	}

	// Ensure JWT_SECRET is configured
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not configured. Please set it in .env file.")
	}

	// Initialize database
	// Initialize database
	config.InitDB()

	// Automatically migrate tables to reflect model changes
	if err := config.DB.AutoMigrate(
		&models.User{},
		&models.Patient{},
		&models.MedicalRecord{},
		&models.Appointment{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed successfully!")

	// Initialize Gin router
	router := gin.Default()

	// CORS Middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Register public routes (no authentication required)
	routes.RegisterAuthRoutes(router)

	// Public route: Fetch all patients
	router.GET("/api/patients", func(c *gin.Context) {
		log.Println("Fetching all patients...")
		var patients []models.Patient
		if err := config.DB.Preload("User").Find(&patients).Error; err != nil {
			log.Printf("Failed to fetch patients: %v", err)
			c.JSON(500, gin.H{"error": "Failed to fetch patients", "details": err.Error()})
			return
		}

		log.Printf("Fetched %d patients", len(patients))
		c.JSON(200, gin.H{"patients": patients})
	})

	router.GET("/api/patients/:id", func(c *gin.Context) {
		id := c.Param("id")
		log.Printf("Fetching details for patient ID: %s", id)

		var patient models.Patient
		if err := config.DB.Preload("User").
			Preload("MedicalRecords.Doctor").
			Preload("MedicalRecords.Appointment").
			First(&patient, "id = ?", id).Error; err != nil {
			log.Printf("Error fetching patient: %v", err)
			c.JSON(404, gin.H{"error": "Patient not found"})
			return
		}

		log.Printf("Successfully fetched patient: %+v", patient)
		c.JSON(200, gin.H{"patient": patient})
	})

	// Protected routes (require authentication)
	protected := router.Group("/api")
	protected.Use(func(c *gin.Context) {
		// JWT authentication logic
		role := c.GetString("role") // Assume the role is set by JWT middleware
		if role == "" {
			log.Println("Unauthorized access attempt")
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	})
	{
		// Protected route: Profile
		protected.GET("/profile", func(c *gin.Context) {
			userID := c.GetString("user_id")
			email := c.GetString("email")
			role := c.GetString("role")

			log.Printf("Profile accessed by user ID: %s", userID)
			c.JSON(200, gin.H{
				"user_id": userID,
				"email":   email,
				"role":    role,
			})
		})

		// Protected route: Dashboard
		protected.GET("/dashboard", func(c *gin.Context) {
			log.Println("Dashboard accessed")
			c.JSON(200, gin.H{"message": "Welcome to the dashboard!"})
		})
	}

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}
