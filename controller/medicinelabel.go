package controller

import (
	"net/http"

	"github.com/Pet002/Project-sa-65/entity"
	"github.com/gin-gonic/gin"
)

// ----------------------------------------- using -----------------------------------

// POST /MedicineUse
func CreateMedicineUse(c *gin.Context) {
	var MedicineUse entity.MedicineUse
	if err := c.ShouldBindJSON(&MedicineUse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&MedicineUse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": MedicineUse})
}

// GET /MedicineUse/:id
func GetMedicineUse(c *gin.Context) {
	var medicineUse entity.MedicineUse
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM medicine_uses WHERE id = ?", id).Scan(&medicineUse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if tx := entity.DB().Where("id = ?", id).First(&medicineLabel); tx.RowsAffected == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "medicine_use not found"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"data": medicineUse})
}

// GET /MedicineUse
func ListMedicineUse(c *gin.Context) {
	var MedicineUse []entity.MedicineUse
	if err := entity.DB().Raw("SELECT * FROM medicine_uses").Scan(&MedicineUse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": MedicineUse})
}

// DELETE /MedicineUse/:id
func DeleteMedicineUse(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM medicine_use WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine_use not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /MedicineUse
func UpdateMedicineUse(c *gin.Context) {
	var MedicineUse entity.MedicineUse
	if err := c.ShouldBindJSON(&MedicineUse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", MedicineUse.ID).First(&MedicineUse); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine_use not found"})
		return
	}

	if err := entity.DB().Save(&MedicineUse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": MedicineUse})
}

// ----------------------------------------- warning -----------------------------------

// POST /warning
func CreateWarning(c *gin.Context) {
	var Warning entity.Warning
	if err := c.ShouldBindJSON(&Warning); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Warning).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": Warning})
}

// GET /warning/:id
func GetWarning(c *gin.Context) {
	var warning entity.Warning
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM warnings WHERE id = ?", id).Scan(&warning).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if tx := entity.DB().Where("id = ?", id).First(&warning); tx.RowsAffected == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "warning not found"})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{"data": warning})
}

// GET /Warning
func ListWarning(c *gin.Context) {
	var warning []entity.Warning
	if err := entity.DB().Raw("SELECT * FROM warnings").Scan(&warning).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if err := entity.DB().Raw("SELECT * FROM warnings").Scan(&warning).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"data": warning})
}

// DELETE /warning/:id
func DeleteWarning(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM warning WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "warning not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Warning
func UpdateWarning(c *gin.Context) {
	var Warning entity.Warning
	if err := c.ShouldBindJSON(&Warning); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Warning.ID).First(&Warning); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "warning not found"})
		return
	}

	if err := entity.DB().Save(&Warning).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Warning})
}

//--------------------------------------------- Medicine Label---------------------------------------------------------------------

// POST /MedicineLabel
func CreateMedicineLabel(c *gin.Context) {

	var medicinelabel entity.MedicineLabel
	var medicineuse entity.MedicineUse
	var warning entity.Warning
	var employee entity.Employee

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร watchVideo
	if err := c.ShouldBindJSON(&medicinelabel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา medicineuse ด้วย id
	if tx := entity.DB().Where("id = ?", medicinelabel.MedicineUseID).First(&medicineuse); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine_use not found"})
		return
	}

	// 10: ค้นหา warning ด้วย id
	if tx := entity.DB().Where("id = ?", medicinelabel.WarningID).First(&warning); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "warning not found"})
		return
	}

	// 11: ค้นหา employee ด้วย id
	if tx := entity.DB().Where("id = ?", medicinelabel.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	// 12: สร้าง MedicineLabel
	wv := entity.MedicineLabel{
		MedicineUse:   medicineuse,                 // โยงความสัมพันธ์กับ Entity medicineuse
		Warning:       warning,                     // โยงความสัมพันธ์กับ Entity warning
		Employee:      employee,                    // โยงความสัมพันธ์กับ Entity employee
		RecordingDate: medicinelabel.RecordingDate, // ตั้งค่าฟิลด์ RecordingDate
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": wv})
}

// GET /MedicineLabel/:id
func GetMedicineLabel(c *gin.Context) {
	var medicineLabel entity.MedicineLabel
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM medicine_labels WHERE id = ?", id).Preload("MedicineUse").Preload("Warning").Find(&medicineLabel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicineLabel})
}

// GET /MedicineLabel
func ListMedicineLabel(c *gin.Context) {
	var medicineLabel []entity.MedicineLabel
	if err := entity.DB().Raw("SELECT * FROM medicine_labels").Preload("Employee").Preload("MedicineUse").Preload("Warning").Find(&medicineLabel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if err := entity.DB().Preload("MedicineUse").Preload("Warning").Preload("Employee").Raw("SELECT * FROM medicine_labels").Find(&MedicineLabel).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"data": medicineLabel})
}

// DELETE /MedicineLabel/:id
func DeleteMedicineLabel(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM medicine_labels WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine_label not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /MedicineLabel
func UpdateMedicineLabel(c *gin.Context) {
	var MedicineLabel entity.MedicineLabel
	if err := c.ShouldBindJSON(&MedicineLabel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", MedicineLabel.ID).First(&MedicineLabel); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine_label not found"})
		return
	}

	if err := entity.DB().Save(&MedicineLabel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": MedicineLabel})
}
