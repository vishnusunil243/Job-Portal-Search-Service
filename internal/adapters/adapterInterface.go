package adapters

import "github.com/vishnusunil243/Job-Portal-Search-Service/entities"

type AdapterInterface interface {
	AddSearchHistory(entities.SearchHistory) error
	GetSearchHistory(userId string) (entities.SearchHistory, error)
	UpdateSearchHistory(entities.SearchHistory) error
}
