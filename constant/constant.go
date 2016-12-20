package constant

// error constants
const (
	ErrAuthorizationError      = "Authorization Error"
	ErrInsufficientPermissions = "Insufficient Permissions"
)

// exchange types
const (
	OkCoinCn     = "okcoin.cn"
	Huobi        = "huobi"
	Poloniex     = "poloniex"
	Btcc         = "btcc"
	Chbtc        = "chbtc"
	OkcoinFuture = "okcoin.future"
	OandaV20     = "oanda.v20"
)

// log types
const (
	ERROR  = -1
	INFO   = 0
	PROFIT = 1
	BUY    = 2
	SELL   = 3
	CANCEL = 4
)

// delete log time types
const (
	LastTime = "0"
	Day      = "1"
	Week     = "2"
	Month    = "3"
)

// trade types
const (
	TradeTypeBuy        = "BUY"
	TradeTypeSell       = "SELL"
	TradeTypeLong       = "LONG"
	TradeTypeShort      = "SHORT"
	TradeTypeLongClose  = "LONG_CLOSE"
	TradeTypeShortClose = "SHORT_CLOSE"
)

// stock types (will useless)
const (
	BTC = "BTC"
	LTC = "LTC"
)

// some variables
var (
	Consts        = []string{"BTC", "LTC", "M", "M5", "M15", "M30", "H", "D", "W"}
	ExchangeTypes = []string{OkCoinCn, Huobi, Poloniex, Btcc, Chbtc, OkcoinFuture, OandaV20}
)
