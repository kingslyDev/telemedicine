package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kingslyDev/telemedicine/server/config"
	"github.com/kingslyDev/telemedicine/server/models"
	"golang.org/x/crypto/bcrypt"
)

// Struktur input untuk login
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"` // Validasi tambahan untuk memastikan format email
	Password string `json:"password" binding:"required"`    // Password wajib diisi
}

// Handler untuk login
func LoginHandler(c *gin.Context) {
	var input LoginInput
	// Bind JSON dari request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cari user berdasarkan email
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"}) // Email tidak ditemukan
		return
	}

	// Bandingkan password yang di-hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"}) // Password salah
		return
	}

	// Pastikan variabel lingkungan `JWT_SECRET` ada
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT secret is not configured"})
		return
	}

	// Buat JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,                         // ID pengguna dari database
		"email":    user.Email,                      // Email pengguna
		"role":     user.Role,                       // Role pengguna (misal: patient, doctor)
		"exp":      time.Now().Add(72 * time.Hour).Unix(), // Token berlaku selama 72 jam
	})

	// Tanda tangani token dengan secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Kirim token ke client
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
