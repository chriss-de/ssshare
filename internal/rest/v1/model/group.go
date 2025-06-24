package model

type GroupLinks struct {
	SelfLink
}

type Group struct {
	ID    string     `json:"id"`               // role ID
	Links GroupLinks `json:"_links,omitempty"` // REST links to other resources
}
