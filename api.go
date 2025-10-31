package cregis

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

const (
	CoinsURI                     = "/api/v1/coins"
	TradeURI                     = "/api/v1/trade/page"
	AddressCreateURI             = "/api/v1/address/create"
	AddressBatchCreateURI        = "/api/v1/batch/address/create"
	AddressInnerURI              = "/api/v1/address/inner"
	AddressLegalURI              = "/api/v1/address/legal"
	AddressUpdateURI             = "/api/v1/address/update"
	CollectionURI                = "/api/v1/collection"
	PayoutURI                    = "/api/v1/payout"
	PayoutV2URI                  = "/api/v2/payout"
	payoutQueryURI               = "/api/v1/payout/query"
	SubAddressBalanceURI         = "/api/v1/sub_address_balance"
	SubAddressWithdrawalURI      = "/api/v1/sub_address_withdrawal"
	SubAddressWithdrawalTradeURI = "/api/v1/sub_address_withdrawal/info"
)

type API struct {
	url    string
	apiKey string
	pid    int64
}

func NewClient(url, apiKey string, pid int64) *API {
	return &API{
		url:    url,
		apiKey: apiKey,
		pid:    pid,
	}
}

func (a *API) Coins() (*Result[Coins], error) {
	mp := make(map[string]any)
	mp["pid"] = a.pid
	resp, err := a.send(a.url+CoinsURI, mp)
	if err != nil {
		return nil, err
	}
	r := Result[Coins]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *API) Transactions(opts ...Option) (*Result[Page[[]Transaction]], error) {
	mp := make(map[string]any)
	mp["pid"] = a.pid
	for _, opt := range opts {
		opt(mp)
	}
	resp, err := a.send(a.url+TradeURI, mp)
	if err != nil {
		return nil, err
	}
	r := Result[Page[[]Transaction]]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *API) AddressCreate(chainId string, opts ...Option) (*Result[Address], error) {
	mp := make(map[string]any)
	mp["pid"] = a.pid
	mp["chain_id"] = chainId
	for _, opt := range opts {
		opt(mp)
	}
	resp, err := a.send(a.url+AddressCreateURI, mp)
	if err != nil {
		return nil, err
	}
	r := Result[Address]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *API) AddressBatchCreate(chainId, number string, opts ...Option) (*Result[Address], error) {
	mp := make(map[string]any)
	mp["pid"] = a.pid
	mp["chain_id"] = chainId
	mp["number"] = number
	for _, opt := range opts {
		opt(mp)
	}
	resp, err := a.send(a.url+AddressBatchCreateURI, mp)
	if err != nil {
		return nil, err
	}
	r := Result[Address]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *API) AddressInner(chainId, address string) (*Result[AddressInner], error) {
	mp := make(map[string]any)
	mp["pid"] = a.pid
	mp["address"] = address
	mp["chain_id"] = chainId
	resp, err := a.send(a.url+AddressInnerURI, mp)
	if err != nil {
		return nil, err
	}
	r := Result[AddressInner]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *API) AddressLegal(chainId, address string) (*Result[AddressLegal], error) {
	mp := make(map[string]any)
	mp["pid"] = a.pid
	mp["address"] = address
	mp["chain_id"] = chainId
	resp, err := a.send(a.url+AddressLegalURI, mp)
	if err != nil {
		return nil, err
	}
	r := Result[AddressLegal]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *API) AddressUpdate(address string, opts ...Option) (*Result[any], error) {
	mp := make(map[string]any)
	mp["pid"] = a.pid
	mp["address"] = address
	for _, opt := range opts {
		opt(mp)
	}
	resp, err := a.send(a.url+AddressUpdateURI, mp)
	if err != nil {
		return nil, err
	}
	r := Result[any]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *API) Collection(currency, fromAddress, toAddress string, opts ...Option) (*Result[Collection], error) {
	mp := make(map[string]any)
	mp["pid"] = a.pid
	mp["currency"] = currency
	mp["from_address"] = fromAddress
	mp["to_address"] = toAddress
	for _, opt := range opts {
		opt(mp)
	}
	resp, err := a.send(a.url+CollectionURI, mp)
	if err != nil {
		return nil, err
	}
	r := Result[Collection]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *API) Payout(currency, address, amount, thirdPartyId string, opts ...Option) (*Result[Payout], error) {
	mp := make(map[string]any)
	mp["pid"] = a.pid
	mp["currency"] = currency
	mp["address"] = address
	mp["amount"] = amount
	mp["third_party_id"] = thirdPartyId
	for _, opt := range opts {
		opt(mp)
	}
	resp, err := a.send(a.url+PayoutURI, mp)
	if err != nil {
		return nil, err
	}
	r := Result[Payout]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *API) PayoutV2(currency, toAddress, amount, thirdPartyId string, opts ...Option) (*Result[Payout], error) {
	mp := make(map[string]any)
	mp["pid"] = a.pid
	mp["currency"] = currency
	mp["to_address"] = toAddress
	mp["amount"] = amount
	mp["third_party_id"] = thirdPartyId
	for _, opt := range opts {
		opt(mp)
	}
	resp, err := a.send(a.url+PayoutV2URI, mp)
	if err != nil {
		return nil, err
	}
	r := Result[Payout]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *API) PayoutQuery(cid int64) (*Result[PayoutQuery], error) {
	mp := make(map[string]any)
	mp["pid"] = a.pid
	mp["cid"] = cid
	resp, err := a.send(a.url+payoutQueryURI, mp)
	if err != nil {
		return nil, err
	}
	r := Result[PayoutQuery]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *API) SubAddressBalance(currency string, opts ...Option) (*Result[Page[SubAddressBalance]], error) {
	mp := make(map[string]any)
	mp["pid"] = a.pid
	mp["currency"] = currency
	for _, opt := range opts {
		opt(mp)
	}
	resp, err := a.send(a.url+SubAddressBalanceURI, mp)
	if err != nil {
		return nil, err
	}
	r := Result[Page[SubAddressBalance]]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *API) SubAddressWithdrawal(currency, toAddress, amount, thirdPartyId string, opts ...Option) (*Result[SubAddressWithdrawal], error) {
	mp := make(map[string]any)
	mp["pid"] = a.pid
	mp["to_address"] = toAddress
	mp["currency"] = currency
	mp["amount"] = amount
	mp["third_party_id"] = thirdPartyId
	for _, opt := range opts {
		opt(mp)
	}
	resp, err := a.send(a.url+SubAddressWithdrawalURI, mp)
	if err != nil {
		return nil, err
	}
	r := Result[SubAddressWithdrawal]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *API) SubAddressWithdrawalTrade(cid int64) (*Result[PayoutQuery], error) {
	mp := make(map[string]any)
	mp["pid"] = a.pid
	mp["cid"] = cid
	resp, err := a.send(a.url+SubAddressWithdrawalTradeURI, mp)
	if err != nil {
		return nil, err
	}
	r := Result[PayoutQuery]{}
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *API) send(url string, params map[string]any) (string, error) {
	params["timestamp"] = time.Now().UnixMilli()
	params["nonce"] = generateRandomString(6)
	params["sign"] = Sign(params, a.apiKey)
	payload, _ := json.Marshal(params)
	body := strings.NewReader(string(payload))
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	r, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if !gjson.ParseBytes(r).Get("code").Exists() {
		return "", fmt.Errorf("%v", string(r))
	}
	return string(r), nil
}

func generateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLength := len(charset)
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomString[i] = charset[rand.Intn(charsetLength)]
	}
	return string(randomString)
}
