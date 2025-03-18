package service

type ExpenseService interface {
	Add(details string, amount float64) (Expense, error)
}

type expenseService struct{}

func NewExpenseService() ExpenseService {
	return &expenseService{}
}

func (e *expenseService) Add(details string, amount float64) (Expense, error) {
	return Expense{}, nil
}