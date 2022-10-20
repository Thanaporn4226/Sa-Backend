package controller

import (
	"net/http"

	"github.com/Pet002/Project-sa-65/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

//---------------------------------type-----------------------------------

// POST /type
func CreateType(c *gin.Context) {
	var MedicineType entity.MedicineType
	if err := c.ShouldBindJSON(&MedicineType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&MedicineType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": MedicineType})
}

// GET /type/:id
func GetType(c *gin.Context) {
	var MedicineType entity.MedicineType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM medicine_types WHERE id = ?", id).Scan(&MedicineType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": MedicineType})
}

// GET /type
func ListType(c *gin.Context) {
	var MedicineType []entity.MedicineType
	if err := entity.DB().Raw("SELECT * FROM medicine_types").Scan(&MedicineType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": MedicineType})
}

// DELETE /type/:id
func DeleteType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM medicine_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MedicineType not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /type
func UpdateType(c *gin.Context) {
	var MedicineType entity.MedicineType
	if err := c.ShouldBindJSON(&MedicineType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", MedicineType.ID).First(&MedicineType); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type not found"})
		return
	}

	if err := entity.DB().Save(&MedicineType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": MedicineType})
}

//-----------------------------------storage--------------------------------

// POST /storage
func CreateStorage(c *gin.Context) {
	var storage entity.Storage
	if err := c.ShouldBindJSON(&storage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&storage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": storage})
}

// GET /storage/:id
func GetStorage(c *gin.Context) {
	var storage entity.Storage
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM storages WHERE id = ?", id).Scan(&storage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": storage})
}

// GET /storage
func ListStorage(c *gin.Context) {
	var storage []entity.Storage
	if err := entity.DB().Raw("SELECT * FROM storages").Scan(&storage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": storage})
}

// DELETE /storage/:id
func DeleteStorage(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM storages WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "storages not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /storage
func UpdateStorage(c *gin.Context) {
	var storage entity.Storage
	if err := c.ShouldBindJSON(&storage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", storage.ID).First(&storage); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "storages not found"})
		return
	}

	if err := entity.DB().Save(&storage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": storage})
}

// POST /Medicine
func CreateMedicine(c *gin.Context) {

	var medicine entity.Medicine
	var employee entity.Employee
	var storage entity.Storage
	var medicine_type entity.MedicineType

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร medicine
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา employee ด้วย id
	if tx := entity.DB().Where("id = ?", medicine.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 10: ค้นหา type ด้วย id
	if tx := entity.DB().Where("id = ?", medicine.TypeID).First(&medicine_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type not found"})
		return
	}

	// 11: ค้นหา storage ด้วย id
	if tx := entity.DB().Where("id = ?", medicine.StorageID).First(&storage); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "storage not found"})
		return
	}
	// 12: สร้าง medicine
	m := entity.Medicine{
		Employee: employee,        // โยงความสัมพันธ์กับ Entity Employee
		Storage:  storage,         // โยงความสัมพันธ์กับ Entity Storage
		Type:     medicine_type,   // โยงความสัมพันธ์กับ Entity Type
		Name:     medicine.Name,   // ตั้งค่าฟิลด์ Name
		MFD:      medicine.MFD,    // ตั้งค่าฟิลด์ MFD
		EXP:      medicine.EXP,    // ตั้งค่าฟิลด์ EXP
		Amount:   medicine.Amount, // ตั้งค่าฟิลด์ Amount
	}

	// ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&m).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": m})
}

// GET /medicine/:id
func GetMedicine(c *gin.Context) {
	var medicine entity.Medicine
	id := c.Param("id")
	if err := entity.DB().Preload("Employee").Preload("Storage").Preload("Type").Raw("SELECT * FROM medicines WHERE id = ?", id).Find(&medicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": medicine})
}

// GET /medicine
func ListMedicine(c *gin.Context) {
	var medicine []entity.Medicine
	if err := entity.DB().Preload("Employee").Preload("Storage").Preload("Type").Raw("SELECT * FROM medicines").Find(&medicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicine})
}

// DELETE /medicine/:id
func DeleteMedicine(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM medicines WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /medicine
func UpdateMedicine(c *gin.Context) {
	var medicine entity.Medicine
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", medicine.ID).First(&medicine); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine not found"})
		return
	}

	if err := entity.DB().Save(&medicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicine})
}
