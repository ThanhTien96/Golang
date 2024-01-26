package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{}	`json:"paging,omitempty"`
	Filter interface{}	`json:"filter,omitempty"`
}
