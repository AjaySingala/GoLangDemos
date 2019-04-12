package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Print(i)
		i++
	}

	fmt.Println()
	for j := 0; j < 10; j++ {
		fmt.Print(j)
	}

	fmt.Println()
	for x := 1; x <= 10; x++ {
		if x%2 == 0 {
			fmt.Println(x, "is even")
		} else {
			fmt.Println(x, "is odd")
		}

		if x%3 == 0 {
			fmt.Println(x, "is divisible by 3")
		}
		if x%4 == 0 {
			fmt.Println(x, "is divisible by 4")
		}
		if x%5 == 0 {
			fmt.Println(x, "is divisible by 5")
		}
	}
}
