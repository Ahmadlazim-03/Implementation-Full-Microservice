package persistence

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"backend/internal/places/domain/entity"
	"backend/internal/places/domain/repository"
	shareddomain "backend/internal/shared/domain"
	"backend/internal/shared/valueobject"
)

type placeDoc struct {
	ID          string    `bson:"_id"`
	CategoryID  string    `bson:"category_id"`
	Name        string    `bson:"name"`
	Latitude    float64   `bson:"latitude"`
	Longitude   float64   `bson:"longitude"`
	Address     string    `bson:"address"`
	Description string    `bson:"description"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

type MongoPlaceRepository struct {
	coll *mongo.Collection
}

func NewMongoPlaceRepository(db *mongo.Database) *MongoPlaceRepository {
	return &MongoPlaceRepository{coll: db.Collection("places")}
}

func (r *MongoPlaceRepository) Save(ctx context.Context, p *entity.Place) error {
	doc := placeDoc{
		ID:          p.ID(),
		CategoryID:  p.CategoryID(),
		Name:        p.Name(),
		Latitude:    p.Location().Latitude(),
		Longitude:   p.Location().Longitude(),
		Address:     p.Address(),
		Description: p.Description(),
		CreatedAt:   p.CreatedAt(),
		UpdatedAt:   p.UpdatedAt(),
	}
	_, err := r.coll.InsertOne(ctx, doc)
	return err
}

func (r *MongoPlaceRepository) FindByID(ctx context.Context, id string) (*entity.Place, error) {
	var d placeDoc
	err := r.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&d)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, shareddomain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return toPlace(d)
}

func (r *MongoPlaceRepository) FindAll(ctx context.Context, filter repository.PlaceFilter) ([]*entity.Place, error) {
	q := bson.M{}
	if filter.CategoryID != "" {
		q["category_id"] = filter.CategoryID
	}
	if filter.Search != "" {
		q["name"] = bson.M{"$regex": filter.Search, "$options": "i"}
	}

	opts := options.Find().SetLimit(filter.Limit).SetSkip(filter.Skip).SetSort(bson.D{{Key: "created_at", Value: -1}})
	cur, err := r.coll.Find(ctx, q, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	out := []*entity.Place{}
	for cur.Next(ctx) {
		var d placeDoc
		if err := cur.Decode(&d); err != nil {
			return nil, err
		}
		p, err := toPlace(d)
		if err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, nil
}

func toPlace(d placeDoc) (*entity.Place, error) {
	coord, err := valueobject.NewCoordinate(d.Latitude, d.Longitude)
	if err != nil {
		return nil, err
	}
	return entity.Hydrate(d.ID, d.CategoryID, d.Name, coord, d.Address, d.Description, d.CreatedAt, d.UpdatedAt), nil
}
