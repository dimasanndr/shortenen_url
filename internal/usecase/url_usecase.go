package usecase

import (
	"context"
	"time"
	"url_shortener_api/internal/repository"
)

type UrlUsecase interface {
	CreateUrl(ctx context.Context, longUrl string, customCode string, expiresAt *time.Time) (string, error)
}
type urlUsecase struct {
	urlRepo repository.UrlRepository
}

func NewUrlUsecase(urlRepo repository.UrlRepository) UrlUsecase {
	return &urlUsecase{
		urlRepo: urlRepo,
	}
}

func (u *urlUsecase) CreateUrl(ctx context.Context, longUrl string, customCode string, expiresAt *time.Time) (string, error) {
	return u.urlRepo.CreateShortUrl(ctx, longUrl, customCode)
}
