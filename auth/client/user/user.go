package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

type Client struct {
	cl  *http.Client
	url string
}

type User struct {
	ID        int    `json:"id"`
	Role      int    `json:"role,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func New(host string) *Client {
	return &Client{
		cl: &http.Client{
			Transport: &http.Transport{
				DialContext:         (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
				TLSHandshakeTimeout: 5 * time.Second,
			},
			Timeout: 10 * time.Second,
		},
		url: host,
	}
}

func (c *Client) UserByID(ctx context.Context, id int) (*User, error) {
	if id == 0 {
		return nil, errors.New("empty id")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/users/v1/user/%d", c.url, id), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	res, err := c.cl.Do(req)
	if err != nil {
		return nil, fmt.Errorf("perform request: %w", err)
	}
	defer res.Body.Close()

	var u User
	if err := json.NewDecoder(res.Body).Decode(&u); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return &u, nil
}

func (c *Client) UserByEmailAndPassword(ctx context.Context, email, password string) (*User, error) {
	if email == "" {
		return nil, errors.New("empty email")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/users/v1/user?email=%s&pwd=%s", c.url, email, password), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	res, err := c.cl.Do(req)
	if err != nil {
		return nil, fmt.Errorf("perform request: %w", err)
	}
	defer res.Body.Close()

	var u User
	if err := json.NewDecoder(res.Body).Decode(&u); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return &u, nil
}
