package entities

import "github.com/google/uuid"

type SearchHistory struct {
	Id      uuid.UUID
	UserId  uuid.UUID
	Keyword string
}
