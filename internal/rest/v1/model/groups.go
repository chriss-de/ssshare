package model

// Groups structure of a list of Groups
type Groups struct {
	Groups []Group         `json:"groups"`
	Links  *GroupsLinks    `json:"_links,omitempty"`
	Paging *PagingResponse `json:"_paging,omitempty"`
}

type GroupsLinks struct {
	SelfLink
}
