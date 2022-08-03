package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/todolist-project-rhttraining/src/userservice/config"
	"github.com/todolist-project-rhttraining/src/userservice/pkg/repository/db"
)

type DatabaseRepo struct {
	conn    *pgx.Conn
	queries *db.Queries
}

func NewDatabaseRepo(ctx context.Context, conf config.DBConfig) (DatabaseRepo, error) {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		conf.User, conf.Pass, conf.Host, conf.Port, conf.Name, conf.SslMode)

	connCfg, err := pgx.ParseConfig(connString)
	if err != nil {
		return DatabaseRepo{}, err
	}

	conn, err := pgx.ConnectConfig(ctx, connCfg)
	if err != nil {
		return DatabaseRepo{}, err
	}

	return DatabaseRepo{
		conn:    conn,
		queries: db.New(conn),
	}, nil
}

func (db DatabaseRepo) Query() db.Querier {
	return db.queries
}

func (db DatabaseRepo) Tx(ctx context.Context, fn func(tx db.Querier) error) error {
	tx, err := db.conn.BeginTx(ctx, pgx.TxOptions{})
	qtx := db.queries.WithTx(tx)
	if err != nil {
		return err
	}
	defer func() {
		rollbackErr := tx.Rollback(ctx)
		if rollbackErr != nil && !errors.Is(rollbackErr, pgx.ErrTxClosed) {
			err = rollbackErr
		}
	}()

	fErr := fn(qtx)
	if fErr != nil {
		_ = tx.Rollback(ctx) // ignore rollback error as there is already an error to return
		return fErr
	}

	return tx.Commit(ctx)
}
