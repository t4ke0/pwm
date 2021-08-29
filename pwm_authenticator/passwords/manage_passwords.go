package passwords

import (
	"golang.org/x/crypto/bcrypt"
)

type HashedPassword []byte

func Hash(password []byte) (HashedPassword, error) {
	var cost int = 4
	gen, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		return nil, err
	}
	return gen, nil
}

// String
func (hp HashedPassword) String() string {
	return string(hp)
}

// IsCorrectPassword
func (hp HashedPassword) IsCorrectPassword(pw []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hp, pw); err != nil {
		return false
	}
	return true
}

// ToHashedPassword
func ToHashedPassword(pw string) HashedPassword {
	return HashedPassword([]byte(pw))
}
