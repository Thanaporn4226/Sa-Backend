package controller

import (
	"net/http"

	"github.com/Pet002/Project-sa-65/entity"
	"github.com/gin-gonic/gin"
)

//-------------------------------PayMedicine ---------------------------

// List all PayMedicine
// GET /paymedicines
func ListPayMedicine(c *gin.Context) {
	var paymedicines []entity.PayMedicine
	if err := entity.DB().Preload("MedicineLabel").Preload("Prescription").Preload("Employee").Raw("SELECT * FROM pay_medicines").Find(&paymedicines).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": paymedicines,
	})

}

// GET /paymedicines/:id
func GetPayMedicine(c *gin.Context) {
	var payMedicine entity.PayMedicine
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM pay_medicines WHERE id = ?", id).Scan(&payMedicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": payMedicine,
		},
	)
}

// POST /paymedicines
func CreatePayMedicine(c *gin.Context) {
	//main
	var payMedicine entity.PayMedicine
	//relation
	var login entity.Login
	var employee entity.Employee
	var medicinelabel entity.MedicineLabel
	var prescription entity.Prescription

	//bind data จาก frontend มาไว้ในนี้
	if err := c.ShouldBindJSON(&payMedicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if tx := entity.DB().Where("id = ?", payMedicine.EmployeeID).First(&login); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "login not found"})
		return
	}

	if tx := entity.DB().Where("login_id = ?", login.ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "login not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", payMedicine.MedicineLabelID).First(&medicinelabel); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicinelabel not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", payMedicine.PrescriptionID).First(&prescription); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "perscription not found"})
		return
	}

	//สร้าง payMedicine

	pm := entity.PayMedicine{
		Amount:        payMedicine.Amount,
		Price:         payMedicine.Price,
		PayDate:       payMedicine.PayDate,
		MedicineLabel: medicinelabel,
		Prescription:  prescription,
		Employee:      employee,
	}

	if err := entity.DB().Create(&pm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Create payMedicine success",
		"data":   pm,
	})

}

//PATCH /paymedicines

func UpdatePayMedicine(c *gin.Context) {
	var paymedicine entity.PayMedicine
	if err := c.ShouldBindJSON(&paymedicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if tx := entity.DB().Where("id = ?", paymedicine.ID).First(&paymedicine); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "paymedicine not found"})
		return
	}

	if err := entity.DB().Save(&paymedicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Update Success",
		"data":   paymedicine,
	})
}

// DELETE /paymedicines/:id
func DeletePayMedicine(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM pay_medicines WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payMedicine not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

//------------------------------------------ Pay medicine Communication ---------------
