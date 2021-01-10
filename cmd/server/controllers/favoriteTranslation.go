package controllers

import (
	"context"
	"errors"
	"transign/cmd/server/config"
	"transign/cmd/server/models"
	pb "transign/gen"

	"gorm.io/gorm"
)

// FavoriteTranslationServer has methods for service FavoriteTranslation
type FavoriteTranslationServer struct {
	pb.UnimplementedFavoriteTranslationServer
}

// GetFavorite is a function for handling retrieving favorite translations
func (s *FavoriteTranslationServer) GetFavorite(context context.Context, in *pb.UUIDMessage) (*pb.GetFavoriteTranslationResponse, error) {
	var translations []models.FavoriteTranslation
	var requestHistory []string
	config.DB.Where(&models.FavoriteTranslation{UUID: in.Uuid}).Find(&translations)
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
	query := config.DB.Where(&models.FavoriteTranslation{UUID: in.Uuid, Text: in.Text})
	if err := query.First(&translations).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// create one
		translation := models.FavoriteTranslation{UUID: in.Uuid, Text: in.Text}
		config.DB.Create(&translation)
	} else {
		// remove one
		query.Delete(&translations)
	}

	return s.GetFavorite(context, &pb.UUIDMessage{Uuid: in.Uuid})

	// return response with updated db status
}
