package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"backend/internal/places/domain/entity"
	"backend/internal/places/domain/repository"
	"backend/internal/shared/valueobject"
)

// CachedPlaceRepository is a DECORATOR around the real PlaceRepository.
// Reads check Redis first; writes invalidate the cache.
// The domain still sees a plain PlaceRepository — caching is invisible.
type CachedPlaceRepository struct {
	inner repository.PlaceRepository
	rdb   *redis.Client
	ttl   time.Duration
}

func NewCachedPlaceRepository(inner repository.PlaceRepository, rdb *redis.Client, ttl time.Duration) *CachedPlaceRepository {
	if ttl <= 0 {
		ttl = 5 * time.Minute
	}
	return &CachedPlaceRepository{inner: inner, rdb: rdb, ttl: ttl}
}

// ---- serializable view used only inside the cache layer ----
type cachedPlace struct {
	ID, CategoryID, Name, Address, Description string
	Lat, Lng                                   float64
	CreatedAt, UpdatedAt                       time.Time
}

func toCached(p *entity.Place) cachedPlace {
	return cachedPlace{
		ID: p.ID(), CategoryID: p.CategoryID(), Name: p.Name(),
		Address: p.Address(), Description: p.Description(),
		Lat: p.Location().Latitude(), Lng: p.Location().Longitude(),
		CreatedAt: p.CreatedAt(), UpdatedAt: p.UpdatedAt(),
	}
}

func fromCached(c cachedPlace) (*entity.Place, error) {
	coord, err := valueobject.NewCoordinate(c.Lat, c.Lng)
	if err != nil {
		return nil, err
	}
	return entity.Hydrate(c.ID, c.CategoryID, c.Name, coord, c.Address, c.Description, c.CreatedAt, c.UpdatedAt), nil
}

// ---- helpers ----
func listKey(f repository.PlaceFilter) string {
	return fmt.Sprintf("places:list:%s:%s:%d:%d", f.CategoryID, f.Search, f.Limit, f.Skip)
}

func (r *CachedPlaceRepository) invalidateLists(ctx context.Context) {
	// Best-effort: SCAN + DEL all list keys. Cheap because keys are short.
	iter := r.rdb.Scan(ctx, 0, "places:list:*", 0).Iterator()
	for iter.Next(ctx) {
		_ = r.rdb.Del(ctx, iter.Val()).Err()
	}
}

// ---- PlaceRepository implementation ----

func (r *CachedPlaceRepository) Save(ctx context.Context, p *entity.Place) error {
	if err := r.inner.Save(ctx, p); err != nil {
		return err
	}
	r.invalidateLists(ctx)
	return nil
}

func (r *CachedPlaceRepository) FindByID(ctx context.Context, id string) (*entity.Place, error) {
	key := "places:id:" + id
	if raw, err := r.rdb.Get(ctx, key).Bytes(); err == nil {
		var c cachedPlace
		if json.Unmarshal(raw, &c) == nil {
			if p, err := fromCached(c); err == nil {
				return p, nil
			}
		}
	}
	p, err := r.inner.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if data, err := json.Marshal(toCached(p)); err == nil {
		_ = r.rdb.Set(ctx, key, data, r.ttl).Err()
	}
	return p, nil
}

func (r *CachedPlaceRepository) FindAll(ctx context.Context, filter repository.PlaceFilter) ([]*entity.Place, error) {
	key := listKey(filter)
	if raw, err := r.rdb.Get(ctx, key).Bytes(); err == nil {
		var list []cachedPlace
		if json.Unmarshal(raw, &list) == nil {
			out := make([]*entity.Place, 0, len(list))
			for _, c := range list {
				if p, err := fromCached(c); err == nil {
					out = append(out, p)
				}
			}
			return out, nil
		}
	}

	items, err := r.inner.FindAll(ctx, filter)
	if err != nil {
		return nil, err
	}
	cached := make([]cachedPlace, 0, len(items))
	for _, p := range items {
		cached = append(cached, toCached(p))
	}
	if data, err := json.Marshal(cached); err == nil {
		_ = r.rdb.Set(ctx, key, data, r.ttl).Err()
	}
	return items, nil
}

// FindNearby is intentionally NOT cached — query parameters are continuous
// (latitude/longitude floats) so cache hit ratio would be ~0 and would
// just waste memory. Calls go straight to Mongo.
func (r *CachedPlaceRepository) FindNearby(ctx context.Context, center valueobject.Coordinate, radiusMeters float64, limit int64) ([]*entity.Place, error) {
	return r.inner.FindNearby(ctx, center, radiusMeters, limit)
}
