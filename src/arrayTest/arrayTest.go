package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	x[2] = 50
	fmt.Println(x)
	fmt.Println(x[2])

	avg()

	// declare arrays with values.
	y := [5]float64{98, 93, 77, 82, 83}
	fmt.Println(y)
}

func avg() {
	var x [5]float64

	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	fmt.Println()
	var total float64 = 0
	for i := 0; i < len(x); i++ {
		fmt.Print(x[i], " ")
		total += x[i]
	}
	fmt.Println()
	fmt.Println("Total:", total)
	fmt.Println("Average:", total/float64(len(x)))

	total = 0
	//for i, value := range x {		// error, "i" declared but not used.
	for _, value := range x {
		total += value
	}
	fmt.Println()
	fmt.Println("Total:", total)
	fmt.Println("Average:", total/float64(len(x)))
}
