package main

func maxMessages(thresh int) int {
	maxMsg, totalCost := 0, 0
	for i := 0; ; i++ {
		totalCost = totalCost + 100 + (1 * i)
		maxMsg = i
		if totalCost >= thresh {
			return maxMsg
		}
	}
}
