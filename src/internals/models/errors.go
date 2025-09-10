package models

import "errors"

var ErrNoRecord = errors.New("SQL : The SQL is No Record")
var ErrInvalidCredentials = errors.New("model : invalid credential")
var ErrDuplicateEmail = errors.New("models : duplicate email")
