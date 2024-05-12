package main

func bulkSend(numMessages int) float64 {
	totalCost := 0.00
	for i := 0; i < numMessages; i++ {
		totalCost = totalCost + 1.00 + (0.01 * float64(i))
	}

	return totalCost
}
