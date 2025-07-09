package postgres

import (
	"context"
	"github.com/PavelParvadov/grpc_order_book_service/book-service/internal/domain/model"
)

func (s *Storage) GetBooks(ctx context.Context) ([]model.Book, error) {
	query, err := s.DbPool.Query(ctx, "SELECT * FROM books")
	if err != nil {
		return nil, ErrBookNotFound
	}
	var books []model.Book

	err = query.Scan(&books)
	if err != nil {
		return nil, err
	}
	return books, nil

}

func (s *Storage) Save(ctx context.Context, name, author string) (int64, error) {
	tx, err := s.DbPool.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)
	var BookId int64
	if err = tx.QueryRow(ctx, "INSERT INTO books (name, author) VALUES ($1, $2) returning id", name, author).Scan(&BookId); err != nil {
		return 0, err
	}

	return BookId, tx.Commit(ctx)
}
