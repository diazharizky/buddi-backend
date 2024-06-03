package ctlusers

import (
	"github.com/diazharizky/buddi-backend/internal/interfaces"
	"github.com/diazharizky/buddi-backend/internal/repositories"
	"gorm.io/gorm"
)

type Controller struct {
	UserRepository interfaces.UserRepository
}

func NewDefault(dbConnection *gorm.DB) (ctl Controller) {
	ctl.UserRepository = repositories.NewUserRepository(dbConnection)
	return
}
