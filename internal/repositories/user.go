package repositories

import "gorm.io/gorm"

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(dbConnection *gorm.DB) userRepository {
	return userRepository{
		db: dbConnection,
	}
}
