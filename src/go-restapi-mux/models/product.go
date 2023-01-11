package models

type Product struct {
	Id    int64   `gorm:"primaryKey" json:"id"`
	Name  string  `gorm:"type:varchar(300)" json:"name"`
	Stock int32   `gorm:"type:int(5)" json:"stock"`
	Price float64 `gorm:"type:decimal(14,2)" json:"price"`
}
