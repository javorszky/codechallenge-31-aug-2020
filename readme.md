# 31 August 2020

This week’s question:
Given an array of numbers that represent stock prices (where each number is the price for a certain day), find 2 days when you should buy and sell your stock for the highest profit.

Example:

```
$ stockBuySell([110, 180, 260, 40, 310, 535, 695])
$ “buy on day 4, sell on day 7”
```

## bench test

```
goos: darwin
goarch: amd64
pkg: github.com/javorszky/cassidoo-31-aug-2020
Benchmark_stockBuySell10-16        	 1501130	       798 ns/op
Benchmark_stockBuySell100-16       	   67050	     18174 ns/op
Benchmark_stockBuySell1000-16      	    3284	    362877 ns/op
Benchmark_stockBuySell10000-16     	      37	  30012136 ns/op
Benchmark_stockBuySell100000-16    	      15	  69868156 ns/op
PASS
ok  	github.com/javorszky/cassidoo-31-aug-2020	6.919s
```