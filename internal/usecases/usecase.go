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
	res, err := redis.Get(req.UserId.String(), req.Keyword)
	if err != nil {
		return err
	}
	if res < 5 {
		num, err := redis.Add(req.UserId.String(), req.Keyword)
		if err != nil {
			return err
		}
		if num >= 5 {
			if err := search.adapters.AddSearchHistory(req); err != nil {
				return err
			}
		}
	}
	return nil
}
