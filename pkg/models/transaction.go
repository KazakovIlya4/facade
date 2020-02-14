package models

// Record keeps operation info
type Record = struct {
	ID            int
	OperationCode int
	UserID        string
	Amount        int
	Success       bool
}
