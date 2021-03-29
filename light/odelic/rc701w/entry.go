package rc701w

import (
	"errors"

	"github.com/dash-app/remote-go/hex"
	"github.com/dash-app/remote-go/light"
)

func (r *rc701w) Generate(e *light.Entry) ([]*hex.HexCode, error) {
	switch e.Action {
	case "OFF":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x80, 0xC5, 0x08, 0xF7},
					{0x84, 0x51, 0x08, 0xF7},
				},
			},
		}, nil

	case "ON":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x80, 0xC5, 0x9A, 0x65},
					{0x84, 0x51, 0x9A, 0x65},
				},
			},
		}, nil
	case "NIGHT_LIGHT":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x80, 0xC5, 0x2A, 0xD5},
					{0x84, 0x51, 0x2A, 0xD5},
				},
			},
		}, nil

	case "FAV":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x80, 0xC5, 0x2E, 0xD1},
					{0x84, 0x51, 0x2E, 0xD1},
				},
			},
		}, nil
	case "TO_BRIGHT":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x80, 0xC5, 0x9E, 0x61},
					{0x84, 0x51, 0x9E, 0x61},
				},
			},
		}, nil
	case "TO_DIM":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x80, 0xC5, 0x96, 0x69},
					{0x84, 0x51, 0x96, 0x69},
				},
			},
		}, nil
	}
	return nil, errors.New("invalid command provided")
}
