package controllers

import (
	"context"
	"transign/cmd/server/models"
	"transign/configs"
	pb "transign/gen"
)

// FavoriteTranslationServer has methods for service FavoriteTranslation
type FavoriteTranslationServer struct {
	pb.UnimplementedFavoriteTranslationServer
}

// GetFavorite is a function for handling retrieving favorite translations
func (s *FavoriteTranslationServer) GetFavorite(context context.Context, in *pb.UUIDMessage) (*pb.GetFavoriteTranslationResponse, error) {
	var translations []models.FavoriteTranslation
	var requestHistory []string
	configs.DB.Where(&models.FavoriteTranslation{UUID: in.Uuid}).Find(&translations)
	requestHistory = make([]string, len(translations))

	for i, translate := range translations {
		requestHistory[i] = translate.Text
	}

	response := &pb.GetFavoriteTranslationResponse{Favorites: requestHistory}
	return response, nil
}

// ToggleFavorite is a function for handling toggling favorite translations
func (s *FavoriteTranslationServer) ToggleFavorite(context context.Context, in *pb.ToggleFavoriteTranslationRequest) (*pb.GetFavoriteTranslationResponse, error) {
	var translations models.FavoriteTranslation
	query := configs.DB.Where(&models.FavoriteTranslation{UUID: in.Uuid, Text: in.Text})
	var count int64
	query.Find(&translations).Count(&count)
	if count == 0 {
		// create one
		translation := models.FavoriteTranslation{UUID: in.Uuid, Text: in.Text}
		configs.DB.Create(&translation)
	} else {
		// remove one
		query.Delete(&translations)
	}

	return s.GetFavorite(context, &pb.UUIDMessage{Uuid: in.Uuid})

	// return response with updated db status
}
