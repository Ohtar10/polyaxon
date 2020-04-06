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

package connections_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewConnectionsV1CreateConnectionParams creates a new ConnectionsV1CreateConnectionParams object
// with the default values initialized.
func NewConnectionsV1CreateConnectionParams() *ConnectionsV1CreateConnectionParams {
	var ()
	return &ConnectionsV1CreateConnectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConnectionsV1CreateConnectionParamsWithTimeout creates a new ConnectionsV1CreateConnectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConnectionsV1CreateConnectionParamsWithTimeout(timeout time.Duration) *ConnectionsV1CreateConnectionParams {
	var ()
	return &ConnectionsV1CreateConnectionParams{

		timeout: timeout,
	}
}

// NewConnectionsV1CreateConnectionParamsWithContext creates a new ConnectionsV1CreateConnectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewConnectionsV1CreateConnectionParamsWithContext(ctx context.Context) *ConnectionsV1CreateConnectionParams {
	var ()
	return &ConnectionsV1CreateConnectionParams{

		Context: ctx,
	}
}

// NewConnectionsV1CreateConnectionParamsWithHTTPClient creates a new ConnectionsV1CreateConnectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConnectionsV1CreateConnectionParamsWithHTTPClient(client *http.Client) *ConnectionsV1CreateConnectionParams {
	var ()
	return &ConnectionsV1CreateConnectionParams{
		HTTPClient: client,
	}
}

/*ConnectionsV1CreateConnectionParams contains all the parameters to send to the API endpoint
for the connections v1 create connection operation typically these are written to a http.Request
*/
type ConnectionsV1CreateConnectionParams struct {

	/*Body
	  Connection body

	*/
	Body *service_model.V1ConnectionResponse
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the connections v1 create connection params
func (o *ConnectionsV1CreateConnectionParams) WithTimeout(timeout time.Duration) *ConnectionsV1CreateConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the connections v1 create connection params
func (o *ConnectionsV1CreateConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the connections v1 create connection params
func (o *ConnectionsV1CreateConnectionParams) WithContext(ctx context.Context) *ConnectionsV1CreateConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the connections v1 create connection params
func (o *ConnectionsV1CreateConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the connections v1 create connection params
func (o *ConnectionsV1CreateConnectionParams) WithHTTPClient(client *http.Client) *ConnectionsV1CreateConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the connections v1 create connection params
func (o *ConnectionsV1CreateConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the connections v1 create connection params
func (o *ConnectionsV1CreateConnectionParams) WithBody(body *service_model.V1ConnectionResponse) *ConnectionsV1CreateConnectionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the connections v1 create connection params
func (o *ConnectionsV1CreateConnectionParams) SetBody(body *service_model.V1ConnectionResponse) {
	o.Body = body
}

// WithOwner adds the owner to the connections v1 create connection params
func (o *ConnectionsV1CreateConnectionParams) WithOwner(owner string) *ConnectionsV1CreateConnectionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the connections v1 create connection params
func (o *ConnectionsV1CreateConnectionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *ConnectionsV1CreateConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
