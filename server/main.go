package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kingslyDev/telemedicine/server/config"
)

func main() {
	
    router := gin.Default()

	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

    // Inisialisasi database
    config.InitDB()

    // Mengaktifkan CORS
    router.Use(cors.Default())

    // Route sederhana untuk pengujian
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    // Menjalankan server di port 8080
    router.Run(":8080")
}
