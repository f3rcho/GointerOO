package main

// functions variadicas
func sum(values ...int) {
	total := 0
	for _, value := range values {
		total += value
	}
	println(total)
}

func printNames(names ...string) {
	for _, name := range names {
		println(name)
	}
}

// returns with names
func getValues(x int) (doble int, triple int, quad int) {
	doble = x * 2
	triple = x * 3
	quad = x * 4
	return
}

func main() {
	sum(1, 2, 3)
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum(values...)

	printNames("John", "Jane", "Joe")

	println(getValues(2))
	//sum() // error: too few arguments in call to sum()
	//sum(1) // error: too few arguments in call to sum()
	//sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
}
