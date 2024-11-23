package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kingslyDev/telemedicine/server/config"
	"github.com/kingslyDev/telemedicine/server/models"
)

func main() {
    // Memuat file .env jika ada
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

    // Inisialisasi database
    config.InitDB()

    // Aktifkan mode debug GORM (opsional, untuk debugging)
    config.DB = config.DB.Debug()

    // Lakukan migrasi database dengan urutan yang benar
	err = config.DB.AutoMigrate(
		&models.User{},
		&models.Patient{},
		&models.Doctor{},
		&models.Staff{},
		&models.Admin{},
		&models.Appointment{},
		&models.DoctorSchedule{},
		&models.MedicalRecord{},
		&models.MedicalImage{},
		&models.DataMiningResult{},
		&models.LabResult{},
		&models.Notification{},
		&models.AccessControl{},
	)
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    log.Println("Database migration completed")

    // Inisialisasi router Gin
    router := gin.Default()

    // Mengaktifkan CORS untuk mengizinkan akses dari domain lain
    router.Use(cors.Default())

    // Route sederhana untuk pengujian
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    // Tambahkan route lain di sini (misalnya, registrasi atau login)
    // router.POST("/register", RegisterHandler)

    // Menjalankan server di port 8080
    router.Run(":8080")
}
