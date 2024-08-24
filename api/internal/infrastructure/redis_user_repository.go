package infrastructure

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/dodokey/go-chat-app/internal/domain"

	"github.com/go-redis/redis/v8"
)

type RedisUserRepository struct {
    client *redis.Client
}

func NewRedisUserRepository(client *redis.Client) *RedisUserRepository {
    return &RedisUserRepository{client: client}
}

func (r *RedisUserRepository) Save(user *domain.User) error {
    data, err := json.Marshal(user)
    if err != nil {
        return err
    }
    return r.client.Set(context.Background(), user.Username, data, 0).Err()
}

func (r *RedisUserRepository) FindByUsername(username string) (*domain.User, error) {
    data, err := r.client.Get(context.Background(), username).Result()
    if err == redis.Nil {
        return nil, errors.New("user not found")
    } else if err != nil {
        return nil, err
    }

    var user domain.User
    if err := json.Unmarshal([]byte(data), &user); err != nil {
        return nil, err
    }

    return &user, nil
}