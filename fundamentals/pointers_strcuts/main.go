package main

import (
	"errors"
	"fmt"
)

const (
	CASH_IN  = "CASH_IN"
	CASH_OUT = "CASH_OUT"
	TRANSFER = "TRANSFER"
)

type Movement interface {
	cashIn(amount float32)
	cashOut(amount float32) (bool, error)
	withdrawal(account *Account, amount float32) (bool, error)
	statement()
}

type Customer struct {
	name        string
	phoneNumber string
}

type Transaction struct {
	sender       string
	movementType string
	amount       float32
}

type Account struct {
	owner        Customer
	balance      float32
	transactions []Transaction
}

func createTransaction(
	sender *Account,
	amount float32,
	movementType string) Transaction {

	transaction := Transaction{
		sender:       sender.owner.name,
		movementType: movementType,
		amount:       amount,
	}

	sender.transactions = append(sender.transactions, transaction)

	return transaction
}

func (account *Account) cashIn(amount float32) {
	createTransaction(account, amount, CASH_IN)
	account.balance += amount
}

func (account *Account) cashOut(amount float32) (bool, error) {

	if account.balance > amount {
		account.balance -= account.balance
		createTransaction(account, amount, CASH_OUT)
		return true, nil
	}

	return false, errors.New("Insuficient Funds")
}

func (sender *Account) withdrawal(recipient *Account, amount float32) (bool, error) {

	if sender.balance > amount {

		sender.balance -= amount
		createTransaction(sender, amount, CASH_OUT)
		recipient.balance += amount
		createTransaction(recipient, amount, CASH_IN)
		return true, nil
	}

	return false, errors.New("Insuficient Funds")
}

func (account Account) statement() {

	fmt.Printf("Statement for %v\n", account.owner.name)
	for _, transaction := range account.transactions {

		fmt.Printf("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=\n")
		fmt.Printf("Amount: %v\n", transaction.amount)
		fmt.Printf("Movement type: %v\n", transaction.movementType)
		fmt.Printf("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=\n")
	}
}

func main() {

	account1 := Account{
		owner: Customer{
			name:        "Jhon Doe",
			phoneNumber: "5515996578690",
		},
		balance:      0,
		transactions: []Transaction{},
	}

	account2 := Account{
		owner: Customer{
			name:        "Sara Doe",
			phoneNumber: "5511997685543",
		},
		balance:      0,
		transactions: []Transaction{},
	}

	account1.cashIn(100)
	account1.cashIn(100)
	account2.cashIn(300)
	account2.cashIn(400)
	account1.withdrawal(&account2, 50)

	account1.statement()
	account2.statement()

}
