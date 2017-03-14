package seapi

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	ReadInbox   = "read_inbox"
	NoExpiry    = "no_expiry"
	WriteAccess = "write_access"
	PrivateInfo = "private_info"
)

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

// StartExplicit retrieves the URL for beginning the explicit auth flow.
func (a *Auth) StartExplicit(scope, redirectUri, state string) string {
	v := make(url.Values)
	v.Add("client_id", a.ClientID)
	v.Add("scope", scope)
	v.Add("redirect_uri", redirectUri)
	v.Add("state", state)
	return fmt.Sprintf(
		"https://stackexchange.com/oauth?%s",
		v.Encode(),
	)
}

// FinishExplicit completes explicit auth.
func (a *Auth) FinishExplicit(code, redirectUri string) error {
	v := make(url.Values)
	v.Add("client_id", a.ClientID)
	v.Add("client_secret", a.ClientSecret)
	v.Add("code", code)
	v.Add("redirect_uri", redirectUri)
	req, err := http.NewRequest(
		http.MethodPost,
		"https://stackexchange.com/oauth/access_token",
		strings.NewReader(v.Encode()),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New("unable to finish explicit auth")
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	q, err := url.ParseQuery(string(b))
	if err != nil {
		return err
	}
	a.AccessToken = q.Get("access_token")
	return nil
}
