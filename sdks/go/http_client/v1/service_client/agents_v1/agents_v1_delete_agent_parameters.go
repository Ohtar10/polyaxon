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

package agents_v1

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
)

// NewAgentsV1DeleteAgentParams creates a new AgentsV1DeleteAgentParams object
// with the default values initialized.
func NewAgentsV1DeleteAgentParams() *AgentsV1DeleteAgentParams {
	var ()
	return &AgentsV1DeleteAgentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAgentsV1DeleteAgentParamsWithTimeout creates a new AgentsV1DeleteAgentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAgentsV1DeleteAgentParamsWithTimeout(timeout time.Duration) *AgentsV1DeleteAgentParams {
	var ()
	return &AgentsV1DeleteAgentParams{

		timeout: timeout,
	}
}

// NewAgentsV1DeleteAgentParamsWithContext creates a new AgentsV1DeleteAgentParams object
// with the default values initialized, and the ability to set a context for a request
func NewAgentsV1DeleteAgentParamsWithContext(ctx context.Context) *AgentsV1DeleteAgentParams {
	var ()
	return &AgentsV1DeleteAgentParams{

		Context: ctx,
	}
}

// NewAgentsV1DeleteAgentParamsWithHTTPClient creates a new AgentsV1DeleteAgentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAgentsV1DeleteAgentParamsWithHTTPClient(client *http.Client) *AgentsV1DeleteAgentParams {
	var ()
	return &AgentsV1DeleteAgentParams{
		HTTPClient: client,
	}
}

/*AgentsV1DeleteAgentParams contains all the parameters to send to the API endpoint
for the agents v1 delete agent operation typically these are written to a http.Request
*/
type AgentsV1DeleteAgentParams struct {

	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*UUID
	  Uuid identifier of the entity

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the agents v1 delete agent params
func (o *AgentsV1DeleteAgentParams) WithTimeout(timeout time.Duration) *AgentsV1DeleteAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the agents v1 delete agent params
func (o *AgentsV1DeleteAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the agents v1 delete agent params
func (o *AgentsV1DeleteAgentParams) WithContext(ctx context.Context) *AgentsV1DeleteAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the agents v1 delete agent params
func (o *AgentsV1DeleteAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the agents v1 delete agent params
func (o *AgentsV1DeleteAgentParams) WithHTTPClient(client *http.Client) *AgentsV1DeleteAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the agents v1 delete agent params
func (o *AgentsV1DeleteAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the agents v1 delete agent params
func (o *AgentsV1DeleteAgentParams) WithOwner(owner string) *AgentsV1DeleteAgentParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the agents v1 delete agent params
func (o *AgentsV1DeleteAgentParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the agents v1 delete agent params
func (o *AgentsV1DeleteAgentParams) WithUUID(uuid string) *AgentsV1DeleteAgentParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the agents v1 delete agent params
func (o *AgentsV1DeleteAgentParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *AgentsV1DeleteAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
