package model

import (
	"github.com/go-redis/redis/v8"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// mysql connector
const (
	dsn = "root:123456@(localhost:3306)/db_proj?timeout=5000ms&readTimeout=5000ms&writeTimeout=5000ms&charset=utf8mb4&parseTime=true&loc=Local"
)

type mySqlConnector struct {
	single sync.Once
	db     *gorm.DB
}

var connector = mySqlConnector{}

func NewMySqlConnector() *gorm.DB {
	connector.single.Do(func() {
		connector.db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	})

	return connector.db
}

// redis connector
type redisConnector struct {
	single   sync.Once
	clientId int
	client   *redis.Client
}

func newRedisClient(clientId int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Network: "tcp4",
		Addr:    "localhost:6379",
		DB:      clientId,
	})
}

var (
	TokenRedisHandler = redisConnector{clientId: 1}
	WALRedisHandler   = redisConnector{clientId: 2}
)

// 通过此接口获得实际的redis-cli, 懒汉式加载, redisConnector只是封装了单例
func (obj *redisConnector) GetRedisClient() *redis.Client {
	obj.single.Do(func() {
		obj.client = newRedisClient(obj.clientId)
	})

	return obj.client
}

// auto gen
func init() {
	db := NewMySqlConnector()
	db.AutoMigrate(&User{}, &Dish{}, &ConsumeRecord{})
}
