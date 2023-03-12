package service

import (
	"cek/auth"
	"cek/model/web"
)

type UserService interface {
	FindAll(filters *map[string]string) []web.UserResponse
	Registration(request *web.RegistrationRequest) web.UserResponse
	Create(request *web.UserCreateRequest) web.UserResponse
	Update(id *int, request *web.UserUpdateRequest) web.UserResponse
	Delete(id *int)
	Login(nip, password, userAgent, remoteAddress *string) web.TokenResponse
	ChangePassword(auth *auth.AccessDetails, id *int, request *web.UserChangePasswordRequest) web.UserResponse
	ResetPassword(auth *auth.AccessDetails, id *int) web.UserResponse
	RefreshToken(request *web.RefreshTokenCreateRequest, userAgent, remoteAddress *string) web.TokenResponse
}
