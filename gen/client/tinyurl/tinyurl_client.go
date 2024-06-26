// Code generated by go-swagger; DO NOT EDIT.

package tinyurl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tinyurl API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tinyurl API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	V1DomainsShorterned(params *V1DomainsShorternedParams, opts ...ClientOption) (*V1DomainsShorternedOK, error)

	V1TinyurlPost(params *V1TinyurlPostParams, opts ...ClientOption) (*V1TinyurlPostOK, error)

	V1TinyurlRedirect(params *V1TinyurlRedirectParams, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
V1DomainsShorterned domains shortened the most

Domains shortened the most
*/
func (a *Client) V1DomainsShorterned(params *V1DomainsShorternedParams, opts ...ClientOption) (*V1DomainsShorternedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1DomainsShorternedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "V1DomainsShorterned",
		Method:             "GET",
		PathPattern:        "/v1/maxdomainsabbrev",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1DomainsShorternedReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*V1DomainsShorternedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*V1DomainsShorternedDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
V1TinyurlPost URLs shortener

URL Shortener
*/
func (a *Client) V1TinyurlPost(params *V1TinyurlPostParams, opts ...ClientOption) (*V1TinyurlPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1TinyurlPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "V1TinyurlPost",
		Method:             "POST",
		PathPattern:        "/v1/tinyurl",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1TinyurlPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*V1TinyurlPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*V1TinyurlPostDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
V1TinyurlRedirect redirects to original URL

Redirect to original URL
*/
func (a *Client) V1TinyurlRedirect(params *V1TinyurlRedirectParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1TinyurlRedirectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "V1TinyurlRedirect",
		Method:             "GET",
		PathPattern:        "/v1/tinyurl/{tinyurl}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1TinyurlRedirectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
