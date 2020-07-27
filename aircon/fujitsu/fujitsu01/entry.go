package fujitsu01

import (
	"errors"

	"github.com/dash-app/remote-go/aircon"
	"github.com/dash-app/remote-go/remote"
)

func (r *fujitsu01) Generate(e *aircon.Entry) ([]*remote.HexCode, error) {
	code := [][]int{
		{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x00},
	}
	opcode := [][]int{
		{},
	}

	// Operation
	if !e.Operation {
		return []*remote.HexCode{
			{
				Code: [][]int{
					{0x14, 0x63, 0x00, 0x10, 0x10, 0x02, 0xFD},
				},
			},
		}, nil
	}

	// Mode
	switch e.Mode {
	case "auto":
		code[0][9] = 0x00
	case "cool":
		code[0][9] = 0x01
	case "dry":
		code[0][9] = 0x05
	case "heat":
		code[0][9] = 0x04
	default:
		return nil, errors.New("invalid mode provided")
	}

	// Temp
	switch e.Mode {
	case "auto":
		var temp int
		if t, ok := e.Temp.(float64); ok {
			temp = int(t)
		} else if t, ok := e.Temp.(int); ok {
			temp = t
		} else {
			return nil, errors.New("invalid temp_auto provided:1")
		}
		code[0][8] = 0x80
		if temp >= 0.0 {
			code[0][14] = 0x00 + (0x05 * (int((temp) * 2.0)))
		} else {
			switch e.Temp {
			case -0.5:
				code[0][14] = 0xFB
			case -1.0:
				code[0][14] = 0xF6
			case -1.5:
				code[0][14] = 0xF0
			case -2.0:
				code[0][14] = 0xEC
			default:
				return nil, errors.New("invalid temp_auto provided:2")
			}
		}
	case "cool", "dry":
		var temp int
		if t, ok := e.Temp.(float64); ok {
			temp = int(t)
		} else if t, ok := e.Temp.(int); ok {
			temp = t
		} else {
			return nil, errors.New("invalid temp_cool&dry provided")
		}
		code[0][14] = 0x00
		code[0][8] = 0x50 + (0x04 * int((temp-18)*2.0))

	case "heat":
		var temp int
		if t, ok := e.Temp.(float64); ok {
			temp = int(t)
		} else if t, ok := e.Temp.(int); ok {
			temp = t
		} else {
			return nil, errors.New("invalid temp_heat provided")
		}
		code[0][14] = 0x00
		code[0][8] = 0x40 + (0x04 * int((temp-16)*2.0))
	default:
		return nil, errors.New("invalid temp provided")
	}

	//Fan
	switch e.Fan {
	case "auto":
		code[0][10] += 0x00
	case "1":
		code[0][10] += 0x01
	case "2":
		code[0][10] += 0x03
	case "3":
		code[0][10] += 0x06
	case "4":
		code[0][10] += 0x08
	default:
		return nil, errors.New("invalid fan provaided")
	}

	//HorizontalVane
	switch e.HorizontalVane {
	case "switch":
		return []*remote.HexCode{
			{
				Code: [][]int{
					{0x14, 0x63, 0x00, 0x10, 0x10, 0x6C, 0x93},
				},
			},
		}, nil
	case "swing":
		code[0][10] += 0x10
	case "keep":
		//No processing
	default:
		return nil, errors.New("invalid horizontal_vane provided")
	}

	//VerticalVane
	switch e.VerticalVane {
	case "keep":
		//No processing
	case "switch":
		return []*remote.HexCode{
			{
				Code: [][]int{
					{0x14, 0x63, 0x00, 0x10, 0x10, 0x79, 0x86},
				},
			},
		}, nil
	case "swing":
		code[0][10] += 0x20
	default:
		return nil, errors.New("invalid vertical_vane provided")
	}

	// if e.Operation {
	// 	//code_copy = code
	// 	copy(code_copy, code[:])
	// 	code_copy[0][8] += 0x01
	// 	fmt.Printf("%x", code)
	// 	fmt.Printf("%x", code_copy)

	// 	sum := 0
	// 	for _, c := range code[0] {
	// 		sum += c
	// 	}
	// 	//print(sum)
	// 	code[0][len(code[0])-1] = (0xFA0 - sum%256) % 256
	// 	sum = 0
	// 	// print(sum)
	// 	for _, c := range code_copy[0] {
	// 		sum += c
	// 	}
	// 	//print(sum)
	// 	code_copy[0][len(code_copy[0])-1] = (0xFA0 - sum%256) % 256
	// }

	if e.Operation {
		opcode = make([][]int, len(code))
		for i, c := range code {
			opcode[i] = make([]int, len(c))
			copy(opcode[i], c)
		}
		opcode[0][8] += 0x01
		opcode[0][len(opcode[0])-1] = 0x00

		code[0][len(code[0])-1] = cheacksum(code)
		opcode[0][len(opcode[0])-1] = cheacksum(opcode)

	}

	// Create of check_sum
	// choose 8 bit -> 0xNNN -> 0x_NN
	return []*remote.HexCode{
		{
			Code: code, // 通常時の信号
		},
		{
			Code: opcode, // オンのときの信号 (運転開始信号)
			//Interval: 1000,
		},
	}, nil
}

func cheacksum(code [][]int) int {
	sum := 0
	for _, c := range code[0] {
		sum += c
	}
	results := (0xFA0 - sum%256) % 256

	return results
}
