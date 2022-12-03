package transactions

type TransactionMetadata struct {
	AffectedNodes          []AffectedNode `json:"AffectedNodes"`
	PartialDeliveredAmount CurrencyAmount `json:"DeliveredAmount"`
	TransactionIndex       uint64         `json:"TransactionIndex"`
	TransactionResult      string         `json:"TransactionResult"`
	DeliveredAmount        CurrencyAmount `json:"delivered_amount"`
}

type AffectedNode struct {
	CreatedNode  *CreatedNode  `json:"CreatedNode,omitempty"`
	ModifiedNode *ModifiedNode `json:"ModifiedNode,omitempty"`
	DeletedNode  *DeletedNode  `json:"DeletedNode,omitempty"`
}

type CreatedNode struct {
	LedgerEntryType string      `json:"LedgerEntryType,omitempty"`
	LedgerIndex     string      `json:"LedgerIndex,omitempty"`
	NewFields       interface{} `json:"NewFields,omitempty"`
}

type ModifiedNode struct {
	LedgerEntryType   string      `json:"LedgerEntryType"`
	LedgerIndex       string      `json:"LedgerIndex"`
	FinalFields       interface{} `json:"FinalFields"`
	PreviousFields    interface{} `json:"PreviousFields"`
	PreviousTxnID     string      `json:"PreviousTxnID,omitempty"`
	PreviousTxnLgrSeq uint64      `json:"PreviousTxnLgrSeq,omitempty"`
}

type DeletedNode struct {
	LedgerEntryType string      `json:"LedgerEntryType"`
	LedgerIndex     string      `json:"LedgerIndex"`
	FinalFields     interface{} `json:"FinalFields"`
}
