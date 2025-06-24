package filters

type GroupFilter struct {
	GroupIDLike *string `in:"query=id_like" json:"id_like"`
}
