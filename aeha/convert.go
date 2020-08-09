package aeha

// SignalToCode - Hex code to IR Code
func SignalToCode(T int, signal [][]int, interval int) []int {
	var code []int
	for i := range signal {
		c := signal[i]
		code = append(code, T*8, T*4)

		for j := range c {
			for k := 0; k < 8; k++ {
				code = append(code, T)
				if c[j]&(1<<k) != 0 {
					code = append(code, T*3)
				} else {
					code = append(code, T)
				}
			}
		}

		code = append(code, T)
		if i < len(signal)-1 {
			code = append(code, interval)
		}
	}

	return code
}
