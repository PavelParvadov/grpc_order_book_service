package service

import (
	"context"
	"github.com/PavelParvadov/grpc_order_book_service/book-service/internal/domain/model"
	"go.uber.org/zap"
)

type BookService struct {
	Log *zap.Logger
	BookSaver
	BookProvider
}

func NewBookService(log *zap.Logger, bs BookSaver, bp BookProvider) *BookService {
	return &BookService{
		Log:          log,
		BookSaver:    bs,
		BookProvider: bp,
	}
}

type BookSaver interface {
	Save(ctx context.Context, name, author string) (int64, error)
}
type BookProvider interface {
	GetBooks(ctx context.Context) ([]model.Book, error)
	GetBookById(ctx context.Context, id int64) (*model.Book, error)
}
