package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

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
