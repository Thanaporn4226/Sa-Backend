package controller

import (
	"net/http"

	"github.com/Pet002/Project-sa-65/entity"
	"github.com/Pet002/Project-sa-65/services"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

//------------------- Part Role ----------------------------------

// Get /roles
// Get All Role
func ListRoles(c *gin.Context) {
	var roles []entity.Role
	if err := entity.DB().Raw("SELECT * FROM roles").Scan(&roles).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": roles,
	})
}

// Get /role/:id
// Get role by id
func GetRole(c *gin.Context) {
	var role entity.Role
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM roles WHERE id = ?", id).Scan(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": role})
}

// POST /roles
func CreateRole(c *gin.Context) {
	var role entity.Role
	// ShouldBindJSON will Read http body in json and change to struct obj.
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Create role success",
		"data":   role,
	})
}

// PATCH /roles
func UpdateRole(c *gin.Context) {
	var role entity.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//มี id นี้อยู่ใน Database ไหม
	if tx := entity.DB().Where("id = ?", role.ID).First(&role); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := entity.DB().Save(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Update Success",
		"data":   role,
	})
}

// DELETE /roles/:id
func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM roles WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roles not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

//------------------- Part Login ----------------------------------

// List all logins
func ListLogin(c *gin.Context) {
	var logins []entity.Login
	if err := entity.DB().Raw("SELECT * FROM logins").Scan(&logins).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   logins,
	})
}

// Get /login/:id
// Get login
func GetLogin(c *gin.Context) {
	var login entity.Login
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM logins WHERE id = ?", id).Scan(&login).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   login,
	})
}

// PATCH /logins
func UpdateLogin(c *gin.Context) {
	var login entity.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if tx := entity.DB().Where("id = ?", login.ID).First(&login); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := entity.DB().Save(&login).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Update Success",
		"data":   login,
	})
}

// DELETE /logins/:id
func DeleteLogin(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM logins WHETE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

//-------------------------------------- Part Employee -----------------------------

// ไว้สำหรับรับค่าตอน Create Employee ตอน POST
type signup struct {
	//Employee
	Name    string
	Surname string
	//Login
	User     string
	Password string
	//Role
	RoleName string
}

// List All Employee
// GET /employees
func ListEmployee(c *gin.Context) {
	var employees []entity.Employee
	if err := entity.DB().Raw("SELECT * FROM employees").Scan(&employees).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": employees,
	})
}

// GET /employees/:id
func GetEmployee(c *gin.Context) {
	var employee entity.Employee
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM employees WHERE id = ?", id).Scan(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// POST /employee
func CreateEmployee(c *gin.Context) {
	var login entity.Login
	var employee entity.Employee
	var role entity.Role
	//ไว้สำหรับรับค่า
	var signup signup

	//รับค่าจาก Body มาก่อน
	if err := c.ShouldBindJSON(&signup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//ค้นหา role จากใน database
	if tx := entity.DB().Where("name = ?", signup.RoleName).First(&role); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role not found"})
		return
	}

	//ค้นหา login จากใน database
	if tx := entity.DB().Where("user = ?", signup.User).First(&login); tx.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "have user in database"})
		return
	}

	bytes, err := services.Hash(signup.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}
	login = entity.Login{
		User:     signup.User,
		Password: string(bytes),
	}

	//ทำการตรวจสอบความถูกต้องของ struct ก่อนนำไปสร้าง record
	if _, err := govalidator.ValidateStruct(login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//สร้างตาราง Login
	if err := entity.DB().Create(&login).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee = entity.Employee{
		Name:    signup.Name,
		Surname: signup.Surname,
		Login:   login,
		Role:    role,
	}
	//ทำการตรวจสอบความถูกต้องของ struct ก่อนนำไปสร้าง record
	if _, err := govalidator.ValidateStruct(employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// บันทึกค่าลงในตารางหลัก
	if err := entity.DB().Create(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "create login Success",
		"data":   employee,
	})
}

// PATCH /employees
func UpdateEmployee(c *gin.Context) {
	var employee entity.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if tx := entity.DB().Where("id = ?", employee.ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := entity.DB().Save(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Update Success",
		"data":   employee,
	})
}

// DELETE /emoloyees/:id
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM employees WHETE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

func GetEmployeeByLoginID(c *gin.Context) {
	var employee entity.Employee
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM employees WHERE login_id = ?", id).Scan(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}
