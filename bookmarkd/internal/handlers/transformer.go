package handlers

import (
	"time"

	"github.com/dolldot/scrapyard/bookmarkd/generated/ent"
	"github.com/google/uuid"
)

type BookmarkReponse struct {
	ID        int       `json:"id"`
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func bookmarkTransformer(data *ent.Bookmark) *BookmarkReponse {
	return &BookmarkReponse{
		ID:        data.ID,
		UUID:      data.UUID,
		Name:      data.Name,
		URL:       data.URL,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
