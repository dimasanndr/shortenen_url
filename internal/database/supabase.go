package database

import (
	"context"
	"fmt"
	"log"
	"url_shortener_api/internal/config"

	"github.com/jackc/pgx/v5"
)

func NewDbSupabase(ctx context.Context, cfg *config.Config) (*pgx.Conn, error) {
	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v", cfg.DbConfig.DbUser, cfg.DbConfig.DbPassword, cfg.DbConfig.DbHost, cfg.DbConfig.DbPort, cfg.DbConfig.DbName)
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	//test connection
	var version string
	if err := conn.QueryRow(ctx, "SELECT version()").Scan(&version); err != nil {
		return nil, err
	}

	log.Println(version)

	return conn, nil

}
