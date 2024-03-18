package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/vishnusunil243/Job-Portal-Search-Service/entities"
	"github.com/vishnusunil243/Job-Portal-Search-Service/internal/usecases"
	"github.com/vishnusunil243/Job-Portal-proto-files/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SearchService struct {
	usecases *usecases.SearchUsecase
	pb.UnimplementedSearchServiceServer
}

func NewSearchService(usecases *usecases.SearchUsecase) *SearchService {
	return &SearchService{
		usecases: usecases,
	}
}
func (search *SearchService) AddSearchHistory(ctx context.Context, req *pb.SearchRequest) (*emptypb.Empty, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	reqEntity := entities.SearchHistory{
		UserId:  userId,
		Keyword: req.Keyword,
	}
	err = search.usecases.AddSearchHistory(reqEntity)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
