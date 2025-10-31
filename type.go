package cregis

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
	BlockTime   int64             `json:"block_time"`
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
	ChainId      string           `json:"chain_id"`
	TokenId      string           `json:"token_id"`
	Currency     string           `json:"currency"`
	Amount       string           `json:"amount"`
	ThirdPartyId string           `json:"third_party_id"`
	Remark       string           `json:"remark"`
	Status       WithdrawalStatus `json:"status"`
	TxId         string           `json:"txid"`
	BlockHeight  string           `json:"block_height"`
	BlockTime    string           `json:"block_time"`
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
	Pid         int64             `json:"pid" mapstructure:"pid"`
	Cid         int64             `json:"cid" mapstructure:"cid"`
	ChainId     string            `json:"chain_id" mapstructure:"chain_id"`
	TokenId     string            `json:"token_id" mapstructure:"token_id"`
	Currency    string            `json:"currency" mapstructure:"currency"`
	Address     string            `json:"address" mapstructure:"address"`
	Amount      string            `json:"amount" mapstructure:"amount"`
	Status      TransactionStatus `json:"status" mapstructure:"status"`
	TxId        string            `json:"txid" mapstructure:"txid"`
	BlockHeight string            `json:"block_height" mapstructure:"block_height"`
	BlockTime   int64             `json:"block_time" mapstructure:"block_time"`
	Nonce       string            `json:"nonce" mapstructure:"nonce"`
	Timestamp   int64             `json:"timestamp" mapstructure:"timestamp"`
	Sign        string            `json:"sign" mapstructure:"-"`
}

type WithdrawalCallback struct {
	Pid          int64            `json:"pid" mapstructure:"pid"`
	Cid          int64            `json:"cid" mapstructure:"cid"`
	ChainId      string           `json:"chain_id" mapstructure:"chain_id"`
	TokenId      string           `json:"token_id" mapstructure:"token_id"`
	Currency     string           `json:"currency" mapstructure:"currency"`
	Address      string           `json:"address" mapstructure:"address"`
	Amount       string           `json:"amount" mapstructure:"amount"`
	ThirdPartyId string           `json:"third_party_id" mapstructure:"third_party_id"`
	Remark       string           `json:"remark" mapstructure:"remark"`
	Status       WithdrawalStatus `json:"status" mapstructure:"status"`
	TxId         string           `json:"txid" mapstructure:"txid"`
	BlockHeight  string           `json:"block_height" mapstructure:"block_height"`
	BlockTime    int64            `json:"block_time" mapstructure:"block_time"`
	Nonce        string           `json:"nonce" mapstructure:"nonce"`
	Timestamp    int64            `json:"timestamp" mapstructure:"timestamp"`
	Sign         string           `json:"sign" mapstructure:"-"`
}

type SubAddressWithdrawalCallback struct {
	Pid          int64            `json:"pid" mapstructure:"pid"`
	Cid          int64            `json:"cid" mapstructure:"cid"`
	ChainId      string           `json:"chain_id" mapstructure:"chain_id"`
	TokenId      string           `json:"token_id" mapstructure:"token_id"`
	Currency     string           `json:"currency" mapstructure:"currency"`
	FromAddress  string           `json:"from_address" mapstructure:"from_address"`
	ToAddress    string           `json:"to_address" mapstructure:"to_address"`
	Amount       string           `json:"amount" mapstructure:"amount"`
	ThirdPartyId string           `json:"third_party_id" mapstructure:"third_party_id"`
	Remark       string           `json:"remark" mapstructure:"remark"`
	Status       WithdrawalStatus `json:"status" mapstructure:"status"`
	TxId         string           `json:"txid" mapstructure:"txid"`
	BlockHeight  string           `json:"block_height" mapstructure:"block_height"`
	BlockTime    int64            `json:"block_time" mapstructure:"block_time"`
	Nonce        string           `json:"nonce" mapstructure:"nonce"`
	Timestamp    int64            `json:"timestamp" mapstructure:"timestamp"`
	Sign         string           `json:"sign" mapstructure:"-"`
}
