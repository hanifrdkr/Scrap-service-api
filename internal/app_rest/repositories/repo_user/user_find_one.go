package repo_user

import (
	"context"
	"encoding/json"
	"helicopter-hr/internal/app_rest/model"
)

func (u *userRepository) FindOne(ctx context.Context, p model.User) (*model.User, error) {
	const q = `SELECT json_build_object(
    'id',users.id,
    'email',users.email,
    'name',users.name
) FROM users where users.email = $1;`

	var bson []byte
	err := u.db.QueryRowContext(ctx, q, p.Email).Scan(&bson)
	if err != nil {
		return nil, err
	}

	var result model.User
	if err := json.Unmarshal(bson, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
