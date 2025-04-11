package cterr

import "errors"

var (
	ErrorInvalid    = errors.New("error_invalid")
	ErrorNotExisted = errors.New("error_not_existed")
	ErrorUnknown    = errors.New("error_unknown")
)
