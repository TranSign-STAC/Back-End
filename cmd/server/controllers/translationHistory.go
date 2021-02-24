package controllers

import (
	"context"
	"transign/cmd/server/models"
	"transign/configs"
	pb "transign/gen"
)

// TranslationHistoryServer has methods for service translationHistory
type TranslationHistoryServer struct {
	pb.UnimplementedTranslationHistoryServer
}

// GetHistory is a function for handling retrieving list of history request
func (s *TranslationHistoryServer) GetHistory(ctx context.Context, in *pb.UUIDMessage) (*pb.TranslationHistoryResponse, error) {
	var translations []models.Translation
	var requestHistory []*pb.TextToSignLangRequest

	if in.Uuid == "" {
		requestHistory = make([]*pb.TextToSignLangRequest, 0)
		response := &pb.TranslationHistoryResponse{History: requestHistory}
		return response, nil
	}

	configs.DB.Where(&models.Translation{UUID: in.Uuid}).Find(&translations)
	requestHistory = make([]*pb.TextToSignLangRequest, len(translations))

	for i, translation := range translations {
		requestHistory[i] = &pb.TextToSignLangRequest{Uuid: translation.UUID, Text: translation.Text}
	}

	response := &pb.TranslationHistoryResponse{History: requestHistory}
	return response, nil
}
