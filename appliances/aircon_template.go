package appliances

type AirconTemplate struct {
	Operation *ActionTemplate                `json:"operation"`
	Options   map[string]*ActionTemplate     `json:"options,omitempty"`
	Modes     map[string]*AirconModeTemplate `json:"modes"`
}

type AirconModeTemplate struct {
	Temp           *ActionTemplate            `json:"temp,omitempty"`
	Humid          *ActionTemplate            `json:"humid,omitempty"`
	Fan            *ActionTemplate            `json:"fan,omitempty"`
	HorizontalVane *ActionTemplate            `json:"horizontal_vane,omitempty"`
	VerticalVane   *ActionTemplate            `json:"vertical_vane,omitempty"`
	Options        map[string]*ActionTemplate `json:"options,omitempty"`
}
