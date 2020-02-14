package models

type Record = struct {
	ID            int
	OperationCode uint8
	UserID        string
	Amount        uint32
	Success       bool
}
