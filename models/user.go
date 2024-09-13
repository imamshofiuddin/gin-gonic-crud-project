package models

type User struct {
	Id    int64  `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(200)" json:"name"`
	Email string `gorm:"type:varchar(200)" json:"email"`
	Role  int    `gorm:"type:varchar(200)" json:"role"`
}
