// src/handlers/doctor.go
package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kingslyDev/telemedicine/server/config"
	"github.com/kingslyDev/telemedicine/server/models"
)

// UpdateDoctorProfileInput defines the expected input for updating doctor profile
type UpdateDoctorProfileInput struct {
	FirstName         string `json:"first_name" binding:"required"`
	LastName          string `json:"last_name" binding:"required"`
	Specialization    string `json:"specialization"`
	LicenseNumber     string `json:"license_number"`
	YearsOfExperience int    `json:"years_of_experience"`
	Bio               string `json:"bio"`
	AvailableHours    string `json:"available_hours"` // JSON string
}

// GetDoctorProfileHandler retrieves the current doctor's profile
func GetDoctorProfileHandler(c *gin.Context) {
	// Ensure the user has the 'doctor' role
	role, exists := c.Get("role")
	if !exists || role != "doctor" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Get the user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	// Find the doctor associated with the user ID
	var doctor models.Doctor
	if err := config.DB.Where("user_id = ?", userID).First(&doctor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor profile not found"})
		return
	}

	// Respond with the doctor profile
	c.JSON(http.StatusOK, gin.H{"doctor": doctor})
}

// UpdateDoctorProfileHandler handles updating a doctor's profile
func UpdateDoctorProfileHandler(c *gin.Context) {
	// Ensure the user has the 'doctor' role
	role, exists := c.Get("role")
	if !exists || role != "doctor" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Get the user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	// Bind the input
	var input UpdateDoctorProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the doctor associated with the user ID
	var doctor models.Doctor
	if err := config.DB.Where("user_id = ?", userID).First(&doctor).Error; err != nil {
		// If doctor profile does not exist, create one
		doctor = models.Doctor{
			UserID:            userID.(uint),
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create doctor profile"})
			return
		}
	} else {
		// Update existing doctor fields
		doctor.FirstName = input.FirstName
		doctor.LastName = input.LastName
		doctor.Specialization = input.Specialization
		doctor.LicenseNumber = input.LicenseNumber
		doctor.YearsOfExperience = input.YearsOfExperience
		doctor.Bio = input.Bio
		doctor.AvailableHours = input.AvailableHours
		doctor.UpdatedAt = time.Now()

		// Save the updated doctor
		if err := config.DB.Save(&doctor).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update doctor profile"})
			return
		}
	}

	// Respond with the updated profile
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully", "doctor": doctor})
}
