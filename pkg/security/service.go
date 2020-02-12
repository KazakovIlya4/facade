package security

import "fmt"

const (
	securityCodeExampleValid int = 123
	transactionIDExample     int = 1234
)

// Checker checks correctness of `securityCode` corresponding to `transactionID`
type Checker interface {
	Check(transactionID int, securityCode int) bool
}

type security struct{}

func (s *security) Check(transactionID int, securityCode int) bool {
	fmt.Printf("Checking security code %d\n", securityCode)

	//placeholder
	if securityCode == securityCodeExampleValid && transactionID == transactionIDExample {
		return true
	}
	return false
}

// NewChecker returns new instance of Checker implementation
func NewChecker() Checker {
	return &security{}
}
