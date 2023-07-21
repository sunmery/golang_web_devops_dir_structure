package test

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
)

var ctx = context.Background()

func TestRedis(t *testing.T) {
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "password", // no password set
		DB:       0,          // use default DB
	})

	// 存
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// 取
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Errorf("error: %s", err)
	}
	t.Logf("val is: %v", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		t.Error("key2 does not exist")
	} else if err != nil {
		t.Errorf("error: %s", err)
	}

	t.Logf("key2 is %s", val2)

}
