package main

import (
	"fmt"
)

func main() {
	stocks := []uint32{110, 180, 260, 40, 310, 535, 695}

	stockBuySell(stocks)
}

func stockBuySell(prices []uint32) signal {
	//var max, buy uint32
	for i, v := range prices {
		fmt.Printf("index is %d, value is %d\n", i, v)
	}
	return signal{}
}

type signal struct {
	buy int
	sell int
	profit uint32
}


// BuyDay will return the day to buy.
func (s *signal) BuyDay() int {
	return s.buy
}


// SellDay will return the day to sell.
func (s *signal) SellDay() int {
	return s.sell
}

// Profit will return the profit between the two.
func (s *signal) Profit() uint32 {
	return s.profit
}

// String will return the human readable advice given the slice.
func (s *signal) String() string {
	return fmt.Sprintf("buy on day %d, sell on day %d for a profit of %d", s.buy, s.sell, s.profit)
}