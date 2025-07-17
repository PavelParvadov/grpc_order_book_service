package book

import (
	"context"
	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/domain/dto"
	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/domain/models"
	bookServicev1 "github.com/PavelParvadov/grpc_order_book_service/book-service/protoc/gen/go/book-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Client struct {
	Client bookServicev1.BookClient
}

func NewClient(ctx context.Context, addr string) (*Client, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Client{bookServicev1.NewBookClient(cc)}, nil
}

func (c *Client) AddBook(ctx context.Context, in dto.AddBookInput) (dto.AddBookOutput, error) {
	book, err := c.Client.AddBook(ctx, &bookServicev1.AddBookRequest{
		Name:   in.Name,
		Author: in.Author,
	})
	if err != nil {
		_, ok := status.FromError(err)
		if ok {
			return dto.AddBookOutput{}, ErrInternalGRPCServer
		}
		return dto.AddBookOutput{}, err
	}
	return dto.AddBookOutput{
		ID: book.Id,
	}, err
}

func (c *Client) GetBooks(ctx context.Context) ([]models.Book, error) {
	resp, err := c.Client.GetBooks(ctx, &bookServicev1.GetBooksRequest{})
	if err != nil {
		_, ok := status.FromError(err)
		if ok {
			return []models.Book{}, ErrInternalGRPCServer
		}
		return []models.Book{}, err
	}
	books := make([]models.Book, 0, len(resp.GetBooks()))
	for _, book := range resp.GetBooks() {
		books = append(books, models.Book{
			Name:   book.GetName(),
			Author: book.GetAuthor(),
		})
	}
	return books, nil
}
