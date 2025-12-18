package Bank

import (
	"errors"
	"fmt"
	"time"
)

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
	Transactions  []Transaction
}

func NewBankAccount(accountNumber, holderName string, initialBalance float64) *BankAccount {
	account := &BankAccount{
		AccountNumber: accountNumber,
		HolderName:    holderName,
		Balance:       initialBalance,
		Transactions:  []Transaction{},
	}
	if initialBalance > 0 {
		account.Transactions = append(account.Transactions, Transaction{
			Type:      "Initial Deposit",
			Amount:    initialBalance,
			Balance:   initialBalance,
			Timestamp: time.Now(),
		})
	}
	return account
}

func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}

	b.Balance += amount
	b.Transactions = append(b.Transactions, Transaction{
		Type:      "Deposit",
		Amount:    amount,
		Balance:   b.Balance,
		Timestamp: time.Now(),
	})

	return nil
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}

	if amount > b.Balance {
		return errors.New("insufficient funds")
	}

	b.Balance -= amount
	b.Transactions = append(b.Transactions, Transaction{
		Type:      "Withdrawal",
		Amount:    amount,
		Balance:   b.Balance,
		Timestamp: time.Now(),
	})

	return nil
}

func (b *BankAccount) GetBalance() float64 {
	return b.Balance
}

func (b *BankAccount) GetTransactionHistory() []Transaction {
	return b.Transactions
}

func (b *BankAccount) PrintTransactionHistory() {
	if len(b.Transactions) == 0 {
		fmt.Println("No transactions found.")
		return
	}

	for i, t := range b.Transactions {
		fmt.Printf("%d. %s - Amount: %.2f, Balance: %.2f, Time: %s\n",
			i+1, t.Type, t.Amount, t.Balance, t.Timestamp.Format("2006-01-02 15:04:05"))
	}
}
