package food

import "github.com/guregu/null/v6"

type AddFood struct {
	Id          string    `json:"id" binding:"required" extensions:"x-order=1"`
	UserId      string    `json:"userId" binding:"required" extensions:"x-order=2"`
	OwnerId     string    `json:"ownerId" binding:"required" extensions:"x-order=3"`
	LocationId  string    `json:"locationId" binding:"required" extensions:"x-order=4"`
	Images      string    `json:"images" binding:"required" extensions:"x-order=5"`
	Name        string    `json:"name" binding:"required" extensions:"x-order=6"`
	Description string    `json:"description" binding:"required" extensions:"x-order=7"`
	Price       uint64    `json:"price" binding:"required" extensions:"x-order=8"`
	Review      string    `json:"review" binding:"required" extensions:"x-order=9"`
	CreatedAt   null.Time `json:"createdAt" binding:"required" extensions:"x-order=10"`
	UpdatedAt   null.Time `json:"updatedAt" binding:"required" extensions:"x-order=11"`
}

type Food struct {
	Id          string    `json:"id" binding:"required" extensions:"x-order=1"`
	Username    string    `json:"username" binding:"required" extensions:"x-order=2"`
	Ownername   string    `json:"ownername" binding:"required" extensions:"x-order=3"`
	Location    string    `json:"location" binding:"required" extensions:"x-order=4"`
	Images      string    `json:"images" binding:"required" extensions:"x-order=5"`
	Name        string    `json:"name" binding:"required" extensions:"x-order=6"`
	Description string    `json:"description" binding:"required" extensions:"x-order=7"`
	Price       uint64    `json:"price" binding:"required" extensions:"x-order=8"`
	Review      string    `json:"review" binding:"required" extensions:"x-order=9"`
	CreatedAt   null.Time `json:"createdAt" binding:"required" extensions:"x-order=10"`
	UpdatedAt   null.Time `json:"updatedAt" binding:"required" extensions:"x-order=11"`
}
