package contract

type PaginationResponse struct {
	TotalItems  int `json:"total_items" example:"100"`
	TotalPage   int `json:"total_page" example:"10"`
	CurrentPage int `json:"current_page" example:"1"`
	PageLimit   int `json:"page_limit" example:"10"`
}
