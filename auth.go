package seapi

// Auth maintains storage for credentials necessary to access certain methods
// in the API.
type Auth struct {
	Key          string
	ClientID     string
	ClientSecret string
	AccessToken  string
}

// NewRequest creates a new request using the authentication data.
func (a *Auth) NewRequest(method string) *Request {
	return NewRequest(method).Auth(a)
}
