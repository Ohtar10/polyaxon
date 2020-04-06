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

package runs_v1

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

// NewRunsV1CopyRunParams creates a new RunsV1CopyRunParams object
// with the default values initialized.
func NewRunsV1CopyRunParams() *RunsV1CopyRunParams {
	var ()
	return &RunsV1CopyRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRunsV1CopyRunParamsWithTimeout creates a new RunsV1CopyRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRunsV1CopyRunParamsWithTimeout(timeout time.Duration) *RunsV1CopyRunParams {
	var ()
	return &RunsV1CopyRunParams{

		timeout: timeout,
	}
}

// NewRunsV1CopyRunParamsWithContext creates a new RunsV1CopyRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewRunsV1CopyRunParamsWithContext(ctx context.Context) *RunsV1CopyRunParams {
	var ()
	return &RunsV1CopyRunParams{

		Context: ctx,
	}
}

// NewRunsV1CopyRunParamsWithHTTPClient creates a new RunsV1CopyRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRunsV1CopyRunParamsWithHTTPClient(client *http.Client) *RunsV1CopyRunParams {
	var ()
	return &RunsV1CopyRunParams{
		HTTPClient: client,
	}
}

/*RunsV1CopyRunParams contains all the parameters to send to the API endpoint
for the runs v1 copy run operation typically these are written to a http.Request
*/
type RunsV1CopyRunParams struct {

	/*Body
	  Run object

	*/
	Body *service_model.V1Run
	/*EntityOwner
	  Owner of the namespace

	*/
	EntityOwner string
	/*EntityProject
	  Project

	*/
	EntityProject string
	/*EntityUUID
	  Uuid identifier of the entity

	*/
	EntityUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the runs v1 copy run params
func (o *RunsV1CopyRunParams) WithTimeout(timeout time.Duration) *RunsV1CopyRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the runs v1 copy run params
func (o *RunsV1CopyRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the runs v1 copy run params
func (o *RunsV1CopyRunParams) WithContext(ctx context.Context) *RunsV1CopyRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the runs v1 copy run params
func (o *RunsV1CopyRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the runs v1 copy run params
func (o *RunsV1CopyRunParams) WithHTTPClient(client *http.Client) *RunsV1CopyRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the runs v1 copy run params
func (o *RunsV1CopyRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the runs v1 copy run params
func (o *RunsV1CopyRunParams) WithBody(body *service_model.V1Run) *RunsV1CopyRunParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the runs v1 copy run params
func (o *RunsV1CopyRunParams) SetBody(body *service_model.V1Run) {
	o.Body = body
}

// WithEntityOwner adds the entityOwner to the runs v1 copy run params
func (o *RunsV1CopyRunParams) WithEntityOwner(entityOwner string) *RunsV1CopyRunParams {
	o.SetEntityOwner(entityOwner)
	return o
}

// SetEntityOwner adds the entityOwner to the runs v1 copy run params
func (o *RunsV1CopyRunParams) SetEntityOwner(entityOwner string) {
	o.EntityOwner = entityOwner
}

// WithEntityProject adds the entityProject to the runs v1 copy run params
func (o *RunsV1CopyRunParams) WithEntityProject(entityProject string) *RunsV1CopyRunParams {
	o.SetEntityProject(entityProject)
	return o
}

// SetEntityProject adds the entityProject to the runs v1 copy run params
func (o *RunsV1CopyRunParams) SetEntityProject(entityProject string) {
	o.EntityProject = entityProject
}

// WithEntityUUID adds the entityUUID to the runs v1 copy run params
func (o *RunsV1CopyRunParams) WithEntityUUID(entityUUID string) *RunsV1CopyRunParams {
	o.SetEntityUUID(entityUUID)
	return o
}

// SetEntityUUID adds the entityUuid to the runs v1 copy run params
func (o *RunsV1CopyRunParams) SetEntityUUID(entityUUID string) {
	o.EntityUUID = entityUUID
}

// WriteToRequest writes these params to a swagger request
func (o *RunsV1CopyRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param entity.owner
	if err := r.SetPathParam("entity.owner", o.EntityOwner); err != nil {
		return err
	}

	// path param entity.project
	if err := r.SetPathParam("entity.project", o.EntityProject); err != nil {
		return err
	}

	// path param entity.uuid
	if err := r.SetPathParam("entity.uuid", o.EntityUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
