package security

type validator interface {
	Validate()
}

// Checker checks correctness of `securityCode` corresponding to `transactionID`
type Checker interface {
	Check(securityCode, transactionID int) bool
}

type security struct{}

// This function is a placeholder for an external call
func (s *security) Check(_, _ int) (valid bool) {
	valid = true
	return
}

// NewChecker returns new instance of Checker implementation
func NewChecker() Checker {
	return &security{}
}
