package ledger

type NegativeUNL struct {
	DisabledValidators  []DisabledValidatorEntry `json:",omitempty"`
	Flags               uint32
	LedgerEntryType     LedgerEntryType
	ValidatorToDisable  string `json:",omitempty"`
	ValidatorToReEnable string `json:",omitempty"`
}

func (*NegativeUNL) EntryType() LedgerEntryType {
	return NegativeUNLEntry
}

type DisabledValidatorEntry struct {
	DisabledValidator DisabledValidator
}

type DisabledValidator struct {
	FirstLedgerSequence uint32
	PublicKey           string
}
