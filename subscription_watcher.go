package pubcontrol

var (
	defaultSubscriptionWatcher SubscriptionWatcher = &AlwaysTrueSubscriptionWatcher{}
)

// Tracks subscribed channels. For Fanout, this would utilize: https://fanout.io/docs/devguide.html#subscription-feeds
type SubscriptionWatcher interface {
	Start()
	Stop()
	// Checks if subscriptions are being updated without errors. Useful for scenarios where you want
	// to deliver all messages when subscription status is not available, or just general monitoring.
	IsHealthy() bool

	// Checks if a channel is connected.
	IsChannelConnected(channel string) bool
}

type AlwaysTrueSubscriptionWatcher struct{}

func (w *AlwaysTrueSubscriptionWatcher) Start()                                 {}
func (w *AlwaysTrueSubscriptionWatcher) Stop()                                  {}
func (w *AlwaysTrueSubscriptionWatcher) IsHealthy() bool                        { return true }
func (w *AlwaysTrueSubscriptionWatcher) IsChannelConnected(channel string) bool { return true }
