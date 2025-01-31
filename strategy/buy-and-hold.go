package strategy

import (
	gbt "github.com/SlashGordon/gobacktest"
	"github.com/SlashGordon/gobacktest/algo"
)

// BuyAndHold is a basic strategy, which creates a buy signal on every year change
func BuyAndHold() *gbt.Strategy {
	// create a new strategy with an algo stack and load into the backtest
	strategy := gbt.NewStrategy("buy-and-hold-yearly")

	strategy.SetAlgo(
		algo.RunYearly(),         // run on beginning of each year
		algo.CreateSignal("buy"), // always create a buy signal on a data event
	)

	return strategy
}
