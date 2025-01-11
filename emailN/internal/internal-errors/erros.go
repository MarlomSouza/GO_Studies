package internalerrors

import "errors"

var ErrInternal error = errors.New("internal server error")
var ErrNotFound error = errors.New("not found")

func ProcessInternalError(err error) error {
	if errors.Is(err, ErrNotFound) {
		return ErrNotFound
	}
	return ErrInternal
}
