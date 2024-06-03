package repositories

import "gorm.io/gorm"

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(dbConnection *gorm.DB) transactionRepository {
	return transactionRepository{
		db: dbConnection,
	}
}
