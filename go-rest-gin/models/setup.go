package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

// untuk menyimpan koneksi database agar bisa di gunakan di seluruh file project
var DB *gorm.DB

func ConnectDatabase() {
	// koneksi dengan mysql/database
	database, error := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_restapi_gin"))

	// cek jika ada error
	if error != nil {
		panic(error)
	}

	// migrasi/tambah table berdasar pada struct product
	database.AutoMigrate(&Product{})

	// simpan koneksi database ke pointer DB
	DB = database
}