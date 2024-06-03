package ctlwallets

import (
	"github.com/diazharizky/buddi-backend/internal/interfaces"
	"github.com/diazharizky/buddi-backend/internal/repositories"
	"gorm.io/gorm"
)

type Controller struct {
	WalletRepository interfaces.WalletRepository
}

func NewDefault(dbConnection *gorm.DB) (ctl Controller) {
	ctl.WalletRepository = repositories.NewWalletRepository(dbConnection)
	return
}
