package keys_manager_errors

import "errors"

var (
	ErrKeyAlreadyExists   = errors.New("key already exists in the `DB`")
	ErrServerKeyNotExists = errors.New("server key is not yet generated.")
)
