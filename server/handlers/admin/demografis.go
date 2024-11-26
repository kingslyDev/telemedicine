package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kingslyDev/telemedicine/server/config"
	"github.com/kingslyDev/telemedicine/server/models"
)

func GetPatients(c *gin.Context) {
    var patients []models.Patient

    log.Println("Attempting to fetch patients from the database...")

    // Query with preloads
    if err := config.DB.
        Preload("User").
        Preload("MedicalRecords.Doctor").
        Preload("MedicalRecords.Appointment").
        Find(&patients).Error; err != nil {
        log.Printf("Database error: %v", err) // Log the error for debugging
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to fetch patients: " + err.Error(),
        })
        return
    }

    log.Printf("Successfully fetched %d patients", len(patients))
    c.JSON(http.StatusOK, gin.H{"patients": patients})
}

func GetPatientDetails(c *gin.Context) {
    patientID := c.Param("id")

    var patient models.Patient

    // Preload User, MedicalRecords, Doctor, and Appointment relationships
    if err := config.DB.
        Preload("User").
        Preload("MedicalRecords.Doctor").
        Preload("MedicalRecords.Appointment").
        First(&patient, patientID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
            return
        }
        log.Printf("Failed to fetch patient details: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patient details"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"patient": patient})
}
