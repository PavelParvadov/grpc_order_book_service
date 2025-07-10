package suite

import (
	"context"
	"fmt"
	"github.com/PavelParvadov/grpc_order_book_service/book-service/internal/config"
	bookServicev1 "github.com/PavelParvadov/grpc_order_book_service/book-service/protoc/gen/go/book-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"testing"
	"time"
)

type Suite struct {
	T          *testing.T
	Cfg        *config.Config
	BookClient bookServicev1.BookClient
}

func NewSuite(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()
	cfg := config.LoadConfigByPath("../config/config.yaml")
	ctx, canc := context.WithTimeout(context.Background(), 10*time.Second)

	t.Cleanup(func() {
		t.Helper()
		canc()
	})

	cc, err := grpc.DialContext(ctx, fmt.Sprintf("localhost:%d", cfg.GRPCConf.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return ctx, &Suite{
		T:          t,
		Cfg:        cfg,
		BookClient: bookServicev1.NewBookClient(cc),
	}
}
