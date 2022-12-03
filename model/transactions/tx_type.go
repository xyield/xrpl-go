package transactions

type TxType string

const (
	AccountSetTx           TxType = "AccountSet"
	AccountDeleteTx               = "AccountDelete"
	CheckCancelTx                 = "CheckCancel"
	CheckCashTx                   = "CheckCash"
	CheckCreateTx                 = "CheckCreate"
	DepositPreauthTx              = "DepositPreauth"
	EscrowCancelTx                = "EscrowCancel"
	EscrowCreateTx                = "EscrowCreate"
	EscrowFinishTx                = "EscrowFinish"
	OfferCreateTx                 = "OfferCreate"
	OfferCancelTx                 = "OfferCancel"
	PaymentTx                     = "Payment"
	PaymentChannelClaimTx         = "PaymentChannelClaim"
	PaymentChannelCreateTx        = "PaymentChannelCreate"
	PaymentChannelFundTx          = "PaymentChannelFund"
	SetRegularKeyTx               = "SetRegularKey"
	SignerListSetTx               = "SignerListSet"
	TrustSetTx                    = "TrustSet"
	TicketCreateTx                = "TicketCreate"
)
