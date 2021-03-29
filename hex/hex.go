package hex

type HexCode struct {
	Format   CodeFormat
	Code     [][]int
	Raw      []int
	Interval int
}

type CodeFormat int32

const (
	// AEHA - Association for Electric Home Appliances
	AEHA CodeFormat = iota

	// WIP: NEC - NEC Format
	// NEC

	// WIP: SONY - Sony IR Format
	// SONY

	// RAW - Raw IR Format
	RAW
)
