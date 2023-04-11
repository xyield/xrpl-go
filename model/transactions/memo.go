package transactions

type MemoWrapper struct {
	Memo Memo
}

type Memo struct {
	MemoData   string `json:",omitempty"`
	MemoFormat string `json:",omitempty"`
	MemoType   string `json:",omitempty"`
}
