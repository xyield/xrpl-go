package ledger

type FeeSettings struct {
	BaseFee           string
	Flags             uint32
	LedgerEntryType   LedgerEntryType
	ReferenceFeeUnits uint
	ReserveBase       uint
	ReserveIncrement  uint
}

func (*FeeSettings) EntryType() LedgerEntryType {
	return FeeSettingsEntry
}
