package repositories

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LinksRepo struct {
	collection *mongo.Collection
}

func (l *LinksRepo) Create(ctx context.Context, item LinkItem) error {
	_, err := l.collection.InsertOne(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func (l *LinksRepo) GetByID(ctx context.Context, id int64) (string, error) {
	var res LinkItem
	if err := l.collection.FindOne(ctx, bson.M{"id": id}).Decode(&res); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", nil
		}

		return "", err
	}

	return res.Link, nil
}

func NewLinksRepository(db *mongo.Database) *LinksRepo {
	return &LinksRepo{
		collection: db.Collection(linksCollection),
	}
}
