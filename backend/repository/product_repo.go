package repository

import (
	"context"
	"product_manage/config"
	"product_manage/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{
		collection: config.DB.Collection("products"),
	}
}

func (repo *ProductRepository) CreateProduct(product *model.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := repo.collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	return nil
}
func (repo *ProductRepository) UpdateProduct(product *model.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	productID, errID := primitive.ObjectIDFromHex(product.ID)
	if errID != nil {
		return errID
	}
	filter := bson.M{"_id": productID}

	update := bson.M{
		"$set": bson.M{
			"name":        product.Name,
			"price":       product.Price,
			"description": product.Description,
		},
	}
	result, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNilCursor
	}
	return nil
}

func (repo *ProductRepository) DeleteProduct(productID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	productIDCheck, errID := primitive.ObjectIDFromHex(productID)
	if errID != nil {
		return errID
	}
	filter := bson.M{"_id": productIDCheck}

	_, err := repo.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepository) GetAllProducts() ([]model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var products []model.Product
	cursor.All(ctx, &products)
	return products, nil
}

func (repo *ProductRepository) GetProductById(productID string) (model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var product model.Product
	productIDCheck, errID := primitive.ObjectIDFromHex(productID)
	if errID != nil {
		return product, errID
	}
	filter := bson.M{"_id": productIDCheck}
	err := repo.collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return product, err
	}

	return product, nil
}
