package main

import "fmt"

func main() {
	var status string

	fmt.Print("Enter a status (new, active, inactive, disabled): ")
	fmt.Scanf("%s", &status)
	fmt.Println("You entered", status)

	switch status {
	case "new":
		fmt.Println("New customer")
	case "active":
		fmt.Println("Active customer")
	case "inactive":
		fmt.Println("Inactive customer")
	case "disabled":
		fmt.Println("Disabled customer")
	default:
		fmt.Println("Invalid status entered.")

	}
}
