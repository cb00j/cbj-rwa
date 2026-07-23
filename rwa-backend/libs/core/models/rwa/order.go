package rwa

import (
	"time"

	"github.com/shopspring/decimal"
)

// Order maps to public.orders
type Order struct {
	// the order identifying fields
	ID            uint64 `gorm:"column:id;primaryKey" json:"id"`
	ClientOrderID string `gorm:"column:client_order_id;index" json:"clientOrderId"` // used to send Alpaca as the unique identifier for the order provided by the client(this project)
	AccountID     uint64 `gorm:"column:account_id;index" json:"accountId"`          // foreign key to [Account.ID]
	EventLogID    uint64 `gorm:"column:event_log_id;index" json:"eventLogId"`       // foreign key to [EventLog.ID],the log id from the blockchain event emitted by the OrderContract

	// the order parameters
	Symbol      string          `gorm:"column:symbol;index" json:"symbol"` // stock symbol, e.g., AAPL.cbj, TSLA.cbj
	AssetType   AssetType       `gorm:"column:asset_type" json:"assetType"`
	Side        OrderSide       `gorm:"column:side" json:"side"`         // Buy/Sell
	Type        OrderType       `gorm:"column:type" json:"type"`         // Market/Limit/...
	Quantity    decimal.Decimal `gorm:"column:quantity" json:"quantity"` // the quantity of stock to buy/sell
	Price       decimal.Decimal `gorm:"column:price" json:"price"`
	StopPrice   decimal.Decimal `gorm:"column:stop_price" json:"stopPrice"`
	TimeInForce string          `gorm:"column:time_in_force" json:"timeInForce"` // DAY/GTC/IOC...,see Order.sol for details
	Status      OrderStatus     `gorm:"column:status" json:"status"`

	// the order execution details
	FilledQuantity    decimal.Decimal `gorm:"column:filled_quantity" json:"filledQuantity"`       // accumulated filled quantity
	FilledPrice       decimal.Decimal `gorm:"column:filled_price" json:"filledPrice"`             // average filled price calculated by VWAP
	RemainingQuantity decimal.Decimal `gorm:"column:remaining_quantity" json:"remainingQuantity"` // remaining quantity to be filled

	// on-chain transaction details
	EscrowAmount   decimal.Decimal `gorm:"column:escrow_amount" json:"escrowAmount"`            // the amount of escrowed in the OrderContract for this order
	EscrowAsset    string          `gorm:"column:escrow_asset" json:"escrowAsset"`              // the asset of the escrowed in the OrderContract for this order, e.g., USDM, AAPL.cbj, TSLA.cbj
	RefundAmount   decimal.Decimal `gorm:"column:refund_amount" json:"refundAmount"`            // if the order is canceled or partially filled, the remaining escrowed amount will be refunded to the user
	ContractTxHash string          `gorm:"column:contract_tx_hash;index" json:"contractTxHash"` // submitOrder transaction hash on the OrderContract
	ExecuteTxHash  string          `gorm:"column:execute_tx_hash" json:"executeTxHash"`         // markExecuted transaction hash on the OrderContract
	CancelTxHash   string          `gorm:"column:cancel_tx_hash" json:"cancelTxHash"`

	// the brokerage provider details
	ExternalOrderID string `gorm:"column:external_order_id" json:"externalOrderId"` // the order id from the brokerage provider
	Provider        string `gorm:"column:provider" json:"provider"`                 // the brokerage provider, e.g., Alpaca, Binance, etc.

	// fee details
	Commission      string `gorm:"column:commission" json:"commission"`
	CommissionAsset string `gorm:"column:commission_asset" json:"commissionAsset"` // the asset of the commission

	Metadata string `gorm:"column:metadata;type:text" json:"metadata"` // additional metadata for the order
	Notes    string `gorm:"column:notes;type:text" json:"notes"`       // additional notes for the order

	// timestamps maintained by GORM
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`

	// records the timestamps of the order status changes
	SubmittedAt *time.Time `gorm:"column:submitted_at" json:"submittedAt"`
	AcceptedAt  *time.Time `gorm:"column:accepted_at" json:"acceptedAt"`
	FilledAt    *time.Time `gorm:"column:filled_at" json:"filledAt"`
	CancelledAt *time.Time `gorm:"column:cancelled_at" json:"cancelledAt"`
	ExpiredAt   *time.Time `gorm:"column:expired_at" json:"expiredAt"`

	// reconciler tracking fields
	SettlementAttempts      int        `gorm:"column:settlement_attempts" json:"settlementAttempts"` // the number of settlement attempts made for this order
	LastSettlementAttemptAt *time.Time `gorm:"column:last_settlement_attempt_at" json:"lastSettlementAttemptAt"`
	LastSettlementError     string     `gorm:"column:last_settlement_error" json:"lastSettlementError"`
	NeedsManualReview       bool       `gorm:"column:needs_manual_review" json:"needsManualReview"` // whether the order needs manual review due to repeated settlement failures
}

// OrderExecution maps to public.order_executions
type OrderExecution struct {
	ID              uint64          `gorm:"column:id;primaryKey" json:"id"`
	OrderID         uint64          `gorm:"column:order_id;index" json:"orderId"`
	ExecutionID     string          `gorm:"column:execution_id" json:"executionId"`
	Quantity        decimal.Decimal `gorm:"column:quantity" json:"quantity"`
	Price           decimal.Decimal `gorm:"column:price" json:"price"`
	Commission      string          `gorm:"column:commission" json:"commission"`
	CommissionAsset string          `gorm:"column:commission_asset" json:"commissionAsset"`
	Provider        string          `gorm:"column:provider" json:"provider"`
	ExternalID      string          `gorm:"column:external_id" json:"externalId"`
	ExecutedAt      time.Time       `gorm:"column:executed_at" json:"executedAt"`
	CreatedAt       time.Time       `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

// OrderFilter represents order filter conditions
type OrderFilter struct {
	AccountID string    `json:"account_id"`
	Symbol    string    `json:"symbol"`
	Side      string    `json:"side"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	Source    string    `json:"source"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Limit     int       `json:"limit"`
	Offset    int       `json:"offset"`
}

func (*Order) TableName() string { return "orders" }

func (*OrderExecution) TableName() string { return "order_executions" }
