package redisx

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
)

var _r *redis.Client
var once sync.Once

type Redis struct {
	cfg Config
	db  *redis.Client
}

func NewRedis(cfg Config) (*Redis, error) {
	_r = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.Db,
	})

	if _, _err := _r.Ping(context.Background()).Result(); _err != nil {
		return nil, _err
	}
	return &Redis{cfg, _r}, nil

}

func (this *Redis) GetDb() (db *redis.Client) {
	return this.db
}
