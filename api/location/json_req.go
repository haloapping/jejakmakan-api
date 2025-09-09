package location

type AddReq struct {
	District   string `json:"district" binding:"required" extensions:"x-order=1"`
	City       string `json:"city" binding:"required" extensions:"x-order=2"`
	Province   string `json:"province" binding:"required" extensions:"x-order=3"`
	PostalCode string `json:"postalCode" binding:"required" extensions:"x-order=4"`
	Details    string `json:"details" binding:"required" extensions:"x-order=5"`
}

type UpdateReq struct {
	District   string `json:"district" extensions:"x-order=1"`
	City       string `json:"city" extensions:"x-order=2"`
	Province   string `json:"province" extensions:"x-order=3"`
	PostalCode string `json:"postalCode" extensions:"x-order=4"`
	Details    string `json:"details" extensions:"x-order=5"`
}
