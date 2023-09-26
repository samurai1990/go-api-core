package errors

import "errors"

var ErrBadInput = errors.New("bad input")
var ErrDB = errors.New("error database")
var ErrEnvironmentValue = errors.New("not set environment variable")
var ErrInternal = errors.New("internal error")
var ErrDuplicateKey=errors.New("duplicate key value violates unique constraint")
