package redisdb

import (
	"encoding/json"
	"fmt"
	"log"
	botdatabase "medium_scraper/inputports/telegram_bot/bot_database"

	"github.com/go-redis/redis"
)

type repo struct {
	rdb *redis.Client
}

func NewRepo(redisUrl string, password string) *repo {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: password,
		DB:       0,
	})

	result := rdb.Ping()
	if result.Err() != nil {
		log.Fatal(result.Err())
	}
	return &repo{
		rdb: rdb,
	}
}

func (r repo) AddUser(chatId int64, user botdatabase.User) error {
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}
	result := r.rdb.Set(fmt.Sprint(chatId), body, 0)
	return result.Err()
}

func (r repo) IsNewUser(chatId int64) (bool, error) {
	result := r.rdb.Get(fmt.Sprint(chatId))
	if result.Err() != nil {
		if result.Err() == redis.Nil {
			return true, nil
		} else {
			return true, result.Err()
		}
	}
	return false, nil
}

func (r repo) GetUser(chatId int64) (*botdatabase.User, error) {
	result := r.rdb.Get(fmt.Sprint(chatId))
	if result.Err() != nil {
		return nil, result.Err()
	}
	body := result.Val()
	user := new(botdatabase.User)
	err := json.Unmarshal([]byte(body), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r repo) UpdateUserCommand(chatId int64, command string) error {
	user, err := r.GetUser(chatId)
	if err != nil {
		return err
	}
	user.LastCommand = command
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}
	result := r.rdb.Set(fmt.Sprint(chatId), body, 0)
	return result.Err()
}

func (r repo) UpdateUserArticles(chatId int64, articles map[string]string) error {
	user, err := r.GetUser(chatId)
	if err != nil {
		return err
	}
	user.LastArticles = articles
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}
	result := r.rdb.Set(fmt.Sprint(chatId), body, 0)
	return result.Err()
}
