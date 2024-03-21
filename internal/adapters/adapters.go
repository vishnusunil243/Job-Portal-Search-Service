package adapters

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vishnusunil243/Job-Portal-Search-Service/entities"
	"github.com/vishnusunil243/Job-Portal-Search-Service/internal/helper/helperstruct"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

type Adapter struct {
	DB      *gorm.DB
	Mongodb *mongo.Database
}

func NewSearchAdapter(db *gorm.DB, mongodb *mongo.Database) *Adapter {
	return &Adapter{
		DB:      db,
		Mongodb: mongodb,
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
func (search *Adapter) GetSearchHistory(userId string) (entities.SearchHistory, error) {
	selectSearchHistory := `SELECT * FROM search_histories WHERE user_id=$1`
	var res entities.SearchHistory
	if err := search.DB.Raw(selectSearchHistory, userId).Scan(&res).Error; err != nil {
		return entities.SearchHistory{}, err
	}
	return res, nil
}
func (search *Adapter) UpdateSearchHistory(req entities.SearchHistory) error {
	updateQuery := `UPDATE search_histories SET keyword=$1 WHERE user_id=$2`
	if err := search.DB.Exec(updateQuery, req.Keyword, req.UserId).Error; err != nil {
		return err
	}
	return nil
}
func (review *Adapter) UserAddReview(req helperstruct.ReviewHelper) error {
	collection := review.Mongodb.Collection("userreview")
	if collection == nil {
		err := review.Mongodb.CreateCollection(context.Background(), "userreview")
		if err != nil {
			return err
		}
		collection = review.Mongodb.Collection("userreview")
	}
	reviewDoc := bson.M{
		"userId":      req.UserId,
		"companyId":   req.CompanyId,
		"rating":      req.Rating,
		"username":    req.Username,
		"description": req.Description,
		"timestamp":   time.Now(),
	}
	_, err := collection.InsertOne(context.Background(), reviewDoc)
	if err != nil {
		return err
	}
	return nil
}
func (review *Adapter) GetReviewsByCompany(companyId string) ([]bson.M, error) {
	collection := review.Mongodb.Collection("userreview")
	if collection == nil {
		return nil, fmt.Errorf("collection not found")
	}
	filter := bson.M{"companyId": companyId}
	options := options.Find().SetSort(bson.D{{"timestamp", -1}})
	cursor, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var reviews []bson.M
	for cursor.Next(context.Background()) {
		var review bson.M
		err = cursor.Decode(&review)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}
func (review *Adapter) UserDeleteReview(userId string, companyId string) error {
	collection := review.Mongodb.Collection("userreview")
	if collection == nil {
		return fmt.Errorf("collection is empty")
	}
	filter := bson.M{"userId": userId, "companyId": companyId}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
func (review *Adapter) GetReviewCheck(userId, companyId string) (bool, error) {
	collection := review.Mongodb.Collection("userreview")
	if collection == nil {
		return false, fmt.Errorf("collection is empty")
	}
	filter := bson.M{"userId": userId, "companyId": companyId}
	res := collection.FindOne(context.Background(), filter)

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, res.Err()
	}
	return true, nil
}

func (r *Adapter) GetAverageRatingOfCompany(companyId string) (float64, error) {
	collection := r.Mongodb.Collection("userreview")
	if collection == nil {
		return 0, fmt.Errorf("collection is empty")
	}

	pipeline := bson.A{
		bson.M{"$match": bson.M{"companyId": companyId}},
		bson.M{"$group": bson.M{"_id": "$companyId", "averageRating": bson.M{"$avg": "$rating"}}},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return 0, err
	}

	defer cursor.Close(context.Background())
	if !cursor.Next(context.Background()) {
		return 0, fmt.Errorf("no reviews found for company with ID: %s", companyId)
	}

	var result bson.M
	if err := cursor.Decode(&result); err != nil {
		return 0, err
	}

	averageRating, ok := result["averageRating"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid average rating format")
	}

	return averageRating, nil
}
