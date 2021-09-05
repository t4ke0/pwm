package api

// ErrResponse
type ErrResponse struct {
	ErrorMessage string `json:"error_message"`
}

// AuthResponse
type AuthResponse struct {
	JwtToken string `json:"token"`
}
