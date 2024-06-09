package repos

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/dolldot/scrapyard/bookmarkd/generated/ent"
	"github.com/dolldot/scrapyard/bookmarkd/generated/ent/bookmark"
	"github.com/dolldot/scrapyard/bookmarkd/internal/logger"
	"github.com/google/uuid"
)

type BookmarkRepo struct {
	db *ent.Client
}

func NewBookmarkRepo(db *ent.Client) *BookmarkRepo {
	return &BookmarkRepo{db: db}
}

func (r *BookmarkRepo) FindAll(ctx context.Context, number string) ([]*ent.Bookmark, error) {
	data, err := r.db.Bookmark.Query().
		Where(bookmark.AccountNumber(number)).
		Order(
			bookmark.ByCreatedAt(
				sql.OrderDesc(),
			),
		).
		All(ctx)
	if err != nil {
		logger.Errorf("err: %v", err)
		return nil, err
	}

	return data, nil
}

func (r *BookmarkRepo) Create(ctx context.Context, bookmark ent.Bookmark) (*ent.Bookmark, error) {
	data, err := r.db.Bookmark.Create().
		SetName(bookmark.Name).
		SetURL(bookmark.URL).
		SetAccountNumber(bookmark.AccountNumber).
		Save(ctx)
	if err != nil {
		logger.Errorf("err: %v", err)
		return nil, err
	}

	return data, nil
}

func (r *BookmarkRepo) FindByUUID(ctx context.Context, number string, uuid uuid.UUID) (*ent.Bookmark, error) {
	data, err := r.db.Bookmark.Query().
		Where(bookmark.UUID(uuid)).
		Only(ctx)
	if err != nil {
		logger.Errorf("err: %v", err)
		return nil, err
	}

	return data, nil
}

func (r *BookmarkRepo) UpdateByUUID(ctx context.Context, number string, uuid uuid.UUID, content ent.Bookmark) error {
	now := time.Now()

	_, err := r.db.Bookmark.Update().
		SetName(content.Name).
		SetURL(content.URL).
		SetUpdatedAt(now).
		Where(bookmark.UUID(uuid)).
		Save(ctx)

	if err != nil {
		logger.Errorf("err: %v", err)
		return err
	}

	return nil
}

func (r *BookmarkRepo) DeleteByUUID(ctx context.Context, number string, uuid uuid.UUID) error {
	_, err := r.db.Bookmark.Delete().
		Where(bookmark.UUID(uuid)).
		Exec(ctx)

	if err != nil {
		logger.Errorf("err: %v", err)
		return err
	}

	return nil
}
