package handler

import (
	"encoding/json"
	"net/http"
	"time"
	"url_shortener_api/internal/response"
	"url_shortener_api/internal/usecase"
)

type Handler interface {
	CreateShortUrl(w http.ResponseWriter, r *http.Request)
}
type handler struct {
	urlUsecase usecase.UrlUsecase
}

func NewUrlHandler(urlUsecase usecase.UrlUsecase) Handler {
	return &handler{
		urlUsecase: urlUsecase,
	}
}

func (h *handler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var reqShortUrl struct {
		Url        string `json:"url"`
		CustomCode string `json:"custom_code"`
		ExpiresAt  string `json:"expires_at,omitempty"`
	}

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&reqShortUrl); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"The request body is invalid",
			err.Error(),
		)
		return
	}

	// Validate required URL
	if reqShortUrl.Url == "" {
		response.Error(
			w,
			http.StatusBadRequest,
			"URL is required",
			"Please provide a valid URL",
		)
		return
	}

	// Parse expiration date
	var expiresAt *time.Time

	if reqShortUrl.ExpiresAt != "" {
		parsedTime, err := time.Parse(time.RFC3339, reqShortUrl.ExpiresAt)
		if err != nil {
			response.Error(
				w,
				http.StatusBadRequest,
				"Invalid expiration date",
				"expires_at must use RFC3339 format",
			)
			return
		}

		if !parsedTime.After(time.Now()) {
			response.Error(
				w,
				http.StatusBadRequest,
				"Invalid expiration date",
				"Expiration date must be in the future",
			)
			return
		}

		expiresAt = &parsedTime
	}

	// Create short URL
	url, err := h.urlUsecase.CreateUrl(
		r.Context(),
		reqShortUrl.Url,
		reqShortUrl.CustomCode,
		expiresAt,
	)

	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"Failed to create short URL",
			err.Error(),
		)
		return
	}

	response.Success(
		w,
		http.StatusCreated,
		"Short URL created successfully",
		url,
	)
}
