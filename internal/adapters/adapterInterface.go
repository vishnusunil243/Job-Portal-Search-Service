package adapters

import (
	"github.com/vishnusunil243/Job-Portal-Search-Service/entities"
	"github.com/vishnusunil243/Job-Portal-Search-Service/internal/helper/helperstruct"
	"go.mongodb.org/mongo-driver/bson"
)

type AdapterInterface interface {
	AddSearchHistory(entities.SearchHistory) error
	GetSearchHistory(userId string) (entities.SearchHistory, error)
	UpdateSearchHistory(entities.SearchHistory) error
	UserAddReview(req helperstruct.ReviewHelper) error
	GetReviewsByCompany(companyId string) ([]bson.M, error)
	UserDeleteReview(userId string, companyId string) error
	GetReviewCheck(userId, companyId string) (bool, error)
	GetAverageRatingOfCompany(companyId string) (float64, error)
}
