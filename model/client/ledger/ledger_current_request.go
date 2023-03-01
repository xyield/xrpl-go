package ledger

type LedgerCurrentRequest struct {
}

func (*LedgerCurrentRequest) Method() string {
	return "ledger_current"
}
