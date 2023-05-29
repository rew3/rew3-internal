package service

import (
	"encoding/json"
	"math"
)

/*
This class provides pagination information to the query results
@field data data response of the query
@field limit number of entities returned
@field pageNumber pageNumber of the repsonse
@field pagesCount total pagesCount
@field totalCount total number of entites
@author rew3 on 2023/05/11
*/
type PaginatedResult struct {
	Data       []json.RawMessage `bson:"data"`
	Limit      int               `bson:"limit"`
	PageNumber int               `bson:"page_number"`
	PagesCount int               `bson:"pages_count"`
	TotalCount int               `bson:"total_count"`
}

func (pr *PaginatedResult) CountPages() *PaginatedResult {
	pageCount := 0
	if pr.Limit > 0 {
		pageCount = int(math.Ceil(float64(pr.TotalCount) / float64(pr.Limit)))
	}
	return &PaginatedResult{
		Data:       pr.Data,
		Limit:      pr.Limit,
		PageNumber: pr.PageNumber,
		PagesCount: pageCount,
		TotalCount: pr.TotalCount,
	}
}
