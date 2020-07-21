package fujitsu01

import (
	"errors"

	"github.com/dash-app/remote-go/aircon"
	"github.com/dash-app/remote-go/remote"
)

func (r *fujitsu01) Generate(e *aircon.Entry) ([]*remote.HexCode, error) {
	//運転内容変更
	Original_code := [][]int{
		{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x00},
	}
	//ON OFF 運転内容変更
	Action_code := [][]int{
		{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x00},
	}

	// Operation
	if e.Operation {
		// 運転ボタンが押されたときは +0x01 それ以外の操作はいらん
		// code[0][8] += 0x00
		//Action_code[0][8] += 0x01
	} else {
		// code[0][8] -= 0x01
		return []*remote.HexCode{
			{
				Code: [][]int{
					// オフのときの信号
					{0x14, 0x63, 0x00, 0x10, 0x10, 0x02, 0xFD},
				},
			},
		}, nil
	}

	// Mode
	switch e.Mode {
	case "auto":
		Original_code[0][9] = 0x00
		Action_code[0][9] = 0x00
	case "cool":
		Original_code[0][9] = 0x01
		Action_code[0][9] = 0x01
	case "dry":
		Original_code[0][9] = 0x05
		Action_code[0][9] = 0x05
	case "heat":
		Original_code[0][9] = 0x04
		Action_code[0][9] = 0x04
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
			return nil, errors.New("invalid temp_AUTO provided:1")
		}
		Original_code[0][8] = 0x80
		Action_code[0][8] = 0x81
		if temp >= 0.0 {
			Original_code[0][14] = 0x00 + (0x05 * (int((temp) * 2.0)))
			Action_code[0][14] = 0x00 + (0x05 * (int((temp) * 2.0)))
		} else {
			switch e.Temp {
			case -0.5:
				Original_code[0][14] = 0xFB
				Action_code[0][14] = 0xFB
			case -1.0:
				Original_code[0][14] = 0xF6
				Action_code[0][14] = 0xF6
			case -1.5:
				Original_code[0][14] = 0xF0
				Action_code[0][14] = 0xF0
			case -2.0:
				Original_code[0][14] = 0xEC
				Action_code[0][14] = 0xEC
			default:
				return nil, errors.New("invalid temp_AUTO provided:2")
			}
		}
	case "cool", "dry":
		var temp int
		if t, ok := e.Temp.(float64); ok {
			temp = int(t)
		} else if t, ok := e.Temp.(int); ok {
			temp = t
		} else {
			return nil, errors.New("invalid temp_COOL&DRY provided")
		}
		Original_code[0][14] = 0x00
		Action_code[0][14] = 0x00
		Original_code[0][8] = 0x50 + (0x04 * int((temp-18)*2.0))
		Action_code[0][8] = 0x50 + (0x04 * int((temp-18)*2.0)) + 0x01

	case "heat":
		var temp int
		if t, ok := e.Temp.(float64); ok {
			temp = int(t)
		} else if t, ok := e.Temp.(int); ok {
			temp = t
		} else {
			return nil, errors.New("invalid temp_HEAT provided")
		}
		Original_code[0][14] = 0x00
		Action_code[0][14] = 0x00
		Original_code[0][8] = 0x40 + (0x04 * int((temp-16)*2.0))
		Action_code[0][8] = 0x40 + (0x04 * int((temp-16)*2.0)) + 0x01
	default:
		return nil, errors.New("invalid temp provided")
	}

	//Fan
	switch e.Fan {
	case "auto":
		Original_code[0][10] += 0x00
		Action_code[0][10] += 0x00
	case "1":
		Original_code[0][10] += 0x01
		Action_code[0][10] += 0x01
	case "2":
		Original_code[0][10] += 0x03
		Action_code[0][10] += 0x03
	case "3":
		Original_code[0][10] += 0x06
		Action_code[0][10] += 0x06
	case "4":
		Original_code[0][10] += 0x08
		Action_code[0][10] += 0x08
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
		Original_code[0][10] += 0x10
		Action_code[0][10] += 0x10
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
		Original_code[0][10] += 0x20
		Action_code[0][10] += 0x20
	default:
		return nil, errors.New("invalid vertical_vane provided")
	}

	if e.Operation {
		O_sum := 0
		A_sum := 0
		for _, c := range Original_code[0] {
			O_sum += c
		}
		for _, c := range Action_code[0] {
			A_sum += c
		}
		Original_code[0][len(Original_code[0])-1] = (0xFA0 - O_sum%256) % 256
		Action_code[0][len(Action_code[0])-1] = (0xFA0 - A_sum%256) % 256
	}

	// Create of check_sum
	// choose 8 bit -> 0xNNN -> 0x_NN

	return []*remote.HexCode{
		{
			Code: Original_code, // 通常時の信号
		},
		{
			Code:     Action_code, // オンのときの信号 (運転開始信号)
			Interval: 1000,
		},
	}, nil
}
