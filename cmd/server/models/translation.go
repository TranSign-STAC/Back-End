package models

// Translation model is a model which contains information about translations
type Translation struct {
	ID        int32
	UUID      string
	Text      string
	RenderURL string
}
