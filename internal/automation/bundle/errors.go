package bundle

import "errors"

var (
	ErrInvalidBundle      = errors.New("invalid bundle")
	ErrInvalidPackageJSON = errors.New("invalid package.json")
)
