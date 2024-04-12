package usecases

import "github.com/vishnusunil243/Job-Portal-Search-Service/entities"

type SearchUsecaseInterface interface {
	AddSearchHistory(entities.SearchHistory) error
	GetSearchHistory(string) ([]entities.SearchHistory, error)
}
