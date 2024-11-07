package repository

import (
	"context"
	"log"
	"product_manage/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	if db == nil {
		log.Fatal("Database connection is nil")
	}
	return &UserRepository{collection: db.Collection("users")}
}

func (repo *UserRepository) CreateUser(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := repo.collection.InsertOne(ctx, user)
	if err != nil {
		log.Println("Có lỗi xảy ra , kết quả có lỗi !!!")
	}
	return err
}

func (repo *UserRepository) FindByUsername(username string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user model.User
	err := repo.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return &user, err
}
