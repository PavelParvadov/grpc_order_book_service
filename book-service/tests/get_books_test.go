package tests

import (
	bookServicev1 "github.com/PavelParvadov/grpc_order_book_service/book-service/protoc/gen/go/book-service"
	"github.com/PavelParvadov/grpc_order_book_service/book-service/tests/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetBooks_HappyPath(t *testing.T) {
	ctx, s := suite.NewSuite(t)
	resp, err := s.BookClient.GetBooks(ctx, &bookServicev1.GetBooksRequest{})
	require.NoError(t, err)
	assert.NotEmpty(t, resp.GetBooks())

}
