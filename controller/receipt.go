package controller

import (
	"net/http"

	"github.com/Thanaporn4226/Project-sa-65/entity"
	"github.com/gin-gonic/gin"
)

//--------------------------------- Payment Types -----------------------------------

func ListPaymentTypes(c *gin.Context) {
	var paymentTypes []entity.PaymentTypes
	if err := entity.DB().Raw("SELECT * FROM payment_types").Scan(&paymentTypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": paymentTypes})

}
func GetPaymentType(c *gin.Context) {
	var paymenttypes entity.PaymentTypes
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM payment_types WHERE id = ?", id).
		Scan(&paymenttypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": paymenttypes})
}
func DeletePaymentType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM prescriptions WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prescription not found"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": id})

}

// ----------------------------------- Receipt -------------------------------------------------
// GET /receipts
func ListReceipts(c *gin.Context) {
	var receipts []entity.Receipt
	if err := entity.DB().Preload("PayMedicine").Preload("Employee").Preload("Types").Raw("SELECT * FROM receipts").Find(&receipts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": receipts})

}

// GET /receipts/:id
func GetReceipts(c *gin.Context) {
	var receipt entity.Receipt
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM receipts WHERE id = ?", id).
		Scan(&receipt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": receipt})
}

// POST /receipts\
func CreateReceipts(c *gin.Context) {
	//main
	var receipt entity.Receipt
	//sub
	var paymentType entity.PaymentTypes
	var payMedicine entity.PayMedicine
	var employee entity.Employee

	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", receipt.TypesID).First(&paymentType); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", receipt.PayMedicineID).First(&payMedicine); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", receipt.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	rp := entity.Receipt{
		TotalPrice:  receipt.TotalPrice,
		Receive:     receipt.Receive,
		Refund:      receipt.Refund,
		Employee:    employee,
		Types:       paymentType,
		PayMedicine: payMedicine,
	}

	//บันทึก
	if err := entity.DB().Create(&rp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": rp})

}

// DELETE /receipts/:id
func DeleteReceipts(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM receipts WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "receipts not found"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": id})

}
