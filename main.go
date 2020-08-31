package main

import (
	"fmt"
	"sort"
)

func main() {
	s := stockBuySell(slice100000)
	fmt.Printf("%s\n", s.String())
}

func stockBuySell(prices []int) signal {
	maxProfit := maxDiff(prices)
	var changed = false
	var done = false
	var buy, sell int
	var buyValue = 99999
	var sellValue = 0
	for i, v := range prices {
		if done {
			break
		}
		for j, w := range prices {
			changed = false
			// Iterate over the part of the slice that comes after the element we're looking at for "buy".
			if j <= i {
				continue
			}

			// If the element is better than what we had, store it.
			if v < buyValue {
				buyValue = v
				buy = i
				changed = true
			}

			if w > sellValue {
				sell = j
				sellValue = w
				changed = true
			}

			// if we've reached peak profit, abort.
			if changed && sellValue - buyValue >= maxProfit{
				done = true
				break
			}
		}
	}

	// Return the current best option.
	return signal{
		buy: buy,
		bV: buyValue,
		sell: sell,
		sV: sellValue,
		profit: sellValue - buyValue,
	}
}

func maxDiff(in []int) int {
	// Remove duplicates
	normalized := unique(in)
	sort.Ints(normalized)
	return normalized[len(normalized) -1] - normalized[0]
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

type signal struct {
	buy int
	sell int
	bV int
	sV int
	profit int
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
func (s *signal) Profit() int {
	return s.profit
}

// String will return the human readable advice given the slice.
func (s *signal) String() string {
	return fmt.Sprintf("buy on day %d for %d, sell on day %d for %d for a profit of %d", s.buy + 1, s.bV, s.sell + 1, s.sV, s.profit)
}