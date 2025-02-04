package payload

type PaginationResponse struct {
	Data      interface{} `json:"data"`
	Page      uint        `json:"page"`
	Limit     uint        `json:"limit"`
	TotalData uint64      `json:"total_data"`
	LastPage  uint        `json:"last_page"`
}
