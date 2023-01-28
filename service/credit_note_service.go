package service

import (
	"cek/auth"
	"cek/model/web"
)

type TestService interface {
	FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.TestResponse
}
