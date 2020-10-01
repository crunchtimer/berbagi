package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/hourglasshoro/berbagi/src/app/domain/value_object"
	"time"
)

type TwitterAccessTokenRepository struct {
	Redis *redis.Client
}

func NewTwitterAccessTokenRepository(
	redis *redis.Client,
) *TwitterAccessTokenRepository {
	repo := new(TwitterAccessTokenRepository)
	repo.Redis = redis
	return repo
}

func (repo *TwitterAccessTokenRepository) SetTwitterToken(token value_object.AccessToken, ttl time.Duration) (err error) {
	key := fmt.Sprintf("twitterToken:%s", uuid.New())
	err = repo.Redis.Set(ctx, key, token.Token, ttl).Err()
	return
}
