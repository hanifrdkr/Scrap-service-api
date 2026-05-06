package repo_user

import (
	"context"
	"helicopter-hr/internal/app_rest/model"
)

func (u *userRepository) StoreUser(ctx context.Context, p model.User) error {
	const q = `INSERT INTO users (id,name,email,created_at,updated_at) VALUES ($1,$2,$3,$4,$5);`

	_, err := u.db.ExecContext(
		ctx,
		q,
		p.ID,
		p.Name,
		p.Email,
		p.CreatedAt,
		p.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
