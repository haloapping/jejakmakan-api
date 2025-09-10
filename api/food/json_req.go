package food

type AddReq struct {
	UserId      string `json:"userId" binding:"required" extensions:"x-order=1"`
	OwnerId     string `json:"ownerId" binding:"required" extensions:"x-order=2"`
	LocationId  string `json:"locationId" binding:"required" extensions:"x-order=3"`
	Images      string `json:"images" binding:"required" extensions:"x-order=4"`
	Name        string `json:"name" binding:"required" extensions:"x-order=5"`
	Description string `json:"description" binding:"required" extensions:"x-order=6"`
	Price       uint64 `json:"price" binding:"required" extensions:"x-order=7"`
	Review      string `json:"review" binding:"required" extensions:"x-order=8"`
}

type UpdateReq struct {
	UserId      string `json:"userId" binding:"required" extensions:"x-order=1"`
	OwnerId     string `json:"ownerId" binding:"required" extensions:"x-order=2"`
	LocationId  string `json:"locationId" binding:"required" extensions:"x-order=3"`
	Images      string `json:"images" binding:"required" extensions:"x-order=4"`
	Name        string `json:"name" binding:"required" extensions:"x-order=5"`
	Description string `json:"description" binding:"required" extensions:"x-order=6"`
	Price       uint64 `json:"price" binding:"required" extensions:"x-order=7"`
	Review      string `json:"review" binding:"required" extensions:"x-order=8"`
}
