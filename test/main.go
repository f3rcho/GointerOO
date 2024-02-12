package main

func Sum(a int, b int) int {
	return a + b
}

func GetMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Fibonnacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonnacci(n-1) + Fibonnacci(n-2)
}
