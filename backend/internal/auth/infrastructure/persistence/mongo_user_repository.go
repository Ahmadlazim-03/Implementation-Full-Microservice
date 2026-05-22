package persistence

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"backend/internal/auth/domain/entity"
	"backend/internal/auth/domain/valueobject"
	shareddomain "backend/internal/shared/domain"
)

type userDoc struct {
	ID           string    `bson:"_id"`
	Email        string    `bson:"email"`
	PasswordHash string    `bson:"password_hash"`
	Name         string    `bson:"name"`
	Role         string    `bson:"role"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}

type MongoUserRepository struct {
	coll *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) (*MongoUserRepository, error) {
	coll := db.Collection("users")

	// Unique index on email — enforces invariant at infra level too.
	_, err := coll.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return nil, err
	}
	return &MongoUserRepository{coll: coll}, nil
}

func (r *MongoUserRepository) Save(ctx context.Context, u *entity.User) error {
	doc := userDoc{
		ID:           u.ID(),
		Email:        u.Email().String(),
		PasswordHash: u.PasswordHash(),
		Name:         u.Name(),
		Role:         string(u.Role()),
		CreatedAt:    u.CreatedAt(),
		UpdatedAt:    u.UpdatedAt(),
	}
	_, err := r.coll.InsertOne(ctx, doc)
	return err
}

func (r *MongoUserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	var d userDoc
	err := r.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&d)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, shareddomain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return docToEntity(d)
}

func (r *MongoUserRepository) FindByEmail(ctx context.Context, email valueobject.Email) (*entity.User, error) {
	var d userDoc
	err := r.coll.FindOne(ctx, bson.M{"email": email.String()}).Decode(&d)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, shareddomain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return docToEntity(d)
}

func (r *MongoUserRepository) ExistsByEmail(ctx context.Context, email valueobject.Email) (bool, error) {
	count, err := r.coll.CountDocuments(ctx, bson.M{"email": email.String()})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func docToEntity(d userDoc) (*entity.User, error) {
	email, err := valueobject.NewEmail(d.Email)
	if err != nil {
		return nil, err
	}
	return entity.Hydrate(d.ID, email, d.PasswordHash, d.Name, entity.Role(d.Role), d.CreatedAt, d.UpdatedAt), nil
}
