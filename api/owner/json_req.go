package owner

type AddReq struct {
	Images string `json:"images" binding:"required" extensions:"x-order=1"`
	Name   string `json:"name" binding:"required" extensions:"x-order=2"`
}

type UpdateReq struct {
	Images string `json:"images" extensions:"x-order=1"`
	Name   string `json:"name" extensions:"x-order=2"`
}
