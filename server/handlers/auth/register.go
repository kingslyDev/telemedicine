package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kingslyDev/telemedicine/server/config"
	"github.com/kingslyDev/telemedicine/server/models"
	"golang.org/x/crypto/bcrypt"
)

// Input untuk registrasi
type RegisterInput struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role" binding:"required,oneof=patient doctor staff admin"`
}

// Handler untuk registrasi
func RegisterHandler(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Buat user
	user := models.User{
		Username:     input.Username,
		PasswordHash: string(hashedPassword),
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		Role:         input.Role,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Simpan user ke database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Buat entitas spesifik berdasarkan role
	switch input.Role {
	case "doctor":
		doctor := models.Doctor{
			UserID: user.ID, // Foreign key ke User
			// Tambahkan atribut default lain jika diperlukan
		}
		if err := config.DB.Create(&doctor).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create doctor"})
			return
		}
	case "admin":
		admin := models.Admin{
			UserID: user.ID, // Foreign key ke User
			// Tambahkan atribut default lain jika diperlukan
		}
		if err := config.DB.Create(&admin).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin"})
			return
		}
	case "staff":
		staff := models.Staff{
			UserID: user.ID, // Foreign key ke User
			// Tambahkan atribut default lain jika diperlukan
		}
		if err := config.DB.Create(&staff).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create staff"})
			return
		}
	case "patient":
		patient := models.Patient{
			UserID: user.ID, // Foreign key ke User
			// Tambahkan atribut default lain jika diperlukan
		}
		if err := config.DB.Create(&patient).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient"})
			return
		}
	default:
		// Jika role tidak valid, hapus user yang sudah dibuat
		config.DB.Delete(&user)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	// Jika berhasil
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
