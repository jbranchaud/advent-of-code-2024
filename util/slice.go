package util

import "fmt"

func PrintFirstAndLast(name string, list []int) {
	if len(list) > 0 {
		fmt.Printf("List '%s' - First: %d, Last: %d\n",
			name, list[0], list[len(list)-1])
	} else {
		fmt.Printf("List '%s' is empty", name)
	}
}
