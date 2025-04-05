package user

import "errors"

type ctxKeyType string

const userKey ctxKeyType = "user"

var ErrUserNotFound = errors.New("user not found")
