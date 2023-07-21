package pagination

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type PagingRequestInterface interface {
	PagingToString() string
}

type PagingRequest struct {
	Search     string
	Offset     int
	Limit      int
	Sort       string
	SortAlias  string
	SortFields map[string]string
}

func NewPagingRequest(c *gin.Context, sortFields map[string]string) *PagingRequest {
	request := PagingRequest{}

	request.Search = ""
	if src, ok := c.GetQuery("search"); ok {
		request.Search = src
	}

	request.Sort = "id,asc"
	if srt, ok := c.GetQuery("sort"); ok {
		request.Sort = srt
		request.SortAlias = sortAlias(srt, sortFields)
	}

	request.Limit = 25
	if pageSize, ok := c.GetQuery("size"); ok {
		if sz, err := strconv.Atoi(pageSize); err == nil {
			if sz > 0 {
				request.Limit = sz
			}
		}
	}

	request.Offset = 0
	if pageNumber, ok := c.GetQuery("page"); ok {
		if pg, err := strconv.Atoi(pageNumber); err == nil {
			if pg < 1 {
				pg = 1
			}
			request.Offset = (pg - 1) * request.Limit
		}
	}

	return &request
}

func (request *PagingRequest) PagingToString() string {
	return fmt.Sprintf("%s-%d-%s-%d",
		slug.Make(request.Search),
		request.Offset,
		slug.Make(request.Sort),
		request.Limit,
	)
}

func sortAlias(sort string, sortFields map[string]string) string {
	field := "id"
	method := "ASC"

	srts := strings.Split(sort, ",")
	log.Println("sort request ", srts, sortFields)
	if len(srts) >= 1 {
		if f := sortFields[srts[0]]; f != "" {
			field = sortFields[f]
		}
		if len(srts) == 2 {
			if strings.ToLower(srts[1]) == "desc" {
				method = "DESC"
			}
		}
	}

	return fmt.Sprintf("%s %s", field, method)
}
