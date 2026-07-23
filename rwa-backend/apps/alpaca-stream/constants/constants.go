package constants

const (
	StreamTypeTradeUpdates = "trade_updates"
	StreamTypeBars         = "bars"
	StreamTypeQuotes       = "quotes"
	StreamTypeTrades       = "trades"
)

const (
	EventTypeAccepted        = "accepted" // order has been accepted by the exchange,but not routed to the market yet（Non-trading hours）
	EventTypeNew             = "new"      // order has been routed to the market and ready to be filled
	EventTypeFill            = "fill"
	EventTypePartialFill     = "partial_fill"
	EventTypeCanceled        = "canceled"
	EventTypeExpired         = "expired"
	EventTypeRejected        = "rejected"
	EventTypeReplaced        = "replaced"
	EventTypePendingNew      = "pending_new"
	EventTypePendingCancel   = "pending_cancel"
	EventTypePendingReplace  = "pending_replace"
	EventTypeCancelRejected  = "cancel_rejected"
	EventTypeDoneForDay      = "done_for_day"
	EventTypeReplaceRejected = "replace_rejected"
)

const (
	FeedIEX = "iex"
	FeedSIP = "sip"
)

const EnableTradeUpdates = true

const (
	DefaultMarketDataFeed = FeedIEX
)

const (
	DefaultReconnectDelay       = 1
	DefaultMaxReconnectDelay    = 30
	DefaultMaxReconnectAttempts = -1
)

const (
	AuthTimeout   = 5
	ReadDeadline  = 300 // 5 分钟，避免非交易时段频繁断连
	WriteDeadline = 10
	PingInterval  = 120 // 2 分钟发送一次 ping 保活
)
