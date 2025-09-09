package owner

import "github.com/guregu/null/v6"

type Owner struct {
	Id        string    `json:"id" binding:"required" extensions:"x-order=1"`
	Images    string    `json:"images" binding:"required" extensions:"x-order=2"`
	Name      string    `json:"name" binding:"required" extensions:"x-order=3"`
	CreatedAt null.Time `json:"createdAt" binding:"required" extensions:"x-order=4"`
	UpdatedAt null.Time `json:"updatedAt" binding:"required" extensions:"x-order=5"`
}
