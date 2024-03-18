package usecases

import "github.com/vishnusunil243/Job-Portal-Search-Service/entities"

type ServiceUsecase interface {
	AddSearchHistory(entities.SearchHistory) error
}
