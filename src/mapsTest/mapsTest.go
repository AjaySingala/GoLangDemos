package main

import "fmt"

func main() {
	// Maps aka hashtable or dictionary
	// An unordered collection of key-value pairs.
	// Syntax map[key_type]value_type

	// The following gives a runtime error.
	// var x map[string]int // x is a map of strings to ints

	// maps have to be initialized before use. So:
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)
	fmt.Println(x["key"])

	y := make(map[int]int)
	y[1] = 10
	y[2] = 20
	y[3] = 30
	fmt.Println(y)
	fmt.Println(y[1])
	fmt.Println(y[2])
	fmt.Println(y[3])

	// Delete items from a map
	fmt.Println("Deleting from map...")
	delete(y, 1)
	fmt.Println(y)
	fmt.Println(y[1])
	fmt.Println(y[2])
	fmt.Println(y[3])

	chemicalsMap()
	chemicalsMap2()
	chemicalMapWithState()
}

func chemicalMapWithState() {
	fmt.Println("Nested maps")
	elements := map[string]map[string]string{ // Note the type!!!
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	fmt.Println(elements)
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}

func chemicalsMap2() {
	// Shorter way of creating a map, just like an array.
	fmt.Println("Shorter way of map")
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements)
}

func chemicalsMap() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements)
	fmt.Println(elements["Li"])
	fmt.Println(elements["Un"]) // returns nothing as "Un" is not a key.

	// first value gives us the result. If key does not exist, returns nothing
	// second value indicates of the lookup was successful or not
	name, ok := elements["Ne"]
	fmt.Println(name, ok) // Prints "Neon true" since "Ne" exists.

	name1, ok1 := elements["Un"]
	fmt.Println(name1, ok1) // prints " false" as "Un" does not exist.

	if name3, ok3 := elements["Ne"]; ok3 {
		fmt.Println("Found", name3, ok3)
	} else {
		fmt.Println("'Un' not found in the map")
	}

	if name3, ok3 := elements["Un"]; ok3 {
		fmt.Println("Found", name3, ok3)
	} else {
		fmt.Println("'Un' not found in the map")
	}
}
