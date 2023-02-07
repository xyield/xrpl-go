package ledger

type NegativeUNL struct {
	DisabledValidators  []DisabledValidator
	Flags               uint
	LedgerEntryType     LedgerEntryType
	ValidatorToDisable  string
	ValidatorToReEnable string
}

func (*NegativeUNL) EntryType() LedgerEntryType {
	return NegativeUNLEntry
}

type DisabledValidator struct {
	FirstLedgerSequence uint
	PublicKey           string
}
