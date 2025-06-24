package model

// swagger:model
type Href struct {
	Href string `json:"href,omitempty"`
}

// swagger:model
type SelfLink struct {
	Self Href `json:"self"`
}
