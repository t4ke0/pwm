package api

// ListPasswordResponse JSON response of list user passwords endpoint.
type ListPasswordResponse struct {
	PasswordID        int    `json:"password_id"`
	Username          string `json:"username"`
	PlainTextPassword string `json:"plaintext_password"`
	Category          string `json:"category"`
	Site              string `json:"site"`
}

// StorePasswordRequest JSON request to store passwords of a user.
type StorePasswordRequest struct {
	Username          string `json:"username"`
	PlainTextPassword string `json:"plaintext_password"`
	Category          string `json:"category"`
	Site              string `json:"site"`
}

// CredItem
type CredItem string

const (
	Username CredItem = "Username"
	Password          = "Password"
	Category          = "Category"
	Site              = "Site"
)

// String
func (self CredItem) String() string {
	return string(self)
}

// Validate
func (self CredItem) Validate() bool {
	if self != Username && self != Password && self != Category && self != Site {
		return false
	}
	return true
}

// UpdateUserItemsRequest JSON request to update user's creds items such as
// `password`, `site`, etc ....
type UpdateUserItemsRequest struct {
	PasswordID int `json:"password_id"`
	Items      []struct {
		Item  CredItem `json:"item"`
		Value string   `json:"value"`
	} `json:"items"`
}

// DeleteUserCredRequest
type DeleteUserCredRequest struct {
	PasswordID int `json:"password_id"`
}

// Complexity
type Complexity string

const (
	PasswordEasy    Complexity = "Easy"
	PasswordMedium             = "Medium"
	PasswordComplex            = "Complex"
)

func (c Complexity) String() string {
	return string(c)
}

// GeneratePasswordRequest
type GeneratePasswordRequest struct {
	PasswordLength     int64      `json:"password_length"`
	PasswordComplexity Complexity `json:"password_complexity"`
}

// GeneratePasswordResponse
type GeneratePasswordResponse struct {
	GeneratedPassword string `json:"generated_password"`
}
