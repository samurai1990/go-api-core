package errors

import "errors"

var ErrBadInput = errors.New("bad input")
var ErrConnectionDB = errors.New("failed to connect database")
var ErrMigrationDB = errors.New("failed migrate schema")
var ErrEnvironmentValue = errors.New("not set environment variable")
var ErrInternal = errors.New("internal error")
