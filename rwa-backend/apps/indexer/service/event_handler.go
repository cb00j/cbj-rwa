package service

import (
	"context"

	coreTypes "github.com/cb00j/cbj-rwa/rwa-backend/libs/core/types"
	"gorm.io/gorm"
)

type EventHandler interface {
	// ContractType returns the type of contract the handler processes.
	ContractType() coreTypes.ContractType
	// Topic0 returns the topic0 of the event the handler handles.
	Topic0() string
	// HandleEvent handles a single event log. The tx parameter is the database
	// transaction that the handler MUST use for all DB operations to ensure atomicity.
	// Side effects that talk to anything outside this database (Alpaca, other
	// on-chain calls) must NOT happen here — the transaction has not committed
	// yet, so nothing written here is visible to any other connection. Put
	// those side effects in AfterCommit instead.
	HandleEvent(ctx context.Context, tx *gorm.DB, event *coreTypes.EventLogWithId) error

	// AfterCommit runs once the batch transaction containing this event's
	// HandleEvent call has been durably committed — i.e. every row this
	// handler wrote is now guaranteed visible to every other DB connection
	// (alpaca-stream, reconciler, etc). This is the correct place for any
	// action that will cause another system to start reacting based on our
	// data (e.g. calling Alpaca's PlaceOrder, which causes Alpaca to start
	// pushing WS events back at us for an order that must already be
	// queryable by the time those events arrive).
	//
	// Most handlers have nothing to do here and can just `return nil`.
	AfterCommit(ctx context.Context, event *coreTypes.EventLogWithId) error
}
