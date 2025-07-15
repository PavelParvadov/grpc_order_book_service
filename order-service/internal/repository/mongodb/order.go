package mongodb

import (
	"context"
	"github.com/PavelParvadov/grpc_order_book_service/order-service/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (s *Storage) Collection() *mongo.Collection {
	return s.Db.Database("order-service-db").Collection("orders")
}

func (s *Storage) GetOrders(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	find, err := s.Collection().Find(ctx, bson.M{})
	defer find.Close(ctx)
	if err != nil {
		return []models.Order{}, err
	}
	for find.Next(ctx) {
		var order models.Order
		err := find.Decode(&order)
		if err != nil {
			return []models.Order{}, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (s *Storage) Save(ctx context.Context, order models.Order) (string, error) {
	one, err := s.Collection().InsertOne(ctx, order)
	if err != nil {
		return "", err
	}
	id := one.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil

}
