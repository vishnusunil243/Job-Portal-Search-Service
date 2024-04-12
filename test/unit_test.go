package searchServiceTest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vishnusunil243/Job-Portal-Search-Service/internal/adapters"
	"github.com/vishnusunil243/Job-Portal-Search-Service/internal/helper/helperstruct"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

func TestUserAddReview(t *testing.T) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %s", err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("test_db")
	collection := db.Collection("userreview")
	var DB *gorm.DB

	repository := adapters.NewSearchAdapter(DB, db)

	tests := []struct {
		name          string
		req           helperstruct.ReviewHelper
		expectedError error
	}{
		{
			name: "Successful insertion",
			req: helperstruct.ReviewHelper{
				UserId:      "user1",
				CompanyId:   "company1",
				Rating:      5,
				Username:    "user1",
				Description: "Great company",
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repository.UserAddReview(tt.req)
			assert.Equal(t, tt.expectedError, err)

			if err == nil {
				var review bson.M
				err := collection.FindOne(context.Background(), bson.M{"userId": tt.req.UserId}).Decode(&review)
				assert.NoError(t, err)
				assert.Equal(t, tt.req.UserId, review["userId"])
				assert.Equal(t, tt.req.CompanyId, review["companyId"])
				assert.Equal(t, int32(tt.req.Rating), review["rating"])
				assert.Equal(t, tt.req.Username, review["username"])
				assert.Equal(t, tt.req.Description, review["description"])
				_, err = collection.DeleteOne(context.Background(), bson.M{"userId": tt.req.UserId})
				assert.NoError(t, err)
			}
		})
	}
}
