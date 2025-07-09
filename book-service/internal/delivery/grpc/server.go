package grpc

import (
	"context"
	"github.com/PavelParvadov/grpc_order_book_service/book-service/internal/domain/model"
	bookServicev1 "github.com/PavelParvadov/grpc_order_book_service/book-service/protoc/gen/go/book-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookService interface {
	GetBooks(ctx context.Context) ([]model.Book, error)
	AddBook(ctx context.Context, name, author string) (int64, error)
}

type RPCServer struct {
	bookService BookService
	bookServicev1.UnimplementedBookServer
}

func RegisterGRPCServer(server *grpc.Server, bookService BookService) {
	bookServicev1.RegisterBookServer(server, &RPCServer{bookService: bookService})
}

func (s *RPCServer) GetBooks(ctx context.Context, req *bookServicev1.GetBooksRequest) (*bookServicev1.GetBooksResponse, error) {
	books, err := s.bookService.GetBooks(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting books: %v", err)
	}

	pbBooks := make([]*bookServicev1.BookData, 0, len(books))
	for _, book := range books {
		pbBook := &bookServicev1.BookData{
			Name:   book.Name,
			Author: book.Author,
		}
		pbBooks = append(pbBooks, pbBook)
	}
	return &bookServicev1.GetBooksResponse{
		Books: pbBooks,
	}, nil

}

func (s *RPCServer) AddBook(ctx context.Context, req *bookServicev1.AddBookRequest) (*bookServicev1.AddBookResponse, error) {
	if req.GetName() == "" || req.GetAuthor() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name or author is empty")
	}
	id, err := s.bookService.AddBook(ctx, req.Name, req.Author)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error adding book: %v", err)
	}
	return &bookServicev1.AddBookResponse{
		Id: id,
	}, nil

}
