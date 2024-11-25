// src/handlers/admin/adminprofile.go
package admin

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kingslyDev/telemedicine/server/config"
	"github.com/kingslyDev/telemedicine/server/handlers/helpers"
	"github.com/kingslyDev/telemedicine/server/models"
)

// UpdateAdminProfileInput mendefinisikan input yang diharapkan untuk memperbarui profil admin
type UpdateAdminProfileInput struct {
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Privileges string `json:"privileges" binding:"required"` // JSON string
}

// GetProfileHandler mengambil profil admin saat ini
func GetProfileHandler(c *gin.Context) {
	// Verifikasi peran pengguna
	if !helpers.GetUserRole(c, "admin") {
		return
	}

	// Ambil user_id dari konteks
	userID, ok := helpers.GetUserID(c)
	if !ok {
		return
	}

	// Cari profil admin berdasarkan user_id
	var admin models.Admin
	if err := config.DB.Where("user_id = ?", userID).First(&admin).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profil admin tidak ditemukan"})
		return
	}

	// Kembalikan data profil admin
	c.JSON(http.StatusOK, gin.H{"admin": admin})
}

// UpdateProfileHandler memperbarui profil admin
func UpdateProfileHandler(c *gin.Context) {
	// Verifikasi peran pengguna
	if !helpers.GetUserRole(c, "admin") {
		return
	}

	// Ambil user_id dari konteks
	userID, ok := helpers.GetUserID(c)
	if !ok {
		return
	}

	// Bind data JSON ke struct input
	var input UpdateAdminProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi bahwa Privileges adalah JSON yang valid
	var privilegesMap map[string]interface{}
	if err := json.Unmarshal([]byte(input.Privileges), &privilegesMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Privileges harus berupa string JSON yang valid"})
		return
	}

	// Cari profil admin
	var admin models.Admin
	if err := config.DB.Where("user_id = ?", userID).First(&admin).Error; err != nil {
		// Jika profil admin tidak ada, buat baru
		admin = models.Admin{
			UserID:     userID,
			FirstName:  input.FirstName,
			LastName:   input.LastName,
			Privileges: input.Privileges,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		if err := config.DB.Create(&admin).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat profil admin"})
			return
		}
	} else {
		// Perbarui profil admin yang ada
		admin.FirstName = input.FirstName
		admin.LastName = input.LastName
		admin.Privileges = input.Privileges
		admin.UpdatedAt = time.Now()

		// Simpan perubahan
		if err := config.DB.Save(&admin).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui profil admin"})
			return
		}
	}

	// Kembalikan respon sukses
	c.JSON(http.StatusOK, gin.H{"message": "Profil berhasil diperbarui", "admin": admin})
}
