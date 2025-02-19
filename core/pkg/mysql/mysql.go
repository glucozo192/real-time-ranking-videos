package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/glu/video-real-time-ranking/core/pkg/utils"
	"github.com/glu/video-real-time-ranking/ent"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

const (
	maxConn         = 50
	maxConnIdleTime = 1 * time.Minute
	maxConnLifetime = 3 * time.Minute
	minConns        = 10
	lazyConnect     = false
)

// NewMySQLConn pool
func NewMySQLConn(cfg *utils.Config) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "sql.Open")
	}

	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(minConns)
	db.SetConnMaxLifetime(maxConnLifetime)
	db.SetConnMaxIdleTime(maxConnIdleTime)

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "db.Ping")
	}

	return db, nil
}

func SetupEntClient(cfg *utils.Config) (*ent.Client, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	client, err := ent.Open("mysql", dataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "failed opening connection to mysql")
	}
	// Run the auto migration tool.
	_ = client.Schema.Create(context.Background())
	return client, nil
}
