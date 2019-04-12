package main

import "fmt"

func main() {
	var x []float64     // no length provided to a "slice"
	fmt.Println(len(x)) // ZERO
	x = make([]float64, 2)
	x[0] = 1
	fmt.Println(x)

	y := make([]float64, 5) // better way to create a slice.
	fmt.Println(len(y))     // FIVE

	z := make([]float64, 5, 10)
	fmt.Println(len(z)) // FIVE. 10 is the capacity of the slize z.

	// slice[low:high] format
	// low: index of where to start.
	// high: index of where to end, not including the index itself.
	// Can omit low and high.
	arr := [5]float64{1, 2, 3, 4, 5} // array.
	arrSlice := arr[0:5]             // 1,2,3,4,5
	fmt.Println(arrSlice)
	arrSlice = arr[1:4] // 2,3,4
	fmt.Println(arrSlice)
	arrSlice = arr[0:]
	fmt.Println(arrSlice) // 1,2,3,4,5
	arrSlice = arr[:3]
	fmt.Println(arrSlice) // 1,2,3
	arrSlice = arr[0:len(arr)]
	fmt.Println(arrSlice) // 1,2,3,4,5
	arrSlice = arr[:]
	fmt.Println(arrSlice) // 1,2,3,4,5
	arrSlice = arr[:5]
	fmt.Println(arrSlice) // 1,2,3,4,5

	// Slice Functions.
	sliceFunctions()
}

func sliceFunctions() {
	fmt.Println("Slice functions...")
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println("Since slice4 only has capacity for 2 elements, only first two elements are copied from slice3")
	fmt.Println(slice3, slice4)
}
