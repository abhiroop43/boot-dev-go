package main

func getNameCounts(names []string) map[rune]map[string]int {
	returnValue := make(map[rune]map[string]int)
	for _, name := range names {
		if len(name) > 0 {
			firstChar := []rune(name)[0]
			if _, ok := returnValue[firstChar]; !ok {
				returnValue[firstChar] = make(map[string]int)
			}
			returnValue[firstChar][name]++
		}
	}
	return returnValue
}
