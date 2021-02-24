package main

import (
	"context"
	"log"
	"net"
	"os"
	"testing"
	"transign/cmd/server/controllers"
	"transign/cmd/server/models"
	"transign/configs"
	pb "transign/gen"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	os.Setenv("DB_DATABASE", "test_database")
	// set for test environment
	configs.ConnectDB()
	// connect db

	configs.DB.Where("1 = 1").Delete(&models.Translation{})
	configs.DB.Where("1 = 1").Delete(&models.FavoriteTranslation{})
	// reset db

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	pb.RegisterTextToSignLangServer(s, &controllers.TextToSignLangServer{})
	pb.RegisterTranslationHistoryServer(s, &controllers.TranslationHistoryServer{})
	pb.RegisterFavoriteTranslationServer(s, &controllers.FavoriteTranslationServer{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func getConnection(ctx context.Context, t *testing.T) (conn *grpc.ClientConn) {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	return
}

func requestTranslation(ctx context.Context, client pb.TextToSignLangClient) {
	client.Translate(ctx, &pb.TextToSignLangRequest{Uuid: "test", Text: "test"})
}

func TestTextToSignLang(t *testing.T) {
	ctx := context.Background()
	conn := getConnection(ctx, t)
	defer conn.Close()
	client := pb.NewTextToSignLangClient(conn)

	if _, err := client.Translate(ctx, &pb.TextToSignLangRequest{Uuid: "test", Text: "test"}); err != nil {
		t.Fatalf("failed: %v", err)
	}
}

func TestTranslationHistory(t *testing.T) {
	ctx := context.Background()
	conn := getConnection(ctx, t)
	defer conn.Close()
	client := pb.NewTranslationHistoryClient(conn)

	resp, err := client.GetHistory(ctx, &pb.UUIDMessage{Uuid: "test"})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	assert.Equal(t, 1, len(resp.History))
	// one history should be returned

	ttslClient := pb.NewTextToSignLangClient(conn)
	for i := 0; i < 10; i++ {
		requestTranslation(ctx, ttslClient)
	}

	resp, err = client.GetHistory(ctx, &pb.UUIDMessage{Uuid: "test"})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	assert.Equal(t, 11, len(resp.History))
	// eleven history should be returned

	resp, err = client.GetHistory(ctx, &pb.UUIDMessage{})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	assert.Equal(t, 0, len(resp.History))
	// no history should be returned if the request happens with the empty UUID
}

func TestFavoriteTranslation(t *testing.T) {
	ctx := context.Background()
	conn := getConnection(ctx, t)
	defer conn.Close()
	client := pb.NewFavoriteTranslationClient(conn)

	_, err := client.ToggleFavorite(ctx, &pb.ToggleFavoriteTranslationRequest{Uuid: "test", Text: "토글"})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	resp, err := client.GetFavorite(ctx, &pb.UUIDMessage{Uuid: "test"})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	assert.Equal(t, "토글", resp.Favorites[0])
	// test set favorite

	_, err = client.ToggleFavorite(ctx, &pb.ToggleFavoriteTranslationRequest{Uuid: "test", Text: "토글"})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	resp, err = client.GetFavorite(ctx, &pb.UUIDMessage{Uuid: "test"})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	assert.Equal(t, 0, len(resp.Favorites))
	// test unset favorite

	resp, err = client.GetFavorite(ctx, &pb.UUIDMessage{})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	assert.Equal(t, 0, len(resp.Favorites))
	// no favorites should be returned if the request happens with the empty UUID
}
