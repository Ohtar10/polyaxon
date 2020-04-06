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
)

// NewRunsV1CollectRunLogsParams creates a new RunsV1CollectRunLogsParams object
// with the default values initialized.
func NewRunsV1CollectRunLogsParams() *RunsV1CollectRunLogsParams {
	var ()
	return &RunsV1CollectRunLogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRunsV1CollectRunLogsParamsWithTimeout creates a new RunsV1CollectRunLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRunsV1CollectRunLogsParamsWithTimeout(timeout time.Duration) *RunsV1CollectRunLogsParams {
	var ()
	return &RunsV1CollectRunLogsParams{

		timeout: timeout,
	}
}

// NewRunsV1CollectRunLogsParamsWithContext creates a new RunsV1CollectRunLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewRunsV1CollectRunLogsParamsWithContext(ctx context.Context) *RunsV1CollectRunLogsParams {
	var ()
	return &RunsV1CollectRunLogsParams{

		Context: ctx,
	}
}

// NewRunsV1CollectRunLogsParamsWithHTTPClient creates a new RunsV1CollectRunLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRunsV1CollectRunLogsParamsWithHTTPClient(client *http.Client) *RunsV1CollectRunLogsParams {
	var ()
	return &RunsV1CollectRunLogsParams{
		HTTPClient: client,
	}
}

/*RunsV1CollectRunLogsParams contains all the parameters to send to the API endpoint
for the runs v1 collect run logs operation typically these are written to a http.Request
*/
type RunsV1CollectRunLogsParams struct {

	/*Namespace*/
	Namespace string
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project where the run will be assigned

	*/
	Project string
	/*UUID
	  Uuid identifier of the entity

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the runs v1 collect run logs params
func (o *RunsV1CollectRunLogsParams) WithTimeout(timeout time.Duration) *RunsV1CollectRunLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the runs v1 collect run logs params
func (o *RunsV1CollectRunLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the runs v1 collect run logs params
func (o *RunsV1CollectRunLogsParams) WithContext(ctx context.Context) *RunsV1CollectRunLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the runs v1 collect run logs params
func (o *RunsV1CollectRunLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the runs v1 collect run logs params
func (o *RunsV1CollectRunLogsParams) WithHTTPClient(client *http.Client) *RunsV1CollectRunLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the runs v1 collect run logs params
func (o *RunsV1CollectRunLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the runs v1 collect run logs params
func (o *RunsV1CollectRunLogsParams) WithNamespace(namespace string) *RunsV1CollectRunLogsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the runs v1 collect run logs params
func (o *RunsV1CollectRunLogsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOwner adds the owner to the runs v1 collect run logs params
func (o *RunsV1CollectRunLogsParams) WithOwner(owner string) *RunsV1CollectRunLogsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the runs v1 collect run logs params
func (o *RunsV1CollectRunLogsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the runs v1 collect run logs params
func (o *RunsV1CollectRunLogsParams) WithProject(project string) *RunsV1CollectRunLogsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the runs v1 collect run logs params
func (o *RunsV1CollectRunLogsParams) SetProject(project string) {
	o.Project = project
}

// WithUUID adds the uuid to the runs v1 collect run logs params
func (o *RunsV1CollectRunLogsParams) WithUUID(uuid string) *RunsV1CollectRunLogsParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the runs v1 collect run logs params
func (o *RunsV1CollectRunLogsParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *RunsV1CollectRunLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
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
