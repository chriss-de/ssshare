package model

import "math"

type PagingRequest struct {
	First *int `in:"query=first,default=10"`
	Skip  *int `in:"query=skip,default=0"`
	//OrderDirection string `in:"query=orderDirection,default=ASC"`
	//OrderBy        string `in:"query=orderBy,default=name"`
}

type PagingResponse struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	PageCount  int `json:"pageCount"`
	TotalCount int `json:"totalCount"`
}

func NewPaging(total int, first *int, skip *int) *PagingResponse {
	var _first, _skip = 10, 0
	if first != nil {
		_first = *first
	}
	if skip != nil {
		_skip = *skip
	}
	p := &PagingResponse{
		Page:       1 + int(math.Floor(float64(_skip)/float64(_first))),
		PageSize:   _first,
		PageCount:  int(math.Ceil(float64(total) / float64(_first))),
		TotalCount: total,
	}
	return p
}
