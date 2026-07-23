package service

import "go.uber.org/fx"

func LoadModule() fx.Option {
	return fx.Module("service",
		fx.Provide(
			NewOnchainSettler,
			NewOnchainReconciler,
			NewFailedEventReconciler,
		),
		fx.Invoke(
			// Both reconcilers register their own fx.Lifecycle hooks inside
			// their constructors — they must be force-constructed here or
			// fx will never call those constructors at all (the same bug we
			// hit earlier with AlpacaWebSocketService).
			func(_ *OnchainReconciler) {},
			func(_ *FailedEventReconciler) {},
		),
	)
}
