package main

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func addAndSubtract(a, b int) (x, y int) {
	x = add(a, b)
	y = subtract(a, b)
	return
}
