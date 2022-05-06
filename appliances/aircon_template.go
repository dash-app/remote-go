package appliances

type AirconTemplate struct {
	Operation *ActionTemplate                `json:"operation"`
	Modes     map[string]*AirconModeTemplate `json:"modes"`
}

type AirconModeTemplate struct {
	Temp           *ActionTemplate `json:"temp,omitempty"`
	Humid          *ActionTemplate `json:"humid,omitempty"`
	Fan            *ActionTemplate `json:"fan,omitempty"`
	HorizontalVane *ActionTemplate `json:"horizontal_vane,omitempty"`
	VerticalVane   *ActionTemplate `json:"vertical_vane,omitempty"`
}
