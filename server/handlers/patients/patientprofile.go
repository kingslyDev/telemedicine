// src/handlers/patient/patientprofile.go
package patient

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kingslyDev/telemedicine/server/config"
	"github.com/kingslyDev/telemedicine/server/handlers/helpers"
	"github.com/kingslyDev/telemedicine/server/models"
)

// UpdatePatientProfileInput mendefinisikan input yang diharapkan untuk memperbarui profil pasien
type UpdatePatientProfileInput struct {
	FirstName        string `json:"first_name" binding:"required"`
	LastName         string `json:"last_name" binding:"required"`
	DateOfBirth      string `json:"date_of_birth" binding:"required"` // Format: "YYYY-MM-DD"
	Gender           string `json:"gender" binding:"required"`
	Address          string `json:"address"`
	MedicalHistory   string `json:"medical_history"`
	EmergencyContact string `json:"emergency_contact"`
	BloodType        string `json:"blood_type"`
	Allergies        string `json:"allergies"`
}

// GetProfileHandler mengambil profil pasien saat ini
func GetProfileHandler(c *gin.Context) {
	// Verifikasi peran pengguna
	if !helpers.GetUserRole(c, "patient") {
		return
	}

	// Ambil user_id dari konteks
	userID, ok := helpers.GetUserID(c)
	if !ok {
		return
	}

	// Cari profil pasien berdasarkan user_id
	var patient models.Patient
	if err := config.DB.Where("user_id = ?", userID).First(&patient).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profil pasien tidak ditemukan"})
		return
	}

	// Kembalikan data profil pasien
	c.JSON(http.StatusOK, gin.H{"patient": patient})
}

// UpdateProfileHandler memperbarui profil pasien
func UpdateProfileHandler(c *gin.Context) {
	// Verifikasi peran pengguna
	if !helpers.GetUserRole(c, "patient") {
		return
	}

	// Ambil user_id dari konteks
	userID, ok := helpers.GetUserID(c)
	if !ok {
		return
	}

	// Bind data JSON ke struct input
	var input UpdatePatientProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse DateOfBirth
	dob, err := time.Parse("2006-01-02", input.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format tanggal lahir tidak valid. Gunakan YYYY-MM-DD."})
		return
	}

	// Cari profil pasien
	var patient models.Patient
	if err := config.DB.Where("user_id = ?", userID).First(&patient).Error; err != nil {
		// Jika profil pasien tidak ada, buat baru
		patient = models.Patient{
			UserID:            userID,
			FirstName:         input.FirstName,
			LastName:          input.LastName,
			DateOfBirth:       dob,
			Gender:            input.Gender,
			Address:           input.Address,
			MedicalHistory:    input.MedicalHistory,
			EmergencyContact:  input.EmergencyContact,
			BloodType:         input.BloodType,
			Allergies:         input.Allergies,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}
		if err := config.DB.Create(&patient).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat profil pasien"})
			return
		}
	} else {
		// Perbarui profil pasien yang ada
		patient.FirstName = input.FirstName
		patient.LastName = input.LastName
		patient.DateOfBirth = dob
		patient.Gender = input.Gender
		patient.Address = input.Address
		patient.MedicalHistory = input.MedicalHistory
		patient.EmergencyContact = input.EmergencyContact
		patient.BloodType = input.BloodType
		patient.Allergies = input.Allergies
		patient.UpdatedAt = time.Now()

		// Simpan perubahan
		if err := config.DB.Save(&patient).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui profil pasien"})
			return
		}
	}

	// Kembalikan respon sukses
	c.JSON(http.StatusOK, gin.H{"message": "Profil berhasil diperbarui", "patient": patient})
}
