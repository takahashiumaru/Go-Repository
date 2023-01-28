package service

import (
	"cek/auth"
	"cek/model/web"
)

type UserService interface {
	Create(request *web.UserCreateRequest) web.UserResponse
	Delete(id *int)
	Update(id *int, request *web.UserUpdateRequest) web.UserResponse
	ChangePassword(auth *auth.AccessDetails, id *int, request *web.UserChangePasswordRequest) web.UserResponse
	ResetPassword(auth *auth.AccessDetails, id *int) web.UserResponse
	RefreshToken(request *web.RefreshTokenCreateRequest, userAgent, remoteAddress *string) web.TokenResponse
	Login(nip, password, userAgent, remoteAddress *string) web.TokenResponse
	FindAll(filters *map[string]string) []web.UserResponse
}
