package appliances

// Request - payload of remote-go controller
type Request struct {
	aircon *Aircon
	light  *Light
}

// FromAircon - Make request from AC
func FromAircon(e *Aircon) *Request {
	return &Request{aircon: e}
}

// Aircon - Get as Aircon
func (req *Request) Aircon() *Aircon {
	return req.aircon
}

// FromLight - Make request from light
func FromLight(e *Light) *Request {
	return &Request{light: e}
}

// Light - Get as Light
func (req *Request) Light() *Light {
	return req.light
}
