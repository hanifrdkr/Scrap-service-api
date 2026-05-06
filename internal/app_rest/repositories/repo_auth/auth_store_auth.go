package repo_auth

import (
	"context"
	"github.com/google/uuid"
	"helicopter-hr/internal/app_rest/model"
)

func (a *authRepository) StoreAuth(ctx context.Context, p model.Auth) error {
	const q = `INSERT INTO auths (id,user_id,email,password,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6);`

	_, err := a.db.Exec(
		q,
		uuid.New().String(),
		p.UserID,
		p.Email,
		p.Password,
		p.CreatedAt,
		p.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
