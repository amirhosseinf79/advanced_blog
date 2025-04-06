package dto

import "errors"

var ErrUserDoesNotExist = errors.New("user does not exist")
var ErrUserAlreadyExists = errors.New("user already exists")
var ErrInvalidPassword = errors.New("invalid username or password")
var ErrInvalidToken = errors.New("invalid token")
