// Code generated by goa v3.13.2, DO NOT EDIT.
//
// account HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/loafoe/sailpoint/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	account "github.com/loafoe/sailpoint/gen/account"
	goahttp "goa.design/goa/v3/http"
)

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "account" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateAccountPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("account", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the account
// create server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*account.CreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("account", "create", "*account.CreatePayload", v)
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("account", "create", err)
		}
		req.SetBasicAuth(p.Username, p.Password)
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the account
// create endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("account", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "account" service "delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		accountID string
	)
	{
		p, ok := v.(*account.DeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("account", "delete", "*account.DeletePayload", v)
		}
		accountID = p.AccountID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteAccountPath(accountID)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("account", "delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteRequest returns an encoder for requests sent to the account
// delete server.
func EncodeDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*account.DeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("account", "delete", "*account.DeletePayload", v)
		}
		req.SetBasicAuth(p.Username, p.Password)
		return nil
	}
}

// DecodeDeleteResponse returns a decoder for responses returned by the account
// delete endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("account", "delete", resp.StatusCode, string(body))
		}
	}
}
