package tests

import (
	bookServicev1 "github.com/PavelParvadov/grpc_order_book_service/book-service/protoc/gen/go/book-service"
	"github.com/PavelParvadov/grpc_order_book_service/book-service/tests/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddBook_HappyPath(t *testing.T) {
	ctx, s := suite.NewSuite(t)
	req := bookServicev1.AddBookRequest{
		Name:   "Тест",
		Author: "Тест",
	}
	resp, err := s.BookClient.AddBook(ctx, &req)
	require.NoError(t, err)
	assert.NotEmpty(t, resp.Id)

}

func TestAddBook_Fail(t *testing.T) {
	ctx, s := suite.NewSuite(t)
	tests := []struct {
		name        string
		input       *bookServicev1.AddBookRequest
		expectedErr string
	}{
		{
			name: "empty author",
			input: &bookServicev1.AddBookRequest{
				Name:   "test2",
				Author: "",
			},
			expectedErr: "name or author is empty",
		},
		{
			name: "empty name",
			input: &bookServicev1.AddBookRequest{
				Name:   "",
				Author: "test",
			},
			expectedErr: "name or author is empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.BookClient.AddBook(ctx, tt.input)
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}
