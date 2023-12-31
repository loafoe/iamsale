// Code generated by goa v3.13.2, DO NOT EDIT.
//
// aggregate HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/loafoe/iamsale/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	aggregate "github.com/loafoe/iamsale/gen/aggregate"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildAccountsRequest instantiates a HTTP request object with method and path
// set to call the "aggregate" service "accounts" endpoint
func (c *Client) BuildAccountsRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AccountsAggregatePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("aggregate", "accounts", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAccountsRequest returns an encoder for requests sent to the aggregate
// accounts server.
func EncodeAccountsRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*aggregate.AccountsPayload)
		if !ok {
			return goahttp.ErrInvalidType("aggregate", "accounts", "*aggregate.AccountsPayload", v)
		}
		req.SetBasicAuth(p.Username, p.Password)
		return nil
	}
}

// DecodeAccountsResponse returns a decoder for responses returned by the
// aggregate accounts endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeAccountsResponse may return the following errors:
//   - "PermissionDenied" (type *goa.ServiceError): http.StatusForbidden
//   - error: internal error
func DecodeAccountsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
		case http.StatusOK:
			var (
				body AccountsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aggregate", "accounts", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateAccountResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("aggregate", "accounts", err)
			}
			res := NewAccountsAccountOK(body)
			return res, nil
		case http.StatusForbidden:
			var (
				body AccountsPermissionDeniedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aggregate", "accounts", err)
			}
			err = ValidateAccountsPermissionDeniedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aggregate", "accounts", err)
			}
			return nil, NewAccountsPermissionDenied(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("aggregate", "accounts", resp.StatusCode, string(body))
		}
	}
}

// BuildGroupsRequest instantiates a HTTP request object with method and path
// set to call the "aggregate" service "groups" endpoint
func (c *Client) BuildGroupsRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GroupsAggregatePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("aggregate", "groups", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGroupsRequest returns an encoder for requests sent to the aggregate
// groups server.
func EncodeGroupsRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*aggregate.GroupsPayload)
		if !ok {
			return goahttp.ErrInvalidType("aggregate", "groups", "*aggregate.GroupsPayload", v)
		}
		req.SetBasicAuth(p.Username, p.Password)
		return nil
	}
}

// DecodeGroupsResponse returns a decoder for responses returned by the
// aggregate groups endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeGroupsResponse may return the following errors:
//   - "PermissionDenied" (type *goa.ServiceError): http.StatusForbidden
//   - error: internal error
func DecodeGroupsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
		case http.StatusOK:
			var (
				body GroupsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aggregate", "groups", err)
			}
			res := NewGroupsGroupOK(body)
			return res, nil
		case http.StatusForbidden:
			var (
				body GroupsPermissionDeniedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aggregate", "groups", err)
			}
			err = ValidateGroupsPermissionDeniedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aggregate", "groups", err)
			}
			return nil, NewGroupsPermissionDenied(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("aggregate", "groups", resp.StatusCode, string(body))
		}
	}
}

// unmarshalAccountResponseToAggregateAccount builds a value of type
// *aggregate.Account from a value of type *AccountResponse.
func unmarshalAccountResponseToAggregateAccount(v *AccountResponse) *aggregate.Account {
	res := &aggregate.Account{
		ID:     v.ID,
		GUID:   v.GUID,
		Name:   *v.Name,
		Login:  *v.Login,
		Email:  *v.Email,
		Status: v.Status,
	}

	return res
}

// unmarshalGroupResponseToAggregateGroup builds a value of type
// *aggregate.Group from a value of type *GroupResponse.
func unmarshalGroupResponseToAggregateGroup(v *GroupResponse) *aggregate.Group {
	res := &aggregate.Group{
		Name: v.Name,
		ID:   v.ID,
	}

	return res
}
