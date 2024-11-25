// src/handlers/patient.go
package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kingslyDev/telemedicine/server/config"
	"github.com/kingslyDev/telemedicine/server/models"
)

// UpdatePatientProfileInput defines the expected input for updating patient profile
type UpdatePatientProfileInput struct {
	FirstName        string    `json:"first_name" binding:"required"`
	LastName         string    `json:"last_name" binding:"required"`
	DateOfBirth      string    `json:"date_of_birth" binding:"required"` // Expecting "YYYY-MM-DD"
	Gender           string    `json:"gender" binding:"required"`
	Address          string    `json:"address"`
	MedicalHistory   string    `json:"medical_history"`
	EmergencyContact string    `json:"emergency_contact"`
	BloodType        string    `json:"blood_type"`
	Allergies        string    `json:"allergies"`
}

// GetPatientProfileHandler retrieves the current patient's profile
func GetPatientProfileHandler(c *gin.Context) {
	// Ensure the user has the 'patient' role
	role, exists := c.Get("role")
	if !exists || role != "patient" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Get the user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	// Find the patient associated with the user ID
	var patient models.Patient
	if err := config.DB.Where("user_id = ?", userID).First(&patient).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient profile not found"})
		return
	}

	// Respond with the patient profile
	c.JSON(http.StatusOK, gin.H{"patient": patient})
}

// UpdatePatientProfileHandler handles updating a patient's profile
func UpdatePatientProfileHandler(c *gin.Context) {
	// Ensure the user has the 'patient' role
	role, exists := c.Get("role")
	if !exists || role != "patient" {
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
	var input UpdatePatientProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse DateOfBirth
	dob, err := time.Parse("2006-01-02", input.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}

	// Find the patient associated with the user ID
	var patient models.Patient
	if err := config.DB.Where("user_id = ?", userID).First(&patient).Error; err != nil {
		// If patient profile does not exist, create one
		patient = models.Patient{
			UserID:          userID.(uint),
			FirstName:       input.FirstName,
			LastName:        input.LastName,
			DateOfBirth:     dob,
			Gender:          input.Gender,
			Address:         input.Address,
			MedicalHistory:  input.MedicalHistory,
			EmergencyContact: input.EmergencyContact,
			BloodType:       input.BloodType,
			Allergies:       input.Allergies,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}
		if err := config.DB.Create(&patient).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient profile"})
			return
		}
	} else {
		// Update existing patient fields
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

		// Save the updated patient
		if err := config.DB.Save(&patient).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient profile"})
			return
		}
	}

	// Respond with the updated profile
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully", "patient": patient})
}
