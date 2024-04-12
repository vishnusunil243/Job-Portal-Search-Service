package initializer

import (
	"github.com/vishnusunil243/Job-Portal-Search-Service/internal/adapters"
	"github.com/vishnusunil243/Job-Portal-Search-Service/internal/service"
	"github.com/vishnusunil243/Job-Portal-Search-Service/internal/usecases"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func Initializer(db *gorm.DB, mongoDB *mongo.Database) *service.SearchService {
	adapter := adapters.NewSearchAdapter(db, mongoDB)
	usecase := usecases.NewSearchUsecase(adapter)
	service := service.NewSearchService(usecase, adapter, ":8082", ":8081")
	return service
}
