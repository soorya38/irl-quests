package presenter

import "time"

type Quest struct {
	Id          int64    `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Level       int      `json:"level"`
	Tags        []string `json:"tags"`

	CreatedBy int64     `json:"createdby"`
	UpdatedBy int64     `json:"updatedby"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
