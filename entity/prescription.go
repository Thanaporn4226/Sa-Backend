package entity

import (
	"time"

	"gorm.io/gorm"
)

// คนไข้
type Patient struct {
	gorm.Model

	PID     string
	Name    string
	Surname string
	Age     uint
	Gender  string
	Allergy string

	Prescriptions []Prescription `gorm:"foreignKey:PatientID"`
}

// Entity หลัก
type Prescription struct {
	gorm.Model
	PrescriptionID string
	Symptom        string //อาการป่วย
	Case_Time      time.Time

	EmployeeID *uint
	Employee   Employee

	MedicineID *uint
	Medicine   Medicine

	PatientID *uint
	Patient   Patient

	//Link to Pay Medicine
	PayMedicines []PayMedicine `gorm:"foreignKey:PrescriptionID"`
}
