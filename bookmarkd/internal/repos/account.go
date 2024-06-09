package repos

import (
	"context"

	"github.com/dolldot/scrapyard/bookmarkd/generated/ent"
	"github.com/dolldot/scrapyard/bookmarkd/generated/ent/account"
	"github.com/dolldot/scrapyard/bookmarkd/internal/logger"
)

type AccountRepo struct {
	db *ent.Client
}

func NewAccountRepo(db *ent.Client) *AccountRepo {
	return &AccountRepo{db: db}
}

func (r *AccountRepo) CreateAccount(ctx context.Context, number string) (*ent.Account, error) {
	data, err := r.db.Account.Create().
		SetNumber(number).
		Save(ctx)
	if err != nil {
		logger.Errorf("err: %v", err)
		return nil, err
	}

	return data, nil
}

func (r *AccountRepo) FindAccountByNumber(ctx context.Context, number string) (*ent.Account, error) {
	data, err := r.db.Account.Query().
		Where(account.Number(number)).
		Only(ctx)
	if err != nil {
		logger.Errorf("err: %v", err)
		return nil, err
	}

	return data, nil
}
