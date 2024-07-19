package pkg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func InitDb(connection_str string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), connection_str)
	err = dbpool.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	return dbpool
}
