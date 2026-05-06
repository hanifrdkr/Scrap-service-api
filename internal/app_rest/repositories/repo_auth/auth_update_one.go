package repo_auth

import (
	"context"
	"helicopter-hr/internal/app_rest/model"
)

func (a *authRepository) UpdateOne(ctx context.Context, p model.AccessToken) error {
	const q = `UPDATE access_tokens
			SET updated_at = $1, is_revoke = $2
			WHERE access_tokens.user_id = $3;`

	_, err := a.db.ExecContext(ctx, q, p.UpdatedAt, p.IsRevoke, p.UserID)
	if err != nil {
		return err
	}

	return nil
}
