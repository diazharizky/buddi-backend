package services

type createTransactionService struct{}

func NewCreateTransactionService() createTransactionService {
	return createTransactionService{}
}

func (svc createTransactionService) Call() error {
	return nil
}
