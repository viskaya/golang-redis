package pagination

import (
	"math"
)

type PagingResponse struct {
	Request *PagingRequest
	Rows    []interface{}
	MaxRow  int64
}

type PagingResponseJSON struct {
	Search     string        `json:"search"`
	PageNumber int64         `json:"pageNumber"`
	PageSize   int64         `json:"pageSize"`
	Sort       string        `json:"sort"`
	MaxPage    int64         `json:"maxPage"`
	MaxRow     int64         `json:"maxRow"`
	Rows       []interface{} `json:"rows"`
}

func NewPagingResponse(request *PagingRequest, rows []interface{}, maxRow int64) *PagingResponse {
	return &PagingResponse{
		Request: request,
		Rows:    rows,
		MaxRow:  maxRow,
	}
}

func (response *PagingResponse) GetPagingResponse() *PagingResponseJSON {
	if response.Request.Limit < 1 {
		response.Request.Limit = 25
	}

	size := int64(response.Request.Limit)
	if response.MaxRow > 0 && response.MaxRow < size {
		size = response.MaxRow
	}

	pageNumber := int64((response.Request.Offset / response.Request.Limit) + 1)

	maxPage := math.Ceil(float64(response.MaxRow) / float64(response.Request.Limit))
	if maxPage < 1 {
		maxPage = 1
	}

	if maxPage < float64(pageNumber) {
		pageNumber = int64(maxPage)
	}

	if response.Rows == nil {
		response.Rows = make([]interface{}, 0)
	}

	return &PagingResponseJSON{
		Search:     response.Request.Search,
		PageNumber: pageNumber,
		PageSize:   size,
		Sort:       response.Request.Sort,
		MaxPage:    int64(maxPage),
		MaxRow:     response.MaxRow,
		Rows:       response.Rows,
	}
}
