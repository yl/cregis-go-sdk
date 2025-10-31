package cregis

const SuccessCode = "00000"

const SuccessWebhookResponse = "success"

type TransactionStatus int

const (
	TransactionStatusPending TransactionStatus = 0
	TransactionStatusSucceed TransactionStatus = 1
	TransactionStatusFailed  TransactionStatus = 2
)

type WithdrawalStatus int

const (
	PayoutStatusReviewing          WithdrawalStatus = 0
	PayoutStatusSignPassed         WithdrawalStatus = 1
	PayoutStatusSignRejected       WithdrawalStatus = 2
	PayoutStatusReviewCancelled    WithdrawalStatus = 3
	PayoutStatusReviewRejected     WithdrawalStatus = 4
	PayoutStatusSigning            WithdrawalStatus = 5
	PayoutStatusTransactionSucceed WithdrawalStatus = 6
	PayoutStatusTransactionFailed  WithdrawalStatus = 7
)

type AddressStatus int

const (
	AddressStatusEnabled  AddressStatus = 0
	AddressStatusDisabled AddressStatus = 1
)

type TransactionType int

const (
	TransactionTypeIn  TransactionType = 1
	TransactionTypeOut TransactionType = 2
)

type BusinessType int

const (
	BusinessTypeSimple        BusinessType = 0
	BusinessTypeReview        BusinessType = 2
	BusinessTypeDeposit       BusinessType = 3
	BusinessTypeCollectionFee BusinessType = 4
	BusinessTypeCollection    BusinessType = 5
)
