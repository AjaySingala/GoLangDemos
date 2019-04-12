package main

import "fmt"

func main() {
	numbers := []float64{98, 93, 77, 82, 83}
	fmt.Println("Average is:", average(numbers))

	//foo()
	//foobar()

	// Return multiple values
	x, y := multipleRetiurnVals()
	fmt.Println(x, y)

	// Return value and error
	a, b := returnValErr()
	fmt.Println(a, b)

	// Return value and status (true/false)
	c, d := returnValOk()
	fmt.Println(c, d)
	c, d = returnValNotOk()
	fmt.Println(c, d)
}

func returnValNotOk() (int, bool) {
	return 0, false
}

func returnValOk() (int, bool) {
	return 1, true
}

func returnValErr() (int, string) {
	return -1, "This is an error"
}

func multipleRetiurnVals() (int, int) {
	fmt.Println("Returning multiple values from a function")
	return 10, 20
}

func foobar() {
	panic("This function not yet implemented")
}

func foo() {
	fmt.Println("Get numbers from user...")

	numbers2 := make([]float64, 5)
	var input1 float64
	var input2 float64
	var input3 float64
	var input4 float64
	var input5 float64

	fmt.Print("Enter 5 numbers: ")
	fmt.Scanf("%f %f %f %f %f", &input1, &input2, &input3, &input4, &input5)
	fmt.Println("You entered:", input1, input2, input3, input4, input5)
	numbers2[0] = input1
	numbers2[1] = input2
	numbers2[2] = input3
	numbers2[3] = input4
	numbers2[4] = input5
	fmt.Println(numbers2)
	fmt.Println("Average is:", average(numbers2))
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}

	fmt.Println("Total is:", total)
	return total / float64(len(xs))
}
