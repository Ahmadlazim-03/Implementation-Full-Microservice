package persistence

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"backend/internal/review/domain/entity"
)

type reviewDoc struct {
	ID        string    `bson:"_id"`
	PlaceID   string    `bson:"place_id"`
	UserID    string    `bson:"user_id"`
	Rating    int       `bson:"rating"`
	Comment   string    `bson:"comment"`
	CreatedAt time.Time `bson:"created_at"`
}

type MongoReviewRepository struct {
	coll *mongo.Collection
}

func NewMongoReviewRepository(db *mongo.Database) *MongoReviewRepository {
	return &MongoReviewRepository{coll: db.Collection("reviews")}
}

func (r *MongoReviewRepository) Save(ctx context.Context, rv *entity.Review) error {
	_, err := r.coll.InsertOne(ctx, reviewDoc{
		ID: rv.ID(), PlaceID: rv.PlaceID(), UserID: rv.UserID(),
		Rating: rv.Rating(), Comment: rv.Comment(), CreatedAt: rv.CreatedAt(),
	})
	return err
}

func (r *MongoReviewRepository) FindByPlace(ctx context.Context, placeID string) ([]*entity.Review, error) {
	cur, err := r.coll.Find(ctx, bson.M{"place_id": placeID})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	out := []*entity.Review{}
	for cur.Next(ctx) {
		var d reviewDoc
		if err := cur.Decode(&d); err != nil {
			return nil, err
		}
		out = append(out, entity.Hydrate(d.ID, d.PlaceID, d.UserID, d.Rating, d.Comment, d.CreatedAt))
	}
	return out, nil
}
