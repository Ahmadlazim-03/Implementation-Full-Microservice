package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"

	"backend/internal/places/domain/entity"
	"backend/internal/places/domain/repository"
)

// CachedCategoryRepository — categories rarely change, so we cache them for 1h.
type CachedCategoryRepository struct {
	inner repository.CategoryRepository
	rdb   *redis.Client
	ttl   time.Duration
}

func NewCachedCategoryRepository(inner repository.CategoryRepository, rdb *redis.Client, ttl time.Duration) *CachedCategoryRepository {
	if ttl <= 0 {
		ttl = time.Hour
	}
	return &CachedCategoryRepository{inner: inner, rdb: rdb, ttl: ttl}
}

type cachedCat struct{ ID, Name, Icon string }

const allCatsKey = "categories:all"

func (r *CachedCategoryRepository) Save(ctx context.Context, c *entity.Category) error {
	if err := r.inner.Save(ctx, c); err != nil {
		return err
	}
	_ = r.rdb.Del(ctx, allCatsKey).Err()
	return nil
}

func (r *CachedCategoryRepository) FindByID(ctx context.Context, id string) (*entity.Category, error) {
	return r.inner.FindByID(ctx, id) // single ID lookups infrequent — skip cache
}

func (r *CachedCategoryRepository) FindAll(ctx context.Context) ([]*entity.Category, error) {
	if raw, err := r.rdb.Get(ctx, allCatsKey).Bytes(); err == nil {
		var list []cachedCat
		if json.Unmarshal(raw, &list) == nil {
			out := make([]*entity.Category, 0, len(list))
			for _, c := range list {
				out = append(out, entity.NewCategory(c.ID, c.Name, c.Icon))
			}
			return out, nil
		}
	}

	items, err := r.inner.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	cached := make([]cachedCat, 0, len(items))
	for _, c := range items {
		cached = append(cached, cachedCat{ID: c.ID(), Name: c.Name(), Icon: c.Icon()})
	}
	if data, err := json.Marshal(cached); err == nil {
		_ = r.rdb.Set(ctx, allCatsKey, data, r.ttl).Err()
	}
	return items, nil
}
