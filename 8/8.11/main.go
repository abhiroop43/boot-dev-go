package main

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	largestDay := 0

	for _, cost := range costs {
		if cost.day > largestDay {
			largestDay = cost.day
		}
	}

	totalCostsEachDay := make([]float64, largestDay+1)

	for _, cost := range costs {
		totalCostsEachDay[cost.day] += cost.value
	}

	return totalCostsEachDay
}

func createMatrix(rows, cols int) [][]int {
	mat := make([][]int, rows)

	for i := range mat {
		mat[i] = make([]int, cols)
	}

	return mat
}
