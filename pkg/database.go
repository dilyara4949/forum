package pkg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"fmt"
	"time"

)

const (
	maxConn           = 50
	healthCheckPeriod = 3 * time.Minute
	maxConnIdleTime   = 1 * time.Minute
	maxConnLifetime   = 3 * time.Minute
	minConns          = 10
	lazyConnect       = false
)

// NewPgxConn pool
func NewPgxConn(env *Env) (*pgxpool.Pool, error) {

	username := env.DBUser
	password := env.DBPass
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName
	ctx := context.Background()
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		username,
		password,
		dbname,
	)

	poolCfg, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		return nil, err
	}

	poolCfg.MaxConns = maxConn
	poolCfg.HealthCheckPeriod = healthCheckPeriod
	poolCfg.MaxConnIdleTime = maxConnIdleTime
	poolCfg.MaxConnLifetime = maxConnLifetime
	poolCfg.MinConns = minConns
	poolCfg.LazyConnect = lazyConnect

	connPool, err := pgxpool.ConnectConfig(ctx, poolCfg)
	if err != nil {
		return nil, errors.Wrap(err, "pgx.ConnectConfig")
	}

	fmt.Println("sussecc connection")

	return connPool, nil
}



func Close( p *pgxpool.Pool)  {
	if p != nil {
		p.Close()
	}
}

// const (
// 	maxConn           = 50
// 	healthCheckPeriod = 3 * time.Minute
// 	maxConnIdleTime   = 1 * time.Minute
// 	maxConnLifetime   = 3 * time.Minute
// 	minConns          = 10
// 	lazyConnect       = false
// )

// func NewSQLiteConn(env *Env) (*sql.DB, error) {
// 	// dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
// 	// 	env.DBUser,
// 	// 	env.DBPass,
// 	// 	env.DBHost,
// 	// 	env.DBPort,
// 	// 	env.DBName,
// 	// )
// 	// fmt.Println(dbPath)
// 	db, err := sql.Open("sqlite3", "forum.db")
// 	if err != nil {
// 		return nil, errors.New(err.Error() + ", sql.Open")
// 	}

// 	// Set connection pool configurations
// 	db.SetMaxOpenConns(maxConn)
// 	db.SetConnMaxLifetime(maxConnLifetime)
// 	db.SetMaxIdleConns(minConns)

// 	// Check connection health
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	if err := db.PingContext(ctx); err != nil {
// 		db.Close()
// 		return nil, errors.New(err.Error() + ", db.PingContext")
// 	}

// 	fmt.Println("Success connection with db")

// 	return db, nil
// }

// func Close(db *sql.DB) {
// 	if db != nil {
// 		db.Close()
// 	}
// }
