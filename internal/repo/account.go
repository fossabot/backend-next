package repo

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"

	"github.com/penguin-statistics/backend-next/internal/model"
	"github.com/penguin-statistics/backend-next/internal/pkg/pgerr"
)

const AccountMaxRetries = 100

type Account struct {
	db *bun.DB
}

func NewAccount(db *bun.DB) *Account {
	return &Account{db: db}
}

// Before v3.3.7, approximately released at 2022-06-05 01:00, PenguinIDs are generated as 8 digits number string.
// After v3.3.7, newly generated PenguinIDs will be a 9 digits number string.
// PenguinID can start with 0, with generated number padded to the corresponding length with 0.
func generateRandomPenguinId() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%09d", rand.Intn(1e9))
}

func (c *Account) CreateAccountWithRandomPenguinId(ctx context.Context) (*model.Account, error) {
	// retry if account already exists
	for i := 0; i < AccountMaxRetries; i++ {
		account := &model.Account{
			PenguinID: generateRandomPenguinId(),
		}
		_, err := c.db.NewInsert().
			Model(account).
			Returning("account_id").
			Exec(ctx)
		if err != nil {
			log.Warn().Err(err).Int("retry", i).Msg("failed to create account. retrying...")
			continue
		} else if i > 0 {
			log.Info().
				Int("retry", i).
				Str("finalizedPenguinID", account.PenguinID).
				Msg("successfully created account after retry")
		}
		return account, nil
	}

	return nil, pgerr.ErrInternalError.Msg("failed to create account")
}

func (c *Account) GetAccountById(ctx context.Context, accountId string) (*model.Account, error) {
	var account model.Account

	err := c.db.NewSelect().
		Model(&account).
		Column("account_id").
		Where("account_id = ?", accountId).
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, pgerr.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &account, nil
}

func (c *Account) GetAccountByPenguinId(ctx context.Context, penguinId string) (*model.Account, error) {
	var account model.Account

	err := c.db.NewSelect().
		Model(&account).
		Where("penguin_id = ?", penguinId).
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, pgerr.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &account, nil
}

func (c *Account) IsAccountExistWithId(ctx context.Context, accountId int) bool {
	var account model.Account

	err := c.db.NewSelect().
		Model(&account).
		Where("account_id = ?", accountId).
		Scan(ctx, &account)
	if err != nil {
		return false
	}

	return account.AccountID > 0
}
