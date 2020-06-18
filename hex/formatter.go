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

// Format - Return Format
func Format(hex [][]int) []string {
	var line []string
	for _, n := range hex {
		s := fmt.Sprintf("[")
		for j, m := range n {
			if j < len(n)-1 {
				s += fmt.Sprintf("0x%02X, ", m)
			} else {
				s += fmt.Sprintf("0x%02X", m)
			}
		}
		s += fmt.Sprintf("]\n")
		line = append(line, s)
	}
	return line
}
