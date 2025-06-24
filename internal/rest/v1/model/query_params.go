package model

import "github.com/chriss-de/ssshare/internal/rest/v1/filters"

type GroupsQueryParams struct {
	*PagingRequest
	*filters.GroupFilter
}
