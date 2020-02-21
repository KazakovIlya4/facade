package main

import (
	"fmt"

	"facade/pkg/account"
	"facade/pkg/facade"
	"facade/pkg/operation"
)

const (
	withdrawCode = iota
	balanceCode
)

func main() {
	wallets := make(map[string]account.Wallet)
	wallets["AliceID"] = account.NewWallet("Alice", 500)
	wallets["BobID"] = account.NewWallet("Bob", 5000)

	transactionService := operation.NewService()

	operationCodes := make(map[string]int)
	operationCodes["withdraw"] = withdrawCode
	operationCodes["balance"] = balanceCode

	paymentSystem := facade.NewPaymentSystem(wallets, transactionService, operationCodes)
	balance, err := paymentSystem.Balance("AliceID")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s balance %d \n", "Alice", balance)

	err = paymentSystem.Withdraw("AliceID", 450)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Withdrew %d from %s \n", 450, "Alice")

	balance, err = paymentSystem.Balance("AliceID")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s balance %d\n", "Alice", balance)

}
