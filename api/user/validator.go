package user

import (
	"net/mail"
	"unicode"
)

func RegisterValidation(req UserRegisterReq) map[string][]string {
	validation := make(map[string][]string)

	if req.Username == "" {
		validation["username"] = append(validation["username"], "cannot empty")
	}

	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		validation["email"] = append(validation["emai"], "invalid email")
	}

	if req.Fullname == "" {
		validation["fullname"] = append(validation["fullname"], "cannot empty")
	}

	nLowercase := 0
	nUppercase := 0
	nDigit := 0
	nPunct := 0
	for _, r := range req.Password {
		if unicode.IsLower(r) {
			nLowercase++
		}
		if unicode.IsUpper(r) {
			nUppercase++
		}
		if unicode.IsDigit(r) {
			nDigit++
		}
		if unicode.IsPunct(r) {
			nPunct++
		}
	}

	if nLowercase < 1 {
		validation["password"] = append(validation["password"], "minimal 1 lowercase")
	}
	if nUppercase < 1 {
		validation["password"] = append(validation["password"], "minimal 1 uppercase")
	}
	if nDigit < 1 {
		validation["password"] = append(validation["password"], "minimal 1 digit")
	}
	if nPunct < 1 {
		validation["password"] = append(validation["password"], "minimal 1 punctuation")
	}

	if req.ConfirmPassword != req.Password {
		validation["confirmPassword"] = append(validation["confirmPassword"], "confirm password and password is not match")
	}

	return validation
}

func LoginValidation(req UserLoginReq) map[string][]string {
	validation := make(map[string][]string)

	if req.Username == "" {
		validation["username"] = append(validation["username"], "cannot empty")
	}
	if req.Password == "" {
		validation["password"] = append(validation["password"], "cannot empty")
	}

	return validation
}
