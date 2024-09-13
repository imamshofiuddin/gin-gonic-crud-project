package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	ProductName string `gorm:"type:varchar(100)" json:"product_name"`
	Description string `gorm:"type:text" json:"description"`
	BasePrice   int    `gorm:"type:int" json:"base_price"`
	UserId      int64  `gorm:"type:int" json:"user_id"`
}
