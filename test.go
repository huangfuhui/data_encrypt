package main

import "fmt"

func countDigitOne(n int) int {
	start := 1
	count := 0
	for n >= start {
		count += n/(start*10) * start
		left := n%(start*10)
		if left < start {
			start *= 10
			continue
		}
		if left >= (start * 2) {
			count += start
		} else {
			count += left - start + 1
		}
		start *= 10
	}
	return count
}

func main() {
	res := countDigitOne(824883294)
	fmt.Println(res)
}
