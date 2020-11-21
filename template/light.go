package template

type Light struct {
	Mode       *Action `json:"mode"`
	Brightness *Action `json:"brightness"`
	Color      *Action `json:"color"`
}
