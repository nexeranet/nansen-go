package nansen

import "time"

type DateBody struct {
	From *string `json:"from,omitempty"`
	To   *string `json:"to,omitempty"`
}

func NewDateBody(from, to *time.Time) *DateBody {
	var fromStr, toStr string
	if from != nil {
		fromStr = from.Format(time.RFC3339)
	}
	if to != nil {
		toStr = to.Format(time.RFC3339)
	}
	return &DateBody{
		From: &fromStr,
		To:   &toStr,
	}
}

type PaginationBody struct {
	Page    int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
}

const (
	SortingDirectionASC  = "ASC"
	SortingDirectionDESC = "DESC"
)

type SortOrderBody struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

type PaginationResponse struct {
	Page       int  `json:"page"`
	PerPage    int  `json:"per_page"`
	IsLastPage bool `json:"is_last_page"`
}

type TimeFrame string

const (
	TimeFrame5Min   TimeFrame = "5m"
	TimeFrame1Hour  TimeFrame = "1h"
	TimeFrame6Hour  TimeFrame = "6h"
	TimeFrame12Hour TimeFrame = "12h"
	TimeFrameDay    TimeFrame = "1d"
	TimeFrameWeek   TimeFrame = "7d"
)
