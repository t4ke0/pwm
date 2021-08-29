package api

import (
	"strings"
)

// RegisterRequest
type RegisterRequest struct {
	Username Field `json:"username"`
	Password Field `json:"password"`
	Email    Field `json:"email"`
}

// Field string type alias that represents json body field.
type Field string

// IsEmpty check whether a field is empty or not.
func (f Field) IsEmpty() bool {
	if strings.TrimSpace(f.String()) == "" {
		return true
	}
	return false
}

// String returns field as a string.
func (f Field) String() string {
	return string(f)
}
