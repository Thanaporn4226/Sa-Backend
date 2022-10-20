package entity

import (
	"time"

	"gorm.io/gorm"
)

type PayMedicine struct {
	gorm.Model

	Amount  uint
	Price   float64
	PayDate time.Time `valid:"past"`

	MedicineLabelID *uint
	MedicineLabel   MedicineLabel

	PrescriptionID *uint `gorm:"uniqueIndex"` //set Unique for 1 to 1 relational database
	Prescription   Prescription

	EmployeeID *uint
	Employee   Employee

	Receipts []Receipt `gorm:"foreignKey:PayMedicineID"`
}
