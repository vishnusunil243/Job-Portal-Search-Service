package db

import (
	"context"

	"github.com/vishnusunil243/Job-Portal-Search-Service/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(connectTo string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connectTo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&entities.SearchHistory{})
	return db, nil
}
func InitMongoDB(connectTo string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectTo))
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return client.Database("review"), nil
}
