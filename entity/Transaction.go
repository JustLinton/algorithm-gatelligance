package entity

import "github.com/jinzhu/gorm"

// Transaction 算法事务信息
type Transaction struct {
	ID   string `gorm:"primary_key"`
	Type string
}

func InitTransactionEntity(db *gorm.DB) {
	db.AutoMigrate(&Transaction{})
}
