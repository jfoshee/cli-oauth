package api

// TokenResponse represents the response from the token endpoint,
// potentially including access, refresh, and identity tokens.
type TokenResponse struct {
	// AccessToken details.
	AccessToken string
	// RefreshToken for obtaining new access tokens.
	RefreshToken string
	// TokenType, e.g., "bearer".
	TokenType string
	// Scope granted by the access token.
	Scope string
	// IDToken is the OpenID Connect ID token (a JWT) providing
	// information about the authenticated user.
	IDToken string
}

// TokenResponse extracts the token information (including optional
// ID token) from a FormResponse.
func (f FormResponse) TokenResponse() (*TokenResponse, error) {
	accessToken := f.Get("access_token")
	idToken := f.Get("id_token")

	// TODO: Re-evaluate this. probably should check status code and idToken is not always present
	if accessToken == "" { // && idToken == "" {
		return nil, f.Err()
	}

	return &TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: f.Get("refresh_token"),
		TokenType:    f.Get("token_type"),
		Scope:        f.Get("scope"),
		IDToken:      idToken,
	}, nil
}
