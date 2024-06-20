package main

func getLast[T any](s []T) T {
	var zeroVal T
	sliceLength := len(s)
	if sliceLength > 0 {
		return s[sliceLength-1]
	}

	return zeroVal

}
