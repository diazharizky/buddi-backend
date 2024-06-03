package repositories

import "gorm.io/gorm"

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(dbConnection *gorm.DB) walletRepository {
	return walletRepository{
		db: dbConnection,
	}
}
