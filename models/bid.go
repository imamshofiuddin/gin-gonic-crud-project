package models

type Bid struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	ItemId   int64  `gorm:"type:int" json:"item_id"`
	UserId   int64  `gorm:"type:int" json:"user_id"`
	BidPrice int    `gorm:"type:int" json:"bid_price"`
	Status   string `gorm:"type:string" json:"status"`
}
