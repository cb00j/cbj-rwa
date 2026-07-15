package ws

// SubscriptionType represents the type of subscription
type SubscriptionType string

type SubscriptionConfig struct {
	Type    SubscriptionType
	Symbols []string
	Handler MessageHandler
}
