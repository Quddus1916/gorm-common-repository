package gorm_common_repository

// PageResponse is used to return paginated data
type PageResponse[T any] struct {
	// Data holds the original response list
	Data []T `json:"data"`
	// Total represents the total number of items in the DB
	Total int64 `json:"total"`
	// Sort represents the applied sorting
	Sort Sort `json:"sort,omitempty"`
	// Page represents the current page info
	Page Page `json:"page"`
	// FilterParams represents the applied query params for filtering
	FilterParams []FilterParam `json:"filter_params,omitempty"`
}
