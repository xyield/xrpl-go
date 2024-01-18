package server

type LedgerAcceptRequest struct {
}

func (*LedgerAcceptRequest) Method() string {
	return "ledger_accept"
}

func (*LedgerAcceptRequest) Validate() error {
	return nil
}
