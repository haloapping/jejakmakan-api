package location

import "github.com/guregu/null/v6"

type Location struct {
	Id         string    `json:"id" binding:"required" extensions:"x-order=1"`
	District   string    `json:"district" binding:"required" extensions:"x-order=2"`
	City       string    `json:"city" binding:"required" extensions:"x-order=3"`
	Province   string    `json:"province" binding:"required" extensions:"x-order=4"`
	PostalCode string    `json:"postalCode" binding:"required" extensions:"x-order=5"`
	Details    string    `json:"details" binding:"required" extensions:"x-order=6"`
	CreatedAt  null.Time `json:"createdAt" binding:"required" extensions:"x-order=7"`
	UpdatedAt  null.Time `json:"updatedAt" binding:"required" extensions:"x-order=8"`
}
