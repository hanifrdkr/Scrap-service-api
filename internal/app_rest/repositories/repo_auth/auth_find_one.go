package repo_auth

import (
	"context"
	"encoding/json"
	"helicopter-hr/internal/app_rest/model"
)

func (a *authRepository) FindOne(ctx context.Context, p model.Auth) (*model.Auth, error) {
	const q = `SELECT json_build_object(
    'id',auths.id,
    'email',auths.email,
    'password',auths.password
) FROM auths where auths.email = $1;`

	var bson []byte
	err := a.db.QueryRowContext(ctx, q, p.Email).Scan(&bson)
	if err != nil {
		return nil, err
	}

	var result model.Auth
	if err := json.Unmarshal(bson, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
