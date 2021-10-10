package repositories

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CounterRepo struct {
	collection *mongo.Collection
}

func (c *CounterRepo) Get(ctx context.Context, rangeItem *RangeItem) (int64, error) {
	findOptions := options.FindOne()
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})

	var res LinkItem
	if err := c.collection.FindOne(ctx, bson.D{
		{
			Key: "id", Value: bson.D{
				{
					Key: "$gte", Value: rangeItem.Start,
				},
				{
					Key: "$lt", Value: rangeItem.End,
				},
			},
		},
	}, findOptions).Decode(&res); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return 0, nil
		}

		return 0, err
	}

	return res.ID, nil
}

func NewCounterRepository(mongodb *mongo.Database) *CounterRepo {
	return &CounterRepo{
		collection: mongodb.Collection(linksCollection),
	}
}
