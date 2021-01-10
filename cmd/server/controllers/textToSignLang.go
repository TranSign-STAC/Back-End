package controllers

import (
	"context"
	"fmt"

	"transign/cmd/server/config"
	"transign/cmd/server/models"
	pb "transign/gen"
)

// TextToSignLangServer has methods for service textToSignLang
type TextToSignLangServer struct {
	pb.UnimplementedTextToSignLangServer
}

// Translate is a function for handling translation request
func (s *TextToSignLangServer) Translate(ctx context.Context, in *pb.TextToSignLangRequest) (*pb.TextToSignLangResponse, error) {
	/* todo: forward to translation server*/
	/* todo: forward response from translation server*/
	URL := fmt.Sprintf("api.transign.me/rendered/%s.mp4", in.Text)
	translation := models.Translation{UUID: in.Uuid, Text: in.Text, RenderURL: URL}
	config.DB.Create(&translation)
	response := &pb.TextToSignLangResponse{RenderUrl: URL}
	return response, nil
}
