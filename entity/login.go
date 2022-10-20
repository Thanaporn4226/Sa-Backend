package entity

import (
	"gorm.io/gorm"
)

// when Login We will use this table
type Login struct {
	gorm.Model
	User     string `gorm:"uniqueIndex"`
	Password string

	//link foreign Key to Employee table
	Employee []Employee `gorm:"foreignKey:LoginID"`
}

type Role struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`

	//link foreign Key to Employee table
	Employees []Employee `gorm:"foreignKey:RoleID"`
}

// main Table to link to another Feature
type Employee struct {
	gorm.Model
	Name    string
	Surname string

	//recive LoginID from Login Table
	LoginID *uint `gorm:"uniqueIndex"` //set Unique for 1 to 1 relational database
	//To easier for join table
	Login Login

	//recive RoleID from Role Table
	RoleID *uint
	//To easier for join table
	Role Role

	//For Enter Relation
	PayMedicines  []PayMedicine   `gorm:"foreignKey:EmployeeID"`
	Medicine      []Medicine      `gorm:"foreignKey:EmployeeID"`
	Prescription  []Prescription  `gorm:"foreignKey:EmployeeID"`
	Receipt       []Receipt       `gorm:"foreignKey:EmployeeID"`
	MedicineLabel []MedicineLabel `gorm:"foreignKey:EmployeeID"`
}
