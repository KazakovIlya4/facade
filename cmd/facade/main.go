package main

import (
	"fmt"

	"facade/pkg/facade"
	"facade/pkg/security"
	"facade/pkg/wallet"
)

const (
	securityCodeExampleValid int = 123
	transactionIDExample     int = 1234
)

func main() {
	paymentSystem := facade.NewPaymentSystem([]facade.AccountInfo{
		facade.Account("Alice", wallet.NewWallet("Alice", 500)),
	}, security.NewChecker())
	balance, err := paymentSystem.GetBalance("Alice", securityCodeExampleValid, transactionIDExample)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s Balance is %d\n", "Alice", balance)

	err = paymentSystem.GetMoney("Alice", 450, securityCodeExampleValid, transactionIDExample)
	if err != nil {
		fmt.Println(err)
	}

	balance, err = paymentSystem.GetBalance("Alice", securityCodeExampleValid, transactionIDExample)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s Balance is %d\n", "Alice", balance)

}
