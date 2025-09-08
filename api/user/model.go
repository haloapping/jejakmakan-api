package user

import (
	"github.com/guregu/null/v6"
)

type UserRegister struct {
	Id             string    `json:"id" binding:"required" extensions:"x-order=1"`
	ProfilePicture string    `json:"profilePicture" binding:"required" extensions:"x-order=2"`
	Username       string    `json:"username" binding:"required" extensions:"x-order=3"`
	Email          string    `json:"email" binding:"required" extensions:"x-order=4"`
	Fullname       string    `json:"fullname" binding:"required" extensions:"x-order=5"`
	CreatedAt      null.Time `json:"createdAt" binding:"required" extensions:"x-order=6"`
	UpdatedAt      null.Time `json:"updatedAt" binding:"required" extensions:"x-order=7"`
}

type UserLogin struct {
	Id       string `json:"id" binding:"required" extensions:"x-order=1"`
	Username string `json:"username" binding:"required" extensions:"x-order=2"`
	Password string `json:"password" binding:"required" extensions:"x-order=3"`
}

type UserBiodata struct {
	Id             string `json:"id" binding:"required" extensions:"x-order=1"`
	ProfilePicture string `json:"profilePicture" binding:"required" extensions:"x-order=2"`
	Username       string `json:"username" binding:"required" extensions:"x-order=3"`
	Password       string `json:"password" binding:"required" extensions:"x-order=4"`
	Email          string `json:"email" binding:"required" extensions:"x-order=5"`
	Fullname       string `json:"fullname" binding:"required" extensions:"x-order=6"`
}
