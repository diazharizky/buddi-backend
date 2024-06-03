package services

type createWalletService struct{}

func NewCreateWalletService() createWalletService {
	return createWalletService{}
}

func (svc createWalletService) Call() error {
	return nil
}
