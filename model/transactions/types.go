package transactions

// Types commonly used in requests
// TODO type validation and utility methods

type Address string

type XAddress string

type Hash256 []byte

type Hash128 []byte

// TODO XrpCurrencyAmount goes to/from rippled as string
type XrpCurrencyAmount uint64

type Marker any

type NFTokenID Hash256

type NFTokenURI string

type TransferFee uint
