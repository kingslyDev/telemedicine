// src/handlers/helpers/helpers.go
package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserID retrieves the user_id from the context and converts it to uint
func GetUserID(c *gin.Context) (uint, bool) {
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return 0, false
	}

	userID, ok := toUint(userIDInterface)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return 0, false
	}

	return userID, true
}

// GetUserRole retrieves the role from the context and ensures it matches the expected role
func GetUserRole(c *gin.Context, expectedRole string) bool {
	role, exists := c.Get("role")
	if !exists || role != expectedRole {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return false
	}
	return true
}

// Helper function to convert interface{} to uint
func toUint(value interface{}) (uint, bool) {
	switch v := value.(type) {
	case float64:
		return uint(v), true
	case uint:
		return v, true
	case int:
		return uint(v), true
	default:
		return 0, false
	}
}
