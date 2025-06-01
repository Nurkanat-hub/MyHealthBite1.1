package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Payment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	OrderID   string             `bson:"order_id"`
	UserName  string             `bson:"user_name"`
	Amount    float64            `bson:"amount"`
	Status    string             `bson:"status"`
	CreatedAt time.Time          `bson:"created_at"`
}

type PaymentFilter struct {
	UserName string
	Status   string
	FromDate time.Time
	ToDate   time.Time
	Limit    int64
}

type Repository struct {
	Collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *Repository {
	return &Repository{Collection: collection}
}

func (r *Repository) Save(ctx context.Context, p *Payment) error {
	p.CreatedAt = time.Now()
	_, err := r.Collection.InsertOne(ctx, p)
	return err
}

func (r *Repository) GetHistory(ctx context.Context, filter PaymentFilter) ([]*Payment, error) {
	query := bson.M{"user_name": filter.UserName}

	if filter.Status != "" {
		query["status"] = filter.Status
	}

	if !filter.FromDate.IsZero() || !filter.ToDate.IsZero() {
		dateRange := bson.M{}
		if !filter.FromDate.IsZero() {
			dateRange["$gte"] = filter.FromDate
		}
		if !filter.ToDate.IsZero() {
			dateRange["$lte"] = filter.ToDate
		}
		query["created_at"] = dateRange
	}

	findOptions := options.Find()
	if filter.Limit > 0 {
		findOptions.SetLimit(filter.Limit)
	}

	cursor, err := r.Collection.Find(ctx, query, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*Payment
	for cursor.Next(ctx) {
		var p Payment
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		results = append(results, &p)
	}

	return results, nil
}
