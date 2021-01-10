package models

// Translation model is a model which contains information about translations
type Translation struct {
	ID        int32  `gorm:"primaryKey"`
	UUID      string `gorm:"not null"`
	Text      string `gorm:"not null"`
	RenderURL string
}

type FavoriteTranslation struct {
	ID   int32  `gorm:"primaryKey"`
	UUID string `gorm:"not null"`
	Text string `gorm:"not null"`
}
