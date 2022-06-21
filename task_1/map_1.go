package main

import "fmt"

func main() {

	Input := [][]string{
		{"London", "New York"},
		{"New York", "Lima"},
		{"Lima", "Sao Paulo"},
	}

	dict := make(map[string]int)

	// Input: paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
	// Output: "Sao Paulo"

	// Explanation: Starting at "London" city you will reach "Sao Paulo" city which is the destination city. Your trip consist of: "London" -> "New York" -> "Lima" -> "Sao Paulo".

	// a["London"] = "New York"
	// a["New York"] = "Lima"
	// a["Lima"] = "Sao Paulo"

	// fmt.Println(a[Lima])
	// fmt.Println(Input)

	for _, inner_object := range Input {
		for i, k := range inner_object {
			if i == 0 {
				dict[k] = 0
			} else {
				if _, ok := dict[k]; ok {
						delete(dict, k) // no use of delete
					} else {
					dict[k] = 1
				}
			}
		}
	}

	fmt.Println(dict)

	result := ""

	for key, value := range dict {
		if value == 1 {
			result = key
			break
		}
	}
	fmt.Println(result)
}
