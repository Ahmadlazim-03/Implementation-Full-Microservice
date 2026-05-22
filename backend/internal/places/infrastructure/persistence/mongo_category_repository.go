package persistence

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"backend/internal/places/domain/entity"
	shareddomain "backend/internal/shared/domain"
)

type categoryDoc struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
	Icon string `bson:"icon"`
}

type MongoCategoryRepository struct {
	coll *mongo.Collection
}

func NewMongoCategoryRepository(db *mongo.Database) *MongoCategoryRepository {
	return &MongoCategoryRepository{coll: db.Collection("categories")}
}

func (r *MongoCategoryRepository) Save(ctx context.Context, c *entity.Category) error {
	_, err := r.coll.InsertOne(ctx, categoryDoc{ID: c.ID(), Name: c.Name(), Icon: c.Icon()})
	return err
}

func (r *MongoCategoryRepository) FindByID(ctx context.Context, id string) (*entity.Category, error) {
	var d categoryDoc
	err := r.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&d)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, shareddomain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return entity.NewCategory(d.ID, d.Name, d.Icon), nil
}

func (r *MongoCategoryRepository) FindAll(ctx context.Context) ([]*entity.Category, error) {
	cur, err := r.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	out := []*entity.Category{}
	for cur.Next(ctx) {
		var d categoryDoc
		if err := cur.Decode(&d); err != nil {
			return nil, err
		}
		out = append(out, entity.NewCategory(d.ID, d.Name, d.Icon))
	}
	return out, nil
}
