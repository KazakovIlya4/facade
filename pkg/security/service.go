package security

import "fmt"

const (
	securityCodeExampleValid int = 123
)

// Checker checks correctness of `securityCode` corresponding to `transactionID`
type Checker interface {
	Check(transactionID int, securityCode int) bool
}

type security struct{}

func (s *security) Check(_ int, securityCode int) bool {
	fmt.Printf("Checking security code %d\n", securityCode)

	//placeholder
	if securityCode == securityCodeExampleValid {
		return true
	}
	return false
}

// NewChecker returns new instance of Checker implementation
func NewChecker() Checker {
	return &security{}
}
