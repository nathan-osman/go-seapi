package seapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// Request represents a query to the Stack Exchange API to perform an action or
// retrieve data.
type Request struct {
	httpMethod string
	method     string
	params     url.Values
}

// NewRequest creates a request for the specified method.
func NewRequest(method string) *Request {
	method = "/2.2" + method
	return &Request{
		httpMethod: http.MethodGet,
		method:     method,
		params:     make(url.Values),
	}
}

// Auth provides authentication data with the request. Only the fields in auth
// that contain valid data will included in the final URL.
func (r *Request) Auth(auth *Auth) *Request {
	if len(auth.AccessToken) != 0 {
		r.params.Add("access_token", auth.AccessToken)
	}
	if len(auth.Key) != 0 {
		r.params.Add("key", auth.Key)
	}
	return r
}

// Do executes the request and returns the response as a Value if successful.
func (r *Request) Do() (Value, error) {
	req, err := http.NewRequest(
		r.httpMethod,
		fmt.Sprintf(
			"https://api.stackexchange.com%s?%s",
			r.method,
			r.params.Encode(),
		),
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "go-seapi")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	v := make(Value)
	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return nil, err
	}
	if errMsg := v.String("error_message"); len(errMsg) != 0 {
		return nil, errors.New(errMsg)
	}
	return v, nil
}

// Param specifies a query string parameter
func (r *Request) Param(key, value string) *Request {
	r.params.Add(key, value)
	return r
}

// Site specifies the site that the request should be directed to.
func (r *Request) Site(site string) *Request {
	return r.Param("site", site)
}

// Sort sets the sorting order for the items in the response.
func (r *Request) Sort(sort string) *Request {
	return r.Param("sort", sort)
}
