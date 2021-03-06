package models

import "errors"

// The MoneyStorage describes item for sale
type MoneyStorage struct {
	moneyAmount, userDeposit float32
}

func NewMoneyStorage() *MoneyStorage {
	return &MoneyStorage{0, 0}
}

// AddMoney allows user to deposit money
func (m *MoneyStorage) AddMoney(amount float32) {
	m.userDeposit += amount
}

func (m *MoneyStorage) UseMoney(amount float32) (float32, error) {
	if m.userDeposit-amount < 0 {
		return m.userDeposit, errors.New("Not enough money")
	}

	m.userDeposit -= amount
	m.moneyAmount += amount

	return m.userDeposit, nil
}

func (m *MoneyStorage) MakeCashBack() (cashback float32) {
	cashback = m.userDeposit
	m.userDeposit = 0

	return cashback
}

func (m *MoneyStorage) DepositAmount() float32 {
	return m.userDeposit
}
