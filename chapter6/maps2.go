package main

import "fmt"

func main() {
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

	found_by_element(elements, "Li")
	delete(elements, "Li")
	found_by_element(elements, "Li")
}

func found_by_element(elements map[string]string, key string) {
	if name, ok := elements[key]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("Not found")
	}
}
