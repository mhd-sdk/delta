package algorithm

import "fmt"

type Algorithm struct {
	Balance float64
	Logs    []string
}

func NewAlgorithm(initialBalance float64) *Algorithm {
	return &Algorithm{
		Balance: initialBalance,
		Logs:    []string{},
	}
}

func (a *Algorithm) AddBalance(amount float64) {
	a.Balance += amount
	a.AddLog(fmt.Sprintf("Added %.2f to balance", amount))
}

func (a *Algorithm) RemoveBalance(amount float64) error {
	if amount > a.Balance {
		err := fmt.Errorf("insufficient balance to remove %.2f", amount)
		a.AddLog(err.Error())
		return err
	}
	a.Balance -= amount
	a.AddLog(fmt.Sprintf("Removed %.2f from balance", amount))
	return nil
}

func (a *Algorithm) AddLog(message string) {
	a.Logs = append(a.Logs, message)
}

func (a *Algorithm) GetLogs() []string {
	return a.Logs
}
