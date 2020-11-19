package template

type Aircon struct {
	Operation *Action                `json:"operation"`
	Modes     map[string]*AirconMode `json:"modes"`
}

type AirconMode struct {
	Temp           *Action `json:"temp,omitempty"`
	Humid          *Action `json:"humid,omitempty"`
	Fan            *Action `json:"fan,omitempty"`
	HorizontalVane *Action `json:"horizontal_vane,omitempty"`
	VerticalVane   *Action `json:"vertical_vane,omitempty"`
}
