package ledger

type NegativeUNL struct {
	DisabledValidators  []DisabledValidator `json:",omitempty"`
	Flags               uint                `json:",omitempty"`
	LedgerEntryType     LedgerEntryType     `json:",omitempty"`
	ValidatorToDisable  string              `json:",omitempty"`
	ValidatorToReEnable string              `json:",omitempty"`
}

func (*NegativeUNL) EntryType() LedgerEntryType {
	return NegativeUNLEntry
}

type DisabledValidator struct {
	FirstLedgerSequence uint
	PublicKey           string
}
