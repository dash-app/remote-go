package template

type Light struct {
	Mode       *Action `json:"mode"`
	Brightness *Action `json:"brightness"`
	Color      *Action `json:"color"`
}

func (l *Light) CommandToAction(cmd string) *Action {
	if l.Mode.Validate(cmd) == nil {
		return l.Mode
	} else if l.Brightness.Validate(cmd) == nil {
		return l.Brightness
	} else if l.Color.Validate(cmd) == nil {
		return l.Color
	}

	return nil
}
