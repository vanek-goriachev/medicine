package user

import (
	"context"
)

func StoreInContext(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func GetFromContext(ctx context.Context) (User, error) {
	user, ok := ctx.Value(userKey).(User)
	if !ok {
		return User{}, ErrUserNotFound
	}

	return user, nil
}
