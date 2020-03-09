package hex

import "fmt"

// Print - Show hex code
func Print(hex [][]int) {
	for _, n := range hex {
		fmt.Printf("[")
		for j, m := range n {
			if j < len(n)-1 {
				fmt.Printf("0x%02X, ", m)
			} else {
				fmt.Printf("0x%02X", m)
			}
		}
		fmt.Printf("]")
		fmt.Printf("\n")
	}
}
