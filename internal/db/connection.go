package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"os"
)

var (
	Postgres *sql.DB
	Redis    *redis.Client
)

func Init() {
	dsn := os.Getenv("POSTGRES_DSN")
	db, _ := sql.Open("postgres", dsn)

	Postgres = db

	addr := os.Getenv("REDIS_ADDR")
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	Redis = rdb
}
