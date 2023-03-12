package service

import (
	"cek/auth"
	"cek/model/web"
)

type NoteService interface {
	FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.NoteResponse
	Create(auth *auth.AccessDetails, request *web.NoteCreateRequest) web.NoteResponse
	Update(auth *auth.AccessDetails, noteID *int, request *web.NoteUpdateRequest) web.NoteResponse
	Delete(auth *auth.AccessDetails, noteID *int)
}
