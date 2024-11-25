// src/handlers/doctor/doctorprofile.go
package doctor

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kingslyDev/telemedicine/server/config"
	"github.com/kingslyDev/telemedicine/server/handlers/helpers"
	"github.com/kingslyDev/telemedicine/server/models"
)

// UpdateDoctorProfileInput mendefinisikan input yang diharapkan untuk memperbarui profil dokter
type UpdateDoctorProfileInput struct {
	FirstName         string `json:"first_name" binding:"required"`
	LastName          string `json:"last_name" binding:"required"`
	Specialization    string `json:"specialization"`
	LicenseNumber     string `json:"license_number"`
	YearsOfExperience int    `json:"years_of_experience"`
	Bio               string `json:"bio"`
	AvailableHours    string `json:"available_hours"` // JSON string
}

// GetProfileHandler mengambil profil dokter saat ini
func GetProfileHandler(c *gin.Context) {
	// Verifikasi peran pengguna
	if !helpers.GetUserRole(c, "doctor") {
		return
	}

	// Ambil user_id dari konteks
	userID, ok := helpers.GetUserID(c)
	if !ok {
		return
	}

	// Cari profil dokter berdasarkan user_id
	var doctor models.Doctor
	if err := config.DB.Where("user_id = ?", userID).First(&doctor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profil dokter tidak ditemukan"})
		return
	}

	// Kembalikan data profil dokter
	c.JSON(http.StatusOK, gin.H{"doctor": doctor})
}

// UpdateProfileHandler memperbarui profil dokter
func UpdateProfileHandler(c *gin.Context) {
	// Verifikasi peran pengguna
	if !helpers.GetUserRole(c, "doctor") {
		return
	}

	// Ambil user_id dari konteks
	userID, ok := helpers.GetUserID(c)
	if !ok {
		return
	}

	// Bind data JSON ke struct input
	var input UpdateDoctorProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cari profil dokter
	var doctor models.Doctor
	if err := config.DB.Where("user_id = ?", userID).First(&doctor).Error; err != nil {
		// Jika profil dokter tidak ada, buat baru
		doctor = models.Doctor{
			UserID:            userID,
			FirstName:         input.FirstName,
			LastName:          input.LastName,
			Specialization:    input.Specialization,
			LicenseNumber:     input.LicenseNumber,
			YearsOfExperience: input.YearsOfExperience,
			Bio:               input.Bio,
			AvailableHours:    input.AvailableHours,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}
		if err := config.DB.Create(&doctor).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat profil dokter"})
			return
		}
	} else {
		// Perbarui profil dokter yang ada
		doctor.FirstName = input.FirstName
		doctor.LastName = input.LastName
		doctor.Specialization = input.Specialization
		doctor.LicenseNumber = input.LicenseNumber
		doctor.YearsOfExperience = input.YearsOfExperience
		doctor.Bio = input.Bio
		doctor.AvailableHours = input.AvailableHours
		doctor.UpdatedAt = time.Now()

		// Simpan perubahan
		if err := config.DB.Save(&doctor).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui profil dokter"})
			return
		}
	}

	// Kembalikan respon sukses
	c.JSON(http.StatusOK, gin.H{"message": "Profil berhasil diperbarui", "doctor": doctor})
}
