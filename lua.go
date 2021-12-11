package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)
var (
    zIncrScript = redis.NewScript(`
		local score = redis.call("ZSCORE", KEYS[1], ARGV[2])
		if (score==false) then score = 0 else score = math.floor(score) end
		return score + ARGV[1],redis.call("ZADD", KEYS[1], score + ARGV[1], ARGV[2])
	`)
  )
func main() {
	//ExampleClient()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.208.50.51:1429",
		Password: "7a5ebedd6de38430", // no password set
		DB:       0,                  // use default DB
	})

	zIncrScript := redis.NewScript(`
		local score = redis.call("ZSCORE", KEYS[1], ARGV[2])
		if (score==false) then score = 0 else score = math.floor(score) end
		return score + ARGV[1],redis.call("ZADD", KEYS[1], score + ARGV[1], ARGV[2])
	`)

  //递增zset 并返回总数
	score := float64(10)
	decimal, _ := strconv.ParseFloat(fmt.Sprintf("0.%d", time.Now().Unix()), 64)
	score += 1 - decimal
	data, err := zIncrScript.Run(ctx, rdb, []string{"testkey"}, score, 1).Result()
	fmt.Println(data.(int64), err)
	
}

var ctx = context.Background()

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.208.50.51:1429",
		Password: "7a5ebedd6de38430", // no password set
		DB:       0,                  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	key := fmt.Sprintf("gateway:meet_room:live_id_%v", 4594)
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	//val2, err := rdb.Get(ctx, "key2").Result()
	//if err == redis.Nil {
	//	fmt.Println("key2 does not exist")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("key2", val2)
	//}
	// Output: key value
	// key2 does not exist
}

func IsEmpty(err error) bool {
	// Redis
	if errors.Is(err, redis.Nil) {
		return true
	}
	return false
}
