package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID               int
	Name             string
	Balance          float64
	TransactionHistory []string
}

var accounts []Account

const (
	OptionDeposit       = 1
	OptionWithdraw      = 2
	OptionViewBalance   = 3
	OptionTransactionHistory = 4
	OptionExit          = 5
)

func FindAccount(id int) (*Account, error) {
	for i, acc := range accounts {
		if acc.ID == id {
			return &accounts[i], nil
		}
	}
	return nil, errors.New("account not found")
}

func Deposit(accountID int, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	acc, err := FindAccount(accountID)
	if err != nil {
		return err
	}
	acc.Balance += amount
	acc.TransactionHistory = append(acc.TransactionHistory, fmt.Sprintf("Deposited: %.2f", amount))
	return nil
}

func Withdraw(accountID int, amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be greater than zero")
	}
	acc, err := FindAccount(accountID)
	if err != nil {
		return err
	}
	if acc.Balance < amount {
		return errors.New("insufficient balance")
	}
	acc.Balance -= amount
	acc.TransactionHistory = append(acc.TransactionHistory, fmt.Sprintf("Withdrew: %.2f", amount))
	return nil
}

func ViewBalance(accountID int) (float64, error) {
	acc, err := FindAccount(accountID)
	if err != nil {
		return 0, err
	}
	return acc.Balance, nil
}

func ViewTransactionHistory(accountID int) ([]string, error) {
	acc, err := FindAccount(accountID)
	if err != nil {
		return nil, err
	}
	return acc.TransactionHistory, nil
}

func main() {
	accounts = append(accounts, Account{ID: 1, Name: "Vishal", Balance: 5000})
	accounts = append(accounts, Account{ID: 2, Name: "Nishant", Balance: 3000})

	for {
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. View Balance")
		fmt.Println("4. View Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		if choice == OptionExit {
			fmt.Println("Exiting the system. Goodbye!")
			break
		}

		var accountID int
		fmt.Print("Enter Account ID: ")
		fmt.Scanln(&accountID)

		switch choice {
		case OptionDeposit:
			fmt.Print("Enter deposit amount: ")
			var amount float64
			fmt.Scanln(&amount)
			if err := Deposit(accountID, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful.")
			}
		case OptionWithdraw:
			fmt.Print("Enter withdrawal amount: ")
			var amount float64
			fmt.Scanln(&amount)
			if err := Withdraw(accountID, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful.")
			}
		case OptionViewBalance:
			balance, err := ViewBalance(accountID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Current balance: %.2f\n", balance)
			}
		case OptionTransactionHistory:
			history, err := ViewTransactionHistory(accountID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Transaction History:")
				for _, entry := range history {
					fmt.Println(entry)
				}
			}
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
