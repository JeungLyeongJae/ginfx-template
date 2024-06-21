package response

import "ginfx-template/internal/model"

type Page struct {
	PageNumber int           `json:"page_number"`
	PageSize   int           `json:"page_size"`
	TotalCount int64         `json:"total_count"`
	Condition  string        `json:"condition"`
	Users      []*model.User `json:"users"`
}
