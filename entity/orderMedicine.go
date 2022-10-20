package entity

import (
	"time"

	"gorm.io/gorm"
)

type MedicineCompany struct {
	gorm.Model
	Company_Name string
	Location     string

	MedicineOrders []MedicineOrder `gorm:"foreignKey:MedicineCompanyID"`
}

type MedicineOrder struct {
	gorm.Model
	OrderID     uint
	OrderAmount uint
	OrderTime   time.Time
	EmployeeID  *uint
	Employee    Employee

	MedicineID *uint
	Medicine   Medicine

	MedicineCompanyID *uint
	MedicineCompany   MedicineCompany
}
