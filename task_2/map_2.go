// Input: paths = [["B","C"],["D","B"],["C","A"]]
// Output: "A"

// Explanation: All possible trips are:
// "D" -> "B" -> "C" -> "A".
// "B" -> "C" -> "A".
// "C" -> "A".
// "A".
// Clearly the destination city is "A".

package main

import "fmt"

func main() {

	Input := [][]string{
		{"D", "B"},
		{"B", "C"},
		{"C", "A"},
	}

	dict := make(map[string]int)

	for _, inner_object := range Input {
		for i, k := range inner_object {
			if i == 0 {
				dict[k] = 0
			} else {
				if _, ok := dict[k]; ok {
					delete(dict, k)
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
		}
		fmt.Println(result)
	}

}
