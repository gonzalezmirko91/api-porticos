package db

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewPostgres(ctx context.Context, host string, port int, user, password, name, sslmode string) (*Postgres, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", url.QueryEscape(user), url.QueryEscape(password), host, port, name, sslmode)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	cfg.MaxConns = 10
	cfg.MinConns = 0
	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.MaxConnLifetime = 30 * time.Minute
	cfg.HealthCheckPeriod = 30 * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, cfg)

	if err != nil {
		return nil, err
	}

	pg := &Postgres{Pool: pool}

	pingCtx, cancel := context.WithTimeout(ctx, 3*time.Second)

	defer cancel()

	if err := pg.Pool.Ping(pingCtx); err != nil {
		pg.Pool.Close()
		return nil, err
	}

	return pg, nil
}

func (p *Postgres) Close() {
	if p != nil && p.Pool != nil {
		p.Pool.Close()
	}
}
