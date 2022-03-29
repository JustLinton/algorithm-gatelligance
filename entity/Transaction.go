package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Transaction struct {
	ID        string `gorm:"primary_key"`
	Server    int
	Owner     string
	CreatedAt time.Time
	Type      string
	Avatar    string
	Title     string
}

func InitTransaction(db *gorm.DB) {
	db.AutoMigrate(&Transaction{})
}

// // Transaction 算法事务信息
// type Transaction struct {
// 	ID   string `gorm:"primary_key"`
// 	Type string
// }

// func InitTransactionEntity(db *gorm.DB) {
// 	db.AutoMigrate(&Transaction{})
// }
