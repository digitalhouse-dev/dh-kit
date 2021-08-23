package meta

import (
	"strconv"
)

var (
	// DefaultPageSize specifies the default page size
	DefaultPageSize = 25
	// MaxPageSize specifies the maximum page size
	MaxPageSize = 1000
	// PageVar specifies the query parameter name for page number
	PageVar = "page"
	// PageSizeVar specifies the query parameter name for page size
	PageSizeVar = "limit"
)

// Meta object
type Meta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	PageCount  int `json:"page_count"`
	TotalCount int `json:"total_count"`
}

// New creates a new Meta instance.
// The page parameter is 1-based and refers to the current page index/number.
// The perPage parameter refers to the number of items on each page.
// And the total parameter specifies the total number of data items.
// If total is less than 0, it means total is unknown.
func New(page, perPage, total int) *Meta {
	if perPage <= 0 {
		perPage = DefaultPageSize
	}
	if perPage > MaxPageSize {
		perPage = MaxPageSize
	}
	pageCount := -1
	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}
	if page < 1 {
		page = 1
	}

	return &Meta{
		Page:       page,
		PerPage:    perPage,
		TotalCount: total,
		PageCount:  pageCount,
	}
}

// NewFormMap creates a Meta object using the query parameters found in the given map.
// count stands for the total number of items. Use -1 if this is unknown.
func NewFormMap(req map[string]string, count int) *Meta {
	page := parseInt(req[PageVar], 1)
	perPage := parseInt(req[PageSizeVar], DefaultPageSize)
	return New(page, perPage, count)
}

// parseInt parses a string into an integer. If parsing is failed, defaultValue will be returned.
func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}

// Offset returns the OFFSET value that can be used in a SQL statement.
func (p *Meta) Offset() int {
	return (p.Page - 1) * p.PerPage
}

// Limit returns the LIMIT value that can be used in a SQL statement.
func (p *Meta) Limit() int {
	return p.PerPage
}
