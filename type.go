package cregis

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const SuccessCode = "00000"

const SuccessWebhookResponse = "success"

type TransactionStatus int

func (s *TransactionStatus) UnmarshalJSON(bytes []byte) error {
	var field any
	if err := json.Unmarshal(bytes, &field); err != nil {
		return err
	}
	switch v := field.(type) {
	case string:
		n, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*s = TransactionStatus(n)
	case float64:
		*s = TransactionStatus(int(v))
	default:
		return fmt.Errorf("unexpected type: %T", v)
	}
	return nil
}

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

type BlockTime int

func (b *BlockTime) UnmarshalJSON(bytes []byte) error {
	var field any
	if err := json.Unmarshal(bytes, &field); err != nil {
		return err
	}
	switch v := field.(type) {
	case nil:
		*b = 0
		return nil
	case string:
		n, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*b = BlockTime(n)
	case float64:
		*b = BlockTime(int(v))
	default:
		return fmt.Errorf("unexpected type: %T", v)
	}
	return nil
}

type Result[T any] struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type Page[T any] struct {
	Total    int64 `json:"total"`
	PageNum  int64 `json:"page_num"`
	PageSize int64 `json:"page_size"`
	Rows     T     `json:"rows"`
}

type Coin struct {
	CoinName string `json:"coin_name"`
	ChainId  string `json:"chain_id"`
	TokenId  string `json:"token_id"`
}

type Coins struct {
	PayoutCoins  []*Coin `json:"payout_coins"`
	AddressCoins []*Coin `json:"address_coins"`
	OrderCoins   []*Coin `json:"order_coins"`
}

type Transaction struct {
	Pid         int64             `json:"pid"`
	Cid         int64             `json:"cid"`
	ChainId     string            `json:"chain_id"`
	TokenId     string            `json:"token_id"`
	Currency    string            `json:"currency"`
	FromAddress string            `json:"from_address"`
	ToAddress   string            `json:"to_address"`
	Amount      string            `json:"amount"`
	Status      TransactionStatus `json:"status"`
	TxId        string            `json:"txid"`
	BlockHeight string            `json:"block_height"`
	BlockTime   BlockTime         `json:"block_time"`
	Fee         string            `json:"fee"`
}

type AddressLegal struct {
	Result bool `json:"result"`
}

type AddressInner struct {
	Result bool `json:"result"`
}

type Collection struct {
	Cid int64 `json:"cid"`
}

type Payout struct {
	Cid int64 `json:"cid"`
}

type PayoutQuery struct {
	Pid          int64            `json:"pid"`
	Address      string           `json:"address"`
	FromAddress  string           `json:"from_address"`
	ChainId      string           `json:"chain_id"`
	TokenId      string           `json:"token_id"`
	Currency     string           `json:"currency"`
	Amount       string           `json:"amount"`
	ThirdPartyId string           `json:"third_party_id"`
	Remark       string           `json:"remark"`
	TxId         string           `json:"txid"`
	BlockTime    BlockTime        `json:"block_time"`
	BlockHeight  string           `json:"block_height"`
	Memo         string           `json:"memo"`
	Status       WithdrawalStatus `json:"status"`
}

type Address struct {
	Address string `json:"address"`
}

type SubAddressBalance struct {
	Pid        int64  `json:"pid"`
	Address    string `json:"address"`
	Currency   string `json:"currency"`
	Total      string `json:"total"`
	Available  string `json:"available"`
	Processing string `json:"processing"`
}

type SubAddressWithdrawal struct {
	Cid int64 `json:"cid"`
}

type DepositCallback struct {
	Pid         int64             `json:"pid"`
	Cid         int64             `json:"cid"`
	ChainId     string            `json:"chain_id"`
	TokenId     string            `json:"token_id"`
	Currency    string            `json:"currency"`
	Address     string            `json:"address"`
	Amount      string            `json:"amount"`
	Status      TransactionStatus `json:"status"`
	TxId        string            `json:"txid"`
	BlockHeight string            `json:"block_height"`
	BlockTime   BlockTime         `json:"block_time"`
	Nonce       string            `json:"nonce"`
	Timestamp   int64             `json:"timestamp"`
	Sign        string            `json:"sign"`
}

type WithdrawalCallback struct {
	Pid          int64            `json:"pid"`
	Cid          int64            `json:"cid"`
	ChainId      string           `json:"chain_id"`
	TokenId      string           `json:"token_id"`
	Currency     string           `json:"currency"`
	Address      string           `json:"address"`
	Amount       string           `json:"amount"`
	ThirdPartyId string           `json:"third_party_id"`
	Remark       string           `json:"remark"`
	Status       WithdrawalStatus `json:"status"`
	TxId         string           `json:"txid"`
	BlockHeight  string           `json:"block_height"`
	BlockTime    BlockTime        `json:"block_time"`
	Nonce        string           `json:"nonce"`
	Timestamp    int64            `json:"timestamp"`
	Sign         string           `json:"sign"`
}

type SubAddressWithdrawalCallback struct {
	Pid          int64            `json:"pid"`
	Cid          int64            `json:"cid"`
	ChainId      string           `json:"chain_id"`
	TokenId      string           `json:"token_id"`
	Currency     string           `json:"currency"`
	FromAddress  string           `json:"from_address"`
	ToAddress    string           `json:"to_address"`
	Amount       string           `json:"amount"`
	ThirdPartyId string           `json:"third_party_id"`
	Remark       string           `json:"remark"`
	Status       WithdrawalStatus `json:"status"`
	TxId         string           `json:"txid"`
	BlockHeight  string           `json:"block_height"`
	BlockTime    BlockTime        `json:"block_time"`
	Nonce        string           `json:"nonce"`
	Timestamp    int64            `json:"timestamp"`
	Sign         string           `json:"sign"`
}
