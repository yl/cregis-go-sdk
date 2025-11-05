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
	BlockTime    int64            `json:"block_time"`
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
	BlockTime   int64             `json:"block_time"`
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
	BlockTime    int64            `json:"block_time"`
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
	BlockTime    int64            `json:"block_time"`
	Nonce        string           `json:"nonce"`
	Timestamp    int64            `json:"timestamp"`
	Sign         string           `json:"sign"`
}
