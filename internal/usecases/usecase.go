package usecases

import (
	"github.com/vishnusunil243/Job-Portal-Search-Service/entities"
	"github.com/vishnusunil243/Job-Portal-Search-Service/internal/adapters"
	"github.com/vishnusunil243/Job-Portal-Search-Service/internal/helper/redis"
)

type SearchUsecase struct {
	adapters adapters.AdapterInterface
}

func NewSearchUsecase(adapters adapters.AdapterInterface) *SearchUsecase {
	return &SearchUsecase{
		adapters: adapters,
	}
}
func (search *SearchUsecase) AddSearchHistory(req entities.SearchHistory) error {
	num, err := redis.Add(req.UserId.String(), req.Keyword)
	if err != nil {
		return err
	}
	searchHistory, err := search.adapters.GetSearchHistory(req.UserId.String())
	if err != nil {
		return err
	}
	if searchHistory.Keyword == "" {
		if num >= 5 {
			err := search.adapters.AddSearchHistory(req)
			if err != nil {
				return err
			}
		}
	} else {
		if num >= 5 {
			err := search.adapters.UpdateSearchHistory(req)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (search *SearchUsecase) GetSearchHistory(userId string) (entities.SearchHistory, error) {
	res, err := search.adapters.GetSearchHistory(userId)
	if err != nil {
		return entities.SearchHistory{}, err
	}
	return res, nil
}
