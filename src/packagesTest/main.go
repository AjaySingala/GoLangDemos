package main

import "fmt"

//import "packagesTest/math"
import m "packagesTest/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)
}
