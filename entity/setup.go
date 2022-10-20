package entity

import (
	"time"

	"github.com/Pet002/Project-sa-65/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("MedicineRoom.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		//Medicine label
		&MedicineUse{},
		&Warning{},
		&MedicineLabel{},
		//Employee system
		&Login{},
		&Role{},
		&Employee{},
		//Prescription
		&Prescription{},
		&Patient{},
		//Medicine
		&MedicineType{},
		&Storage{},
		&Medicine{},
		//PayMedicine
		&PayMedicine{},
		//Receipt
		&PaymentTypes{},
		&Receipt{},
		//OrderMedicine
		&MedicineCompany{},
		&MedicineOrder{},
	)

	db = database

	//เราจะทำการสร้าง Admin account ไว้สำหรับการสร้าง Employee และ role ต่างๆ

	password, _ := services.Hash("123456")

	admin := Role{
		Name: "admin",
	}
	db.Model(&Role{}).Create(&admin)

	intendant := Role{
		Name: "intendant",
	}

	db.Model(&Role{}).Create(&intendant)

	payment := Role{
		Name: "payment",
	}
	db.Model(&Role{}).Create(&payment)

	pharmacist := Role{
		Name: "pharmacist",
	}
	db.Model(&Role{}).Create(&pharmacist)

	//Admin 1
	admin1 := Login{
		User:     "Admin1",
		Password: string(password),
	}
	admin1err := db.Model(&Login{}).Create(&admin1)
	admin1emp := Employee{
		Name:    "Admin1",
		Surname: "Example1",
		Login:   admin1,
		Role:    admin,
	}
	if admin1err.Error == nil {
		db.Model(&Employee{}).Create(&admin1emp)
	}

	//Admin 2
	admin2 := Login{
		User:     "Admin2",
		Password: string(password),
	}
	admin2err := db.Model(&Login{}).Create(&admin2)
	admin2emp := Employee{
		Name:    "Admin1",
		Surname: "Example1",
		Login:   admin2,
		Role:    admin,
	}
	if admin2err.Error == nil {
		db.Model(&Employee{}).Create(&admin2emp)
	}

	//Intendant 1
	intendant1 := Login{
		User:     "Intendant1",
		Password: string(password),
	}
	inten1err := db.Model(&Login{}).Create(&intendant1)
	inten1emp := Employee{
		Name:    "Intendant1",
		Surname: "Example1",
		Login:   intendant1,
		Role:    intendant,
	}
	if inten1err.Error == nil {
		db.Model(&Employee{}).Create(&inten1emp)
	}

	//Intendant 2
	intendant2 := Login{
		User:     "Intendant2",
		Password: string(password),
	}
	inten2err := db.Model(&Login{}).Create(&intendant2)
	inten2emp := Employee{
		Name:    "Intendant2",
		Surname: "Example1",
		Login:   intendant2,
		Role:    intendant,
	}
	if inten2err.Error == nil {
		db.Model(&Employee{}).Create(&inten2emp)
	}

	//Payment 1
	payment1 := Login{
		User:     "Payment1",
		Password: string(password),
	}
	pay1err := db.Model(&Login{}).Create(&payment1)
	pay1emp := Employee{
		Name:    "Payment1",
		Surname: "Example1",
		Login:   payment1,
		Role:    payment,
	}
	if pay1err.Error == nil {
		db.Model(&Employee{}).Create(&pay1emp)
	}

	//Payment 2
	payment2 := Login{
		User:     "Payment2",
		Password: string(password),
	}
	pay2err := db.Model(&Login{}).Create(&payment2)
	pay2emp := Employee{
		Name:    "Payment2",
		Surname: "Example2",
		Login:   payment2,
		Role:    payment,
	}
	if pay2err.Error == nil {
		db.Model(&Employee{}).Create(&pay2emp)
	}

	//pharmacist 1
	pharmacist1 := Login{
		User:     "Pharmacist1",
		Password: string(password),
	}
	phar1err := db.Model(&Login{}).Create(&pharmacist1)
	phar1emp := Employee{
		Name:    "Pharmacist1",
		Surname: "Example1",
		Login:   pharmacist1,
		Role:    pharmacist,
	}
	if phar1err.Error == nil {
		db.Model(&Employee{}).Create(&phar1emp)
	}

	//pharmacist 2
	pharmacist2 := Login{
		User:     "Pharmacist2",
		Password: string(password),
	}
	phar2err := db.Model(&Login{}).Create(&pharmacist2)
	phar2emp := Employee{
		Name:    "Pharmacist2",
		Surname: "Example2",
		Login:   pharmacist2,
		Role:    pharmacist,
	}
	if phar2err.Error == nil {
		db.Model(&Employee{}).Create(&phar2emp)
	}

	// //-------------------------------------- medicine ----------------------------------------------

	//----------Type-------------------
	tha := MedicineType{
		Tmedicine:  "ยาใช้ภายนอก",
		Utilzation: "ทา",
	}
	db.Model(&MedicineType{}).Create(&tha)

	med := MedicineType{
		Tmedicine:  "ยาใช้ภายใน",
		Utilzation: "เม็ด",
	}
	db.Model(&MedicineType{}).Create(&med)

	cheed := MedicineType{
		Tmedicine:  "ยาใช้ภายใน",
		Utilzation: "ฉีด",
	}
	db.Model(&MedicineType{}).Create(&cheed)

	nam := MedicineType{
		Tmedicine:  "ยาใช้ภายใน",
		Utilzation: "น้ำ",
	}
	db.Model(&MedicineType{}).Create(&nam)

	//-------------Storage----------------
	b1 := Storage{
		Name: "B1",
	}
	db.Model(&Storage{}).Create(&b1)

	b2 := Storage{
		Name: "B2",
	}
	db.Model(&Storage{}).Create(&b2)

	// //-------------Medicine---------------

	med1 := Medicine{
		Employee: inten1emp,
		Name:     "Paracetamol",
		Type:     med,
		MFD:      time.Date(2022, 8, 28, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 8, 28, 12, 0, 0, 0, time.UTC),
		Storage:  b1,
		Amount:   100,
	}
	db.Model(&Medicine{}).Create(&med1)

	//=====
	med2 := Medicine{
		Employee: inten1emp,
		Name:     "Menopain",
		Type:     med,
		MFD:      time.Date(2022, 8, 30, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 8, 30, 12, 0, 0, 0, time.UTC),
		Storage:  b1,
		Amount:   100,
	}
	db.Model(&Medicine{}).Create(&med2)

	//====
	med3 := Medicine{
		Employee: inten2emp,
		Name:     "b-derm",
		Type:     tha,
		MFD:      time.Date(2022, 8, 30, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 8, 30, 12, 0, 0, 0, time.UTC),
		Storage:  b2,
		Amount:   100,
	}
	db.Model(&Medicine{}).Create(&med3)

	//===
	med4 := Medicine{
		Employee: inten1emp,
		Name:     "Cetirizine",
		Type:     med,
		MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
		Storage:  b2,
		Amount:   100,
	}
	db.Model(&Medicine{}).Create(&med4)

	//===
	med5 := Medicine{
		Employee: inten1emp,
		Name:     "ฺBromhexine", //ยาละลายเสมหะ
		Type:     med,
		MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
		Storage:  b2,
		Amount:   100,
	}
	db.Model(&Medicine{}).Create(&med5)

	//===
	med6 := Medicine{
		Employee: inten2emp,
		Name:     "Cenor", //ยาต้านเชื้อแบคทีเรีย
		Type:     med,
		MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
		Storage:  b1,
		Amount:   100,
	}
	db.Model(&Medicine{}).Create(&med6)

	//===
	med7 := Medicine{
		Employee: inten1emp,
		Name:     "Tramadol", //ยาบรรเทาอาการปวดรุนแรง
		Type:     med,
		MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
		Storage:  b1,
		Amount:   100,
	}
	db.Model(&Medicine{}).Create(&med7)

	//====
	med8 := Medicine{
		Employee: inten2emp,
		Name:     "Salol et Menthol Mixture", //ยาธาตุน้ำขาว
		Type:     med,
		MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
		Storage:  b2,
		Amount:   100,
	}
	db.Model(&Medicine{}).Create(&med8)

	//====
	med9 := Medicine{
		Employee: inten1emp,
		Name:     "Atorvastatin", //ยา]f++
		Type:     med,
		MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
		Storage:  b2,
		Amount:   100,
	}
	db.Model(&Medicine{}).Create(&med9)

	// //--------------------------------------------------- Medicine Label part -------------------------------------

	//MedicineUse Data
	mu1 := MedicineUse{
		How_To_Use: "ครั้งละ 1 เม็ด ทุก 4-6 ชั่วโง เวลาปวดหรือมีไข้",
	}
	db.Model(&MedicineUse{}).Create(&mu1)

	mu2 := MedicineUse{
		How_To_Use: "ครั้งละ 1 เม็ด หลังอาหาร เช้า-กลางวัน-เย็น",
	}
	db.Model(&MedicineUse{}).Create(&mu2)

	mu3 := MedicineUse{
		How_To_Use: "ครั้งละ 1 เม็ด ก่อนอาหาร เช้า-กลางวัน-เย็น",
	}
	db.Model(&MedicineUse{}).Create(&mu3)

	// Warning Data
	w1 := Warning{
		Medicine_Warning: "ห้ามใช้เกิน 8 เม็ดต่อวัน",
	}
	db.Model(&Warning{}).Create(&w1)

	w2 := Warning{
		Medicine_Warning: "ทานยาแล้วอาจรู้สึกง่วงซึม",
	}
	db.Model(&Warning{}).Create(&w2)

	w3 := Warning{
		Medicine_Warning: "ไม่ควรใช้เกินกว่าขนาดที่ระบุ",
	}
	db.Model(&Warning{}).Create(&w3)

	//MedicineLabel Data
	ml1 := MedicineLabel{
		RecordingDate: time.Now(),
		Warning:       w1,
		MedicineUse:   mu1,
		Employee:      phar1emp,
	}
	db.Model(&MedicineLabel{}).Create(&ml1)

	ml2 := MedicineLabel{
		RecordingDate: time.Now(),
		Warning:       w2,
		MedicineUse:   mu2,
		Employee:      phar1emp,
	}
	db.Model(&MedicineLabel{}).Create(&ml2)

	ml3 := MedicineLabel{
		RecordingDate: time.Now(),
		Warning:       w3,
		MedicineUse:   mu3,
		Employee:      phar2emp,
	}
	db.Model(&MedicineLabel{}).Create(&ml3)

	//------------------------------------------ Prescription ---------------------------

	//patient1
	patient_1 := Patient{
		PID:     "P0001",
		Name:    "กิตติมากร",
		Surname: "สอนแก้ว",
		Age:     21,
		Gender:  "หญิง",
		Allergy: "ยาแอสไพริน",
	}
	db.Model(&Patient{}).Create(&patient_1)

	//patient2
	patient_2 := Patient{
		PID:     "P0002",
		Name:    "ยศพล",
		Surname: "จันทะนาม",
		Age:     22,
		Gender:  "ชาย",
		Allergy: "ไม่แพ้ยาใดๆ",
	}
	db.Model(&Patient{}).Create(&patient_2)

	//patient3
	patient_3 := Patient{
		PID:     "P0003",
		Name:    "กฤษฎา",
		Surname: "น้อยผา",
		Age:     22,
		Gender:  "ชาย",
		Allergy: "ไม่แพ้ยาใดๆ",
	}
	db.Model(&Patient{}).Create(&patient_3)

	prescription_1 := Prescription{
		PrescriptionID: "P00001",
		Symptom:        "ไข้หวัด",
		Case_Time:      time.Now(),
		Employee:       phar1emp,
		Medicine:       med5,
		Patient:        patient_1,
	}
	db.Model(&Prescription{}).Create(&prescription_1)

	prescription_2 := Prescription{
		PrescriptionID: "P00002",
		Symptom:        "ปวดหัว",
		Case_Time:      time.Now(),
		Employee:       phar2emp,
		Medicine:       med1,
		Patient:        patient_2,
	}
	db.Model(&Prescription{}).Create(&prescription_2)

	prescription_3 := Prescription{
		PrescriptionID: "P00003",
		Symptom:        "แผลถลอก",
		Case_Time:      time.Now(),
		Employee:       phar1emp,
		Medicine:       med6,
		Patient:        patient_2,
	}
	db.Model(&Prescription{}).Create(&prescription_3)

	//------------------------ PayMedicine ----------------------------

	payMedicine1 := PayMedicine{
		Amount:        20,
		Price:         500,
		PayDate:       time.Now(),
		MedicineLabel: ml1,
		Prescription:  prescription_1,
		Employee:      phar1emp,
	}
	db.Model(&PayMedicine{}).Create(&payMedicine1)

	payMedicine2 := PayMedicine{
		Amount:        20,
		Price:         500,
		PayDate:       time.Now(),
		MedicineLabel: ml2,
		Prescription:  prescription_2,
		Employee:      phar2emp,
	}
	db.Model(&PayMedicine{}).Create(&payMedicine2)

	// --------------------------------- Receipt ----------------------
	// Payment type
	tp1 := PaymentTypes{
		TypeName: "ชำระเงินสด",
	}
	db.Model(&PaymentTypes{}).Create(&tp1)

	tp2 := PaymentTypes{
		TypeName: "ชำระด้วยการโอน",
	}
	db.Model(&PaymentTypes{}).Create(&tp2)

	//Receipt
	receipt1 := Receipt{
		Employee:    pay1emp,
		Types:       tp1,
		PayMedicine: payMedicine1,
		Receive:     20,
		Refund:      0,
		TotalPrice:  500,
	}
	db.Model(&Receipt{}).Create(&receipt1)

}
