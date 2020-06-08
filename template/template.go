package template

type Template struct {
	Vendor string  `json:"vendor"`
	Model  string  `json:"model"`
	Kind   string  `json:"kind"`
	Aircon *Aircon `json:"aircon,omitempty"`
}

type Aircon struct {
	Operation *Action                `json:"operation"`
	Modes     map[string]interface{} `json:"modes"`
}

type Action struct {
	Type    ActionType    `json:"type"`
	Default interface{}   `json:"default,omitempty"`
	List    []interface{} `json:"list"`
	Range   *Range        `json:"range"`
	Toggle  *Toggle       `json:"toggle"`
	Shot    *Shot         `json:"shot"`
}

type ActionType int32

const (
	// LIST - Multiple choice
	LIST ActionType = iota

	// RANGE - Numeric range
	RANGE

	// TOGGLE - Switch value for boolean
	TOGGLE

	// SHOT - Raise when pushed
	SHOT
)

// Range - Numeric range
type Range struct {
	Step float64 `json:"step"`
	From float64 `json:"from"`
	To   float64 `json:"to"`
}

// Toggle - Switch value for boolean
type Toggle struct {
	ON  interface{} `json:"on"`
	OFF interface{} `json:"off"`
}

// Shot - Raise when pushed
type Shot struct {
	Value interface{} `json:"value"`
}
