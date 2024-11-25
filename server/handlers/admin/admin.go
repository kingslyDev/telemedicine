// src/handlers/admin.go
package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kingslyDev/telemedicine/server/config"
	"github.com/kingslyDev/telemedicine/server/models"
)

// UpdateAdminProfileInput defines the expected input for updating admin profile
type UpdateAdminProfileInput struct {
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Privileges string `json:"privileges"` // JSON string
}

// GetAdminProfileHandler retrieves the current admin's profile
func GetAdminProfileHandler(c *gin.Context) {
	// Ensure the user has the 'admin' role
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Get the user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	// Find the admin associated with the user ID
	var admin models.Admin
	if err := config.DB.Where("user_id = ?", userID).First(&admin).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin profile not found"})
		return
	}

	// Respond with the admin profile
	c.JSON(http.StatusOK, gin.H{"admin": admin})
}

// UpdateAdminProfileHandler handles updating an admin's profile
func UpdateAdminProfileHandler(c *gin.Context) {
	// Ensure the user has the 'admin' role
	role, exists := c.Get("role")
	if !exists || role != "admin" {
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
	var input UpdateAdminProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the admin associated with the user ID
	var admin models.Admin
	if err := config.DB.Where("user_id = ?", userID).First(&admin).Error; err != nil {
		// If admin profile does not exist, create one
		admin = models.Admin{
			UserID:     userID.(uint),
			FirstName:  input.FirstName,
			LastName:   input.LastName,
			Privileges: input.Privileges,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		if err := config.DB.Create(&admin).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin profile"})
			return
		}
	} else {
		// Update existing admin fields
		admin.FirstName = input.FirstName
		admin.LastName = input.LastName
		admin.Privileges = input.Privileges
		admin.UpdatedAt = time.Now()

		// Save the updated admin
		if err := config.DB.Save(&admin).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update admin profile"})
			return
		}
	}

	// Respond with the updated profile
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully", "admin": admin})
}
