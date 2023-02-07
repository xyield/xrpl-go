package ledger

type FeeSettings struct {
	BaseFee           string
	Flags             uint
	LedgerEntryType   string
	ReferenceFeeUnits uint
	ReserveBase       uint
	ReserveIncrement  uint
}

func (*FeeSettings) EntryType() LedgerEntryType {
	return FeeSettingsEntry
}
