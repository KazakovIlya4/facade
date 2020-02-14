package models

// Record keeps transaction info
type Record = struct {
	ID            int
	OperationCode int
	UserID        string
	Amount        int
	Success       bool
}
