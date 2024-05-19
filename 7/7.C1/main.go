package main

func getPacketSize(message string) int {
	messageLen := len(message)

	if isPrime(messageLen) {
		return 0
	}

	for i := 1; ; i++ {
		if messageLen%i == 0 {
			return (messageLen / i) / 4
		}
	}
}

func isPrime(num int) bool {
	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
