# 31 August 2020

This week’s question:
Given an array of numbers that represent stock prices (where each number is the price for a certain day), find 2 days when you should buy and sell your stock for the highest profit.

Example:

```
$ stockBuySell([110, 180, 260, 40, 310, 535, 695])
$ “buy on day 4, sell on day 7”
```
## solutions considered but ultimately discarded

### binary search

Essentially I would recursively split the slice of ints (array for non-Go people), then check their minimum / maximum, and then discard one half if the other one was a better candidate

The reason I discarded that is if the slice ends up being super long, and the last two ints there are `[..., 9237492387, 1]` achieving the maximum diff, I would need to backtrack on the whole binary search again.

## bench test

```
goos: darwin
goarch: amd64
pkg: github.com/javorszky/cassidoo-31-aug-2020
Benchmark_stockBuySell10-16        	 1469340	       803 ns/op
Benchmark_stockBuySell100-16       	  119161	     10265 ns/op
Benchmark_stockBuySell1000-16      	   10000	    109100 ns/op
Benchmark_stockBuySell10000-16     	    2902	    450896 ns/op
Benchmark_stockBuySell100000-16    	     430	   2990930 ns/op
PASS
ok  	github.com/javorszky/cassidoo-31-aug-2020	7.370s
```