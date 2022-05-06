package appliances

// LightTemplate - Template of Light
type LightTemplate struct {
	Mode       *ActionTemplate `json:"mode"`
	Brightness *ActionTemplate `json:"brightness,omitempty"`
	Color      *ActionTemplate `json:"color,omitempty"`
}

func (l *LightTemplate) CommandToAction(cmd string) *ActionTemplate {
	if l.Mode.Validate(cmd) == nil {
		return l.Mode
	} else if l.Brightness.Validate(cmd) == nil {
		return l.Brightness
	} else if l.Color.Validate(cmd) == nil {
		return l.Color
	}
	return nil
}
