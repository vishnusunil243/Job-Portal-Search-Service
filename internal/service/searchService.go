package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vishnusunil243/Job-Portal-Search-Service/entities"
	"github.com/vishnusunil243/Job-Portal-Search-Service/internal/adapters"
	"github.com/vishnusunil243/Job-Portal-Search-Service/internal/helper/helperstruct"
	"github.com/vishnusunil243/Job-Portal-Search-Service/internal/usecases"
	"github.com/vishnusunil243/Job-Portal-proto-files/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	UserConn pb.UserServiceClient
)

type SearchService struct {
	usecases *usecases.SearchUsecase
	adapters adapters.AdapterInterface
	pb.UnimplementedSearchServiceServer
}

func NewSearchService(usecases *usecases.SearchUsecase, adapters adapters.AdapterInterface) *SearchService {
	return &SearchService{
		usecases: usecases,
		adapters: adapters,
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
func (search *SearchService) GetSearchHistory(ctx context.Context, req *pb.UserId) (*pb.SearchResponse, error) {
	searchHistory, err := search.usecases.GetSearchHistory(req.UserId)
	if err != nil {
		return &pb.SearchResponse{}, err
	}
	res := &pb.SearchResponse{
		Designation: searchHistory.Keyword,
	}

	return res, nil
}
func (review *SearchService) UserAddReview(ctx context.Context, req *pb.UserReviewRequest) (*emptypb.Empty, error) {
	user, err := UserConn.GetUser(context.Background(), &pb.GetUserById{
		Id: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	check, err := review.adapters.GetReviewCheck(req.UserId, req.CompanyId)
	if err != nil {
		return nil, err
	}
	if check {
		return nil, fmt.Errorf("this user has already entered a review for the given company please update the existing one or add a new one")
	}
	reqEntity := helperstruct.ReviewHelper{
		UserId:      req.UserId,
		CompanyId:   req.CompanyId,
		Rating:      int(req.Rating),
		Username:    user.Name,
		Description: req.Description,
	}
	if err := review.adapters.UserAddReview(reqEntity); err != nil {
		return nil, err
	}
	return nil, nil
}
func (r *SearchService) GetCompanyReview(req *pb.ReviewByCompanyId, srv pb.SearchService_GetCompanyReviewServer) error {
	reviews, err := r.adapters.GetReviewsByCompany(req.CompanyId)
	if err != nil {
		return err
	}
	for _, review := range reviews {
		userId, ok := review["userId"]
		if !ok {
			return fmt.Errorf("userId field not present")
		}
		userName, ok := review["username"]
		if !ok {
			return fmt.Errorf("username not present")
		}
		companyId, ok := review["companyId"]
		if !ok {
			return fmt.Errorf("username not present")
		}
		rating, ok := review["rating"]
		if !ok {
			return fmt.Errorf("rating not present")
		}
		des, ok := review["description"]
		if !ok {
			return fmt.Errorf("desription not found")
		}
		res := &pb.ReviewResponse{
			UserId:      userId.(string),
			CompanyId:   companyId.(string),
			Description: des.(string),
			Username:    userName.(string),
			Rating:      rating.(int32),
		}
		if err := srv.Send(res); err != nil {
			return err
		}
	}
	return nil
}
func (review *SearchService) RemoveReview(ctx context.Context, req *pb.UserReviewRequest) (*emptypb.Empty, error) {
	if err := review.adapters.UserDeleteReview(req.UserId, req.CompanyId); err != nil {
		return nil, err
	}
	return nil, nil
}
