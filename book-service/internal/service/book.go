package service

import (
	"context"
	"errors"
	"github.com/PavelParvadov/grpc_order_book_service/book-service/internal/domain/model"
	"github.com/PavelParvadov/grpc_order_book_service/book-service/internal/repository/postgres"
	"go.uber.org/zap"
)

func (b *BookService) GetBooks(ctx context.Context) ([]model.Book, error) {
	books, err := b.BookProvider.GetBooks(ctx)
	if err != nil {
		if errors.Is(err, postgres.ErrBookNotFound) {
			b.Log.Error("Books not found", zap.Error(err))
			return []model.Book{}, ErrBookNotFound
		}
		b.Log.Error("GetBooks error", zap.Error(err))
		return nil, err
	}
	return books, nil
}

func (b *BookService) AddBook(ctx context.Context, name, author string) (int64, error) {
	id, err := b.BookSaver.Save(ctx, name, author)
	if err != nil {
		if errors.Is(err, postgres.ErrBookExists) {
			b.Log.Error("Book already exists", zap.Error(err))
			return 0, ErrBookExists
		}
		b.Log.Error("AddBook error", zap.Error(err))
		return 0, err
	}
	return id, nil
}
