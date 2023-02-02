package transactions

import "encoding/json"

type TicketCreate struct {
	BaseTx
	TicketCount uint
}

func (*TicketCreate) TxType() TxType {
	return TicketCreateTx
}

func UnmarshalTicketCreateTx(data json.RawMessage) (Tx, error) {
	var ret TicketCreate
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
