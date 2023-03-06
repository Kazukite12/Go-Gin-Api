package models

type Product struct {
	Id        int64   `gorm:"primaryKey" json:"id"`
	Produk    string  `gorm:"type:varchar(300)" json:"Produk"`
	Deskripsi string  `gorm:"type:text" json:"deskripsi"`
	Price     float64 `gorm:"type:decimal(10,2)" json:"price"`
	Image     string  `gorm:"type:varchar(300)" json:"Image"`
}
