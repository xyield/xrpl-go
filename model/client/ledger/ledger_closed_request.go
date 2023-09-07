package ledger

type LedgerClosedRequest struct {
}

func (*LedgerClosedRequest) Method() string {
	return "ledger_closed"
}

func (*LedgerClosedRequest) Validate() error {
	return nil
}
