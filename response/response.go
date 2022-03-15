package response

// Response
type Response interface {
	StatusCode() int
	GetBody() ([]byte, error)
	GetHeaders() map[string]string
	Error() string
	GetData() interface{}
}

// MetaResponse object
type MetaResponse struct {
	Page       *int    `json:"page,omitempty"`
	PerPage    *int    `json:"per_page,omitempty"`
	PageCount  *int    `json:"page_count,omitempty"`
	TotalCount *int    `json:"total_count,omitempty"`
	NextCursor *string `json:"next_cursor,omitempty"`
}
