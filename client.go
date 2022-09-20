package yrac

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Configuration struct {
	url    string
	client *http.Client
}

// Error ...
type Error struct {
	RequestURL  string `json:"-"`
	Err         string `json:"error"`
	Description string `json:"error_description"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s -> %s: %s", e.RequestURL, e.Err, e.Description)
}

func NewConfiguration(domain, bearerToken string) *Configuration {
	client := &http.Client{}
	client.Transport = RoundTripFunc(func(request *http.Request) (*http.Response, error) {
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept", "application/json")
		request.Header.Set("Authorization", "Bearer "+bearerToken)
		return http.DefaultTransport.RoundTrip(request)
	})
	return &Configuration{
		url:    domain + "/api",
		client: client,
	}
}

type Service struct {
	cfg *Configuration
}

func NewService(cfg *Configuration) Service {
	return Service{cfg: cfg}
}

func (s *Service) url(path string, query url.Values) (string, error) {
	var (
		u, err = url.Parse(s.cfg.url)
	)
	if err != nil {
		return "", err
	}
	u.Path += path
	u.RawQuery = query.Encode()
	return u.String(), nil
}

type RoundTripFunc func(*http.Request) (*http.Response, error)

func (fn RoundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}

func (s Service) call(ctx context.Context, method, path string, query url.Values, send interface{}, recv interface{}) error {
	var (
		b    []byte
		err1 error
	)
	if send != nil {
		if b, err1 = json.Marshal(send); err1 != nil {
			return err1
		}
	}
	u, err := s.url(path, query)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, method, u, bytes.NewReader(b))
	if err != nil {
		return err
	}
	resp, err := s.cfg.client.Do(req)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if _, err1 = io.Copy(&buf, resp.Body); err1 != nil {
		return err1
	}
	if err = resp.Body.Close(); err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		var e Error
		if err = json.Unmarshal(buf.Bytes(), &e); err != nil {
			return err
		}
		e.RequestURL = method + " " + u
		return e
	}
	if recv != nil {
		return json.Unmarshal(buf.Bytes(), recv)
	}
	return nil
}
