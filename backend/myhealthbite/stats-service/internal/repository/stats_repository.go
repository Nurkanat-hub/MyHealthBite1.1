package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Stats struct {
	UserID          string    `bson:"user_id"`
	TargetCalories  int32     `bson:"target_calories"`
	CurrentCalories int32     `bson:"current_calories"`
	TargetWaterML   int32     `bson:"target_water_ml"`
	CurrentWaterML  int32     `bson:"current_water_ml"`
	UpdatedAt       time.Time `bson:"updated_at"`
}

type StatsRepository struct {
	Collection *mongo.Collection
}

func NewStatsRepository(col *mongo.Collection) *StatsRepository {
	return &StatsRepository{Collection: col}
}

func (r *StatsRepository) InitStats(ctx context.Context, stats Stats) error {
	_, err := r.Collection.InsertOne(ctx, stats)
	return err
}

func (r *StatsRepository) GetStatsByUserID(ctx context.Context, userID string) (Stats, error) {
	var stats Stats
	err := r.Collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&stats)
	return stats, err
}

func (r *StatsRepository) UpdateStats(ctx context.Context, userID string, deltaCal, deltaWater int32) (Stats, error) {
	update := bson.M{
		"$inc": bson.M{
			"current_calories": deltaCal,
			"current_water_ml": deltaWater,
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}
	_, err := r.Collection.UpdateOne(ctx, bson.M{"user_id": userID}, update)
	if err != nil {
		return Stats{}, err
	}
	return r.GetStatsByUserID(ctx, userID)
}

func (r *StatsRepository) ResetDailyStats(ctx context.Context, userID string) (Stats, error) {
	update := bson.M{
		"$set": bson.M{
			"current_calories": 0,
			"current_water_ml": 0,
			"updated_at":       time.Now(),
		},
	}
	_, err := r.Collection.UpdateOne(ctx, bson.M{"user_id": userID}, update)
	if err != nil {
		return Stats{}, err
	}
	return r.GetStatsByUserID(ctx, userID)
}

func (r *StatsRepository) DeleteByUserID(ctx context.Context, userID string) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"user_id": userID})
	return err
}
