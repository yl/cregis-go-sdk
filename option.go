package cregis

type Option func(map[string]any)

func WithCid(cid int64) Option {
	return func(m map[string]any) {
		m["cid"] = cid
	}
}

func WithTxId(txId string) Option {
	return func(m map[string]any) {
		m["txid"] = txId
	}
}

func WithStatus(status int) Option {
	return func(m map[string]any) {
		m["status"] = status
	}
}

func WithTradeType(tradeType int) Option {
	return func(m map[string]any) {
		m["trade_type"] = tradeType
	}
}

func WithBusinessType(businessType int) Option {
	return func(m map[string]any) {
		m["business_type"] = businessType
	}
}

func WithBlockTimeStart(blockTimeStart int64) Option {
	return func(m map[string]any) {
		m["blocktime_start"] = blockTimeStart
	}
}

func WithBlockTimeEnd(blockTimeEnd int64) Option {
	return func(m map[string]any) {
		m["blocktime_end"] = blockTimeEnd
	}
}

func WithChainId(chainId string) Option {
	return func(m map[string]any) {
		m["chain_id"] = chainId
	}
}

func WithTokenId(tokenId string) Option {
	return func(m map[string]any) {
		m["token_id"] = tokenId
	}
}

func WithPageNum(pageNum int) Option {
	return func(m map[string]any) {
		m["page_num"] = pageNum
	}
}

func WithPageSize(pageSize int) Option {
	return func(m map[string]any) {
		m["page_size"] = pageSize
	}
}

func WithAlias(alias string) Option {
	return func(m map[string]any) {
		m["alias"] = alias
	}
}

func WithCallbackUrl(callbackUrl string) Option {
	return func(m map[string]any) {
		m["callback_url"] = callbackUrl
	}
}

func WithAmount(amount string) Option {
	return func(m map[string]any) {
		m["amount"] = amount
	}
}

func WithMaximumBalance(maximumBalance string) Option {
	return func(m map[string]any) {
		m["maximum_balance"] = maximumBalance
	}
}

func WithMinimumBalance(minimumBalance string) Option {
	return func(m map[string]any) {
		m["minimum_balance"] = minimumBalance
	}
}

func WithRemark(remark string) Option {
	return func(m map[string]any) {
		m["remark"] = remark
	}
}

func WithMemo(memo string) Option {
	return func(m map[string]any) {
		m["memo"] = memo
	}
}

func WithFromAddress(fromAddress string) Option {
	return func(m map[string]any) {
		m["from_address"] = fromAddress
	}
}
