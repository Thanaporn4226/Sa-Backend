package controller

import (
	"net/http"

	"github.com/Pet002/Project-sa-65/entity"
	"github.com/gin-gonic/gin"
)

// POST /prescriptions
func CreatePrescription(c *gin.Context) {
	var prescription entity.Prescription
	var patient entity.Patient
	var medicine entity.Medicine
	var employee entity.Employee

	if err := c.ShouldBindJSON(&prescription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", prescription.PatientID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", prescription.MedicineID).First(&medicine); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", prescription.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	pr := entity.Prescription{
		PrescriptionID: prescription.PrescriptionID,
		Patient:        patient,                // โยงความสัมพันธ์กับ Entity Resolution
		Medicine:       medicine,               // โยงความสัมพันธ์กับ Entity Video
		Employee:       employee,               // โยงความสัมพันธ์กับ Entity Playlist
		Symptom:        prescription.Symptom,   // โยงความสัมพันธ์กับ Entity
		Case_Time:      prescription.Case_Time, // ตั้งค่าฟิลด์ watchedTime
	}

	//บันทึก
	if err := entity.DB().Create(&pr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": pr})

}

// GET /prescriptions/:id
func GetPrescription(c *gin.Context) {
	var prescriptions entity.Prescription
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM prescriptions WHERE id = ?", id).
		Scan(&prescriptions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": prescriptions})
}

// GET /prescriptions
func ListPrescription(c *gin.Context) {
	var prescriptions []entity.Prescription
	if err := entity.DB().Preload("Medicine").Preload("Patient").Preload("Employee").Raw("SELECT * FROM prescriptions").Find(&prescriptions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": prescriptions})

}

// DELETE /prescriptions/:id
func DeletePrescription(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM prescriptions WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prescription not found"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /prescriptions
func UpdatePrescription(c *gin.Context) {
	var prescriptions entity.Prescription
	if err := c.ShouldBindJSON(&prescriptions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", prescriptions.ID).First(&prescriptions); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prescription not found"})

	}
	if err := entity.DB().Save(&prescriptions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": prescriptions})
}

// POST /patients
func CreatePatient(c *gin.Context) {
	var patient entity.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": patient})
}

// GET /patients/:id
func GetPatient(c *gin.Context) {
	var patient entity.Patient
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM patients WHERE id = ?", id).Scan(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": patient})
}

// GET /patients
func ListPatient(c *gin.Context) {
	var patient []entity.Patient
	if err := entity.DB().Raw("SELECT * FROM patients").Scan(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patient})

}

// DELETE /patients/:id
func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM patients WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prescription not found"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /patients
func UpdatePatient(c *gin.Context) {
	var patient entity.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", patient.ID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prescription not found"})

	}
	if err := entity.DB().Save(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patient})
}
