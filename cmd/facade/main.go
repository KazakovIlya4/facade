package main

import (
	"fmt"

	"facade/pkg/facade"
	"facade/pkg/transaction"
	"facade/pkg/wallet"
)

func main() {
	wallets := make(map[string]wallet.Wallet)
	wallets["AliceID"] = wallet.NewWallet("Alice", 500)
	wallets["BobID"] = wallet.NewWallet("Bob", 5000)

	securityChecker := transaction.NewChecker()

	paymentSystem := facade.NewPaymentSystem(wallets, securityChecker)
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
