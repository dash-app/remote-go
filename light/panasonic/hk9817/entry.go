package hk9817

import (
	"errors"

	"github.com/dash-app/remote-go/appliances"
	"github.com/dash-app/remote-go/hex"
)

func (r *hk9817) Generate(req *appliances.Request) ([]*hex.HexCode, error) {
	e := req.Light()
	switch e.Action {
	case "OFF":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x2C, 0x52, 0x09, 0x2F, 0x26},
					{0x2C, 0x52, 0x09, 0x2F, 0x26},
					{0x2C, 0x52, 0x09, 0x2F, 0x26},
				},
			},
		}, nil
	case "ON":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x2C, 0x52, 0x09, 0x2C, 0x25},
					{0x2C, 0x52, 0x09, 0x2C, 0x25},
					{0x2C, 0x52, 0x09, 0x2C, 0x25},
				},
			},
		}, nil
	case "NIGHT_LIGHT":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x2C, 0x52, 0x09, 0x2E, 0x27},
					{0x2C, 0x52, 0x09, 0x2E, 0x27},
					{0x2C, 0x52, 0x09, 0x2E, 0x27},
				},
			},
		}, nil
	case "FAV":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x2C, 0x52, 0x09, 0x2D, 0x24},
					{0x2C, 0x52, 0x09, 0x2D, 0x24},
					{0x2C, 0x52, 0x09, 0x2D, 0x24},
				},
			},
		}, nil
	case "COOL":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x2C, 0x52, 0x39, 0x8A, 0xB3},
					{0x2C, 0x52, 0x39, 0x8A, 0xB3},
					{0x2C, 0x52, 0x39, 0x8A, 0xB3},
				},
			},
		}, nil
	case "WARM":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x2C, 0x52, 0x39, 0x8B, 0xB2},
					{0x2C, 0x52, 0x39, 0x8B, 0xB2},
					{0x2C, 0x52, 0x39, 0x8B, 0xB2},
				},
			},
		}, nil
	case "DIM":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x2C, 0x52, 0x39, 0x92, 0xAB},
					{0x2C, 0x52, 0x39, 0x92, 0xAB},
					{0x2C, 0x52, 0x39, 0x92, 0xAB},
				},
			},
		}, nil
	case "TO_BRIGHT":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x2C, 0x52, 0x09, 0x2A, 0x23},
					{0x2C, 0x52, 0x09, 0x2A, 0x23},
					{0x2C, 0x52, 0x09, 0x2A, 0x23},
				},
			},
		}, nil
	case "TO_DIM":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x2C, 0x52, 0x09, 0x2B, 0x22},
					{0x2C, 0x52, 0x09, 0x2B, 0x22},
					{0x2C, 0x52, 0x09, 0x2B, 0x22},
				},
			},
		}, nil
	case "TO_WARM":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x2C, 0x52, 0x39, 0x91, 0xA8},
					{0x2C, 0x52, 0x39, 0x91, 0xA8},
					{0x2C, 0x52, 0x39, 0x91, 0xA8},
				},
			},
		}, nil
	case "TO_COOL":
		return []*hex.HexCode{
			{
				Code: [][]int{
					{0x2C, 0x52, 0x39, 0x90, 0xA9},
					{0x2C, 0x52, 0x39, 0x90, 0xA9},
					{0x2C, 0x52, 0x39, 0x90, 0xA9},
				},
			},
		}, nil
	}

	return nil, errors.New("invalid command provided")
}
