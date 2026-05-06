package repo_auth

import (
	"context"
	"github.com/google/uuid"
	"helicopter-hr/internal/app_rest/model"
)

func (a *authRepository) StoreAccessToken(ctx context.Context, p model.AccessToken) error {
	const q = `INSERT INTO access_tokens (id,user_id,token,refresh_token,is_revoke,expired_at,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8);`

	_, err := a.db.Exec(
		q,
		uuid.New().String(),
		p.UserID,
		p.Token,
		p.RefreshToken,
		false,
		p.ExpiredAt,
		p.CreatedAt,
		p.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
