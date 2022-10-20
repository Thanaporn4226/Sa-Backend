package entity

import (
	"gorm.io/gorm"
)

type PaymentTypes struct {
	gorm.Model

	TypeName string
	Receip   []Receipt `gorm:"foreignKey:TypesID"`
}

type Receipt struct {
	gorm.Model
	TotalPrice int
	Receive    int
	Refund     int

	EmployeeID *uint
	Employee   Employee

	TypesID *uint
	Types   PaymentTypes

	PayMedicineID *uint
	PayMedicine   PayMedicine
}
