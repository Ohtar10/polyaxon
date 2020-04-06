// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package versions_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new versions v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for versions v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	VersionsV1GetLogHandler(params *VersionsV1GetLogHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*VersionsV1GetLogHandlerOK, *VersionsV1GetLogHandlerNoContent, error)

	VersionsV1GetVersions(params *VersionsV1GetVersionsParams, authInfo runtime.ClientAuthInfoWriter) (*VersionsV1GetVersionsOK, *VersionsV1GetVersionsNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  VersionsV1GetLogHandler versions v1 get log handler API
*/
func (a *Client) VersionsV1GetLogHandler(params *VersionsV1GetLogHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*VersionsV1GetLogHandlerOK, *VersionsV1GetLogHandlerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVersionsV1GetLogHandlerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VersionsV1_GetLogHandler",
		Method:             "GET",
		PathPattern:        "/api/v1/log_handler",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VersionsV1GetLogHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *VersionsV1GetLogHandlerOK:
		return value, nil, nil
	case *VersionsV1GetLogHandlerNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VersionsV1GetLogHandlerDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  VersionsV1GetVersions gets current user
*/
func (a *Client) VersionsV1GetVersions(params *VersionsV1GetVersionsParams, authInfo runtime.ClientAuthInfoWriter) (*VersionsV1GetVersionsOK, *VersionsV1GetVersionsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVersionsV1GetVersionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VersionsV1_GetVersions",
		Method:             "GET",
		PathPattern:        "/api/v1/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VersionsV1GetVersionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *VersionsV1GetVersionsOK:
		return value, nil, nil
	case *VersionsV1GetVersionsNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VersionsV1GetVersionsDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
