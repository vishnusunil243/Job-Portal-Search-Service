package initializer

import (
	"github.com/vishnusunil243/Job-Portal-Search-Service/internal/adapters"
	"github.com/vishnusunil243/Job-Portal-Search-Service/internal/service"
	"github.com/vishnusunil243/Job-Portal-Search-Service/internal/usecases"
	"gorm.io/gorm"
)

func Initializer(db *gorm.DB) *service.SearchService {
	adapter := adapters.NewSearchAdapter(db)
	usecase := usecases.NewSearchUsecase(adapter)
	service := service.NewSearchService(usecase)
	return service
}
