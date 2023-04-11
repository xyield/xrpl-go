package stream

type StreamType string

const (
	LedgerStreamType      StreamType = "ledgerClosed"
	ValidationStreamType  StreamType = "validationReceived"
	TransactionStreamType StreamType = "transaction"
	PeerStatusStreamType  StreamType = "peerStatusChange"
	// TODO example lists OrderBookStreamType as "transaction"
	OrderBookStreamType StreamType = TransactionStreamType
	ConsensusStreamType StreamType = "consensusPhase"
)
