package panasonic01

import (
	"github.com/dash-app/remote-go/hex"
	"github.com/dash-app/remote-go/light"
)

func (r *panasonic01) Generate(e *light.Entry) ([]*hex.HexCode, error) {
	magicNum = 0x24
	code := [][]int{
		{0x2C, 0x52, 0x09, 0x2F, 0x26},
		{0x2C, 0x52, 0x09, 0x2F, 0x26},
	}

	sum := 0
	for _, c := range code[1] {
		sum += c
	}
	code[1][len(code[1])-1] = sum & 0xFF
	return []*hex.HexCode{
		{
			Code: code,
		},
	}, nil
}
