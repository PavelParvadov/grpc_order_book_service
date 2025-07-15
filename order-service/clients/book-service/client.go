package book_service

import (
	"context"
	bookServicev1 "github.com/PavelParvadov/grpc_order_book_service/book-service/protoc/gen/go/book-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BookClient struct {
	api bookServicev1.BookClient
}

func NewBookClient(ctx context.Context, addr string) (*BookClient, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := bookServicev1.NewBookClient(cc)
	return &BookClient{client}, nil
}

func (b *BookClient) GetBookByID(ctx context.Context, id int64) (*bookServicev1.BookData, error) {
	book, err := b.api.GetBookById(ctx, &bookServicev1.GetBookByIdRequest{BookId: id})
	if err != nil {
		return &bookServicev1.BookData{}, err
	}
	return book.Book, nil
}
