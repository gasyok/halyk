package repository

import "github.com/gasyok/Assessment/internal/repository/postgres"

type Storage struct {
	txManager    postgres.TransactionManager
	pgRepository *postgres.PgRepository
}

func NewStorage(
	txManager postgres.TransactionManager,
	pgRepository *postgres.PgRepository,
) *Storage {
	return &Storage{
		txManager:    txManager,
		pgRepository: pgRepository,
	}
}
