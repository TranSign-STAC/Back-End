# TODO

## Refactoring for transign api v2

### Gen

-   [ ] Update gen for transign api v2

### Models

-   [ ] Add 'isFavorite' field to [translation model](cmd/server/models/translation.go)
-   [ ] Remove [FavoriteTranslation model](cmd/server/models/translation.go)

### Controllers

-   [ ] In [textToSignLang](cmd/server/controllers/textToSignLang.go)
    -   [ ] Rename TextToSignLangRequest to UUIDTextMessage
    -   [ ] Replace TextToSignLangResponse to Translation
-   [ ] in [translationHistory](cmd/server/controllers/translationHistory.go)
    -   [ ] Replace TranslationHistoryResponse to Translations in GetHistory
    -   [ ] Replace UUIDMessage to RemoveHistoryRequest in RemoveHistory
    -   [ ] Replace RemoveHistoryRequest to Translations in RemoveHistory
-   [ ] in [favoriteTranslation](cmd/server/controllers/favoriteTranslation.go)
    -   [ ] Replace GetFavoriteTranslationResponse to Translations in GetFavorite
    -   [ ] Replace ToggleFavorite to SetFavorite
    -   [ ] Add RemoveFavorite
