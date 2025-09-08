package user

type UserRegisterReq struct {
	Username        string `json:"username" binding:"required" extensions:"x-order=1"`
	Password        string `json:"password" binding:"required" extensions:"x-order=2"`
	ConfirmPassword string `json:"confirmPassword" binding:"required" extensions:"x-order=3"`
	ProfilePicture  string `json:"profilePicture" binding:"required" extensions:"x-order=4"`
	Email           string `json:"email" binding:"required" extensions:"x-order=5"`
	Fullname        string `json:"fullname" binding:"required" extensions:"x-order=6"`
}

type UserLoginReq struct {
	Username string `json:"username" binding:"required" extensions:"x-order=1"`
	Password string `json:"password" binding:"required" extensions:"x-order=2"`
}
