package contract

// QueryRequest represents a simple general query request with key-value maps and pagination
type QueryRequest struct {
	Search  map[string]string `json:"search"`
	Filters map[string]string `json:"filters"`
	Sort    map[string]string `json:"sort"`
	Page    int               `json:"page"`
	Limit   int               `json:"limit"`
}

// NewQueryRequest creates a new QueryRequest with defaults
func NewQueryRequest() *QueryRequest {
	return &QueryRequest{
		Search:  make(map[string]string),
		Filters: make(map[string]string),
		Sort:    make(map[string]string),
		Page:    1,
		Limit:   10,
	}
}

// SetPagination sets pagination parameters
func (q *QueryRequest) SetPagination(page, limit int) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	q.Page = page
	q.Limit = limit
}
