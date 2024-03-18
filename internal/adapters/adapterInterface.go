package adapters

import "github.com/vishnusunil243/Job-Portal-Search-Service/entities"

type AdapterInterface interface {
	AddSearchHistory(entities.SearchHistory) error
}
