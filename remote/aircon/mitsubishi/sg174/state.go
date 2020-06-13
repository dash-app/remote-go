package sg174

type State struct {
	Operation bool                  `json:"operation"`
	Mode      string                `json:"mode"`
	Entry     map[string]*ModeEntry `json:"entry"`
}

type ModeEntry struct {
	// Temp - cool/heat: float / dry: string
	Temp           interface{} `json:"temp"`
	Fan            string      `json:"fan"`
	HorizontalVane string      `json:"horizontal_vane"`
	VerticalVane   string      `json:"vertical_vane"`
}

func (s *State) ToEntry() *Entry {
	return &Entry{
		Operation:      s.Operation,
		Mode:           s.Mode,
		Temp:           s.Entry[s.Mode].Temp,
		Fan:            s.Entry[s.Mode].Fan,
		HorizontalVane: s.Entry[s.Mode].HorizontalVane,
		VerticalVane:   s.Entry[s.Mode].VerticalVane,
	}
}
