package entity

import (
	"time"

	"gorm.io/gorm"
)

type MedicineUse struct {
	gorm.Model
	How_To_Use    string
	MedicineLabel []MedicineLabel `gorm:"foreignKey:MedicineUseID"`
}

type Warning struct {
	gorm.Model
	Medicine_Warning string
	MedicineLabel    []MedicineLabel `gorm:"foreignKey:WarningID"`
}

// Main Entity
type MedicineLabel struct {
	gorm.Model

	RecordingDate time.Time

	// EmployeeID ทำหน้าที่เป็น FK
	EmployeeID *uint
	Employee   Employee

	// WarningID ทำหน้าที่เป็น FK
	WarningID *uint
	Warning   Warning

	// MedicineUseID ทำหน้าที่เป็น FK
	MedicineUseID *uint
	MedicineUse   MedicineUse

	//Link to
	PayMedicine []PayMedicine `gorm:"foreignKey:MedicineLabelID"`
}
