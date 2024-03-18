package adapters

import (
	"github.com/google/uuid"
	"github.com/vishnusunil243/Job-Portal-Search-Service/entities"
	"gorm.io/gorm"
)

type Adapter struct {
	DB *gorm.DB
}

func NewSearchAdapter(db *gorm.DB) *Adapter {
	return &Adapter{
		DB: db,
	}
}
func (search *Adapter) AddSearchHistory(req entities.SearchHistory) error {
	id := uuid.New()
	insertSearchHistory := `INSERT INTO search_histories (id,user_id,keyword) VALUES ($1,$2,$3)`
	if err := search.DB.Exec(insertSearchHistory, id, req.UserId, req.Keyword).Error; err != nil {
		return err
	}
	return nil
}
func (search *Adapter) GetSearchHistory(userId string) ([]entities.SearchHistory, error) {
	selectSearchHistory := `SELECT * FROM search_histories WHERE user_id=$1`
	var res []entities.SearchHistory
	if err := search.DB.Raw(selectSearchHistory, userId).Scan(&res).Error; err != nil {
		return []entities.SearchHistory{}, err
	}
	return res, nil
}
