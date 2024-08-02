package db

import (
	"context"
	"fmt"
	"log"
	"odilbekqazaqov4657/my_blog_backend/config"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func getConfig(cfg config.PgConfig) *pgxpool.Config {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifeTime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	var databaseURL string = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	dbConfig, err := pgxpool.ParseConfig(databaseURL)
	fmt.Println("go")
	if err != nil {
		fmt.Println("err", err)
		return nil
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifeTime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("before acquiring the connection pool to the database!!")
		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("after releasing the connection pool to the database!!")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		log.Println("closed the connection pool to the database!!")
	}

	return dbConfig
}

func ConnDb(cfg config.PgConfig) (*pgxpool.Pool, error) {
	connPool, err := pgxpool.NewWithConfig(context.Background(), getConfig(cfg))
	if err != nil {
		log.Println("error while creating connection to the database !")
	}

	connection, err := connPool.Acquire(context.Background())

	if err != nil {
		log.Println("error while acquiring connection from the database pool !", err)
	}

	defer connection.Release()

	err = connection.Ping(context.Background())

	if err != nil {
		log.Println("Could not piung database")
	}

	fmt.Println("connected successfully to db !")

	return connPool, nil
}
