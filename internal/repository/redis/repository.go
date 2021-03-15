// package redis

// import (
// 	"fmt"

// 	"github.com/go-redis/redis"
// 	"github.com/neel229/linktree-clone/internal/shortener"
// )

// type redisRepo struct {
// 	client *redis.Client
// }

// func newRedisClient(redisURL string) (*redis.Client, error) {
// 	opts, err := redis.ParseURL(redisURL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	client := redis.NewClient(opts)
// 	_, err = client.Ping().Result()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return client, nil
// }

// func NewRedisRepo(redisURL string) (*shortener.RedirectRepository, error) {
// 	repo := &redisRepo{}
// 	client, err := newRedisClient(redisURL)
// 	if err != nil {
// 		return nil, fmt.Errorf("error creating redis repo client: %v", err)
// 	}
// 	repo.client = client
// 	return repo, nil
// }

// func (rr *redisRepo) generateKey(code string) string {
// 	return fmt.Sprintf("redirect: %s", code)
// }

// func (rr *redisRepo) Find(code string) (*shortener.Redirect, error) {
// 	redirect := &shortener.Redirect{}
// 	key := rr.generateKey(code)
// 	data, err := rr.client.HGetAll(key).Result()
// 	if err != nil {
// 		return nil, fmt.Errorf("error finding redirect: %v", err)
// 	}
// 	if len(data == 0) {
// 		return nil, fmt.Errorf("error finding the redirect: %v", err)
// 	}
// 	redirect.Code = data["code"]
// 	redirect.Handles = data["handles"]
// 	return redirect, nil
// }

// func (rr *redisRepo) Store(redirect *shortener.Redirect) error {
// 	key := rr.generateKey(redirect.Code)
// 	data := map[string]interface{}{
// 		"code":    redirect.Code,
// 		"handles": redirect.Handles,
// 	}
// 	_, err := rr.client.HMSet(key, data).Result()
// 	if err != nil {
// 		return fmt.Errorf("error storing the data")
// 	}
// 	return nil
// }