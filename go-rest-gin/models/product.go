package models

type Product struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	NamaProduk string `gorm:"type:varchar(255)" json:"nama_produk"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
}