package server

type LedgerAcceptRequest struct {
}

func (*LedgerAcceptRequest) Method() string {
	return "leder_accept"
}
