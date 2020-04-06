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

// NewRunsV1GetRunsArtifactsLineageParams creates a new RunsV1GetRunsArtifactsLineageParams object
// with the default values initialized.
func NewRunsV1GetRunsArtifactsLineageParams() *RunsV1GetRunsArtifactsLineageParams {
	var ()
	return &RunsV1GetRunsArtifactsLineageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRunsV1GetRunsArtifactsLineageParamsWithTimeout creates a new RunsV1GetRunsArtifactsLineageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRunsV1GetRunsArtifactsLineageParamsWithTimeout(timeout time.Duration) *RunsV1GetRunsArtifactsLineageParams {
	var ()
	return &RunsV1GetRunsArtifactsLineageParams{

		timeout: timeout,
	}
}

// NewRunsV1GetRunsArtifactsLineageParamsWithContext creates a new RunsV1GetRunsArtifactsLineageParams object
// with the default values initialized, and the ability to set a context for a request
func NewRunsV1GetRunsArtifactsLineageParamsWithContext(ctx context.Context) *RunsV1GetRunsArtifactsLineageParams {
	var ()
	return &RunsV1GetRunsArtifactsLineageParams{

		Context: ctx,
	}
}

// NewRunsV1GetRunsArtifactsLineageParamsWithHTTPClient creates a new RunsV1GetRunsArtifactsLineageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRunsV1GetRunsArtifactsLineageParamsWithHTTPClient(client *http.Client) *RunsV1GetRunsArtifactsLineageParams {
	var ()
	return &RunsV1GetRunsArtifactsLineageParams{
		HTTPClient: client,
	}
}

/*RunsV1GetRunsArtifactsLineageParams contains all the parameters to send to the API endpoint
for the runs v1 get runs artifacts lineage operation typically these are written to a http.Request
*/
type RunsV1GetRunsArtifactsLineageParams struct {

	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project under namesapce

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the runs v1 get runs artifacts lineage params
func (o *RunsV1GetRunsArtifactsLineageParams) WithTimeout(timeout time.Duration) *RunsV1GetRunsArtifactsLineageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the runs v1 get runs artifacts lineage params
func (o *RunsV1GetRunsArtifactsLineageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the runs v1 get runs artifacts lineage params
func (o *RunsV1GetRunsArtifactsLineageParams) WithContext(ctx context.Context) *RunsV1GetRunsArtifactsLineageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the runs v1 get runs artifacts lineage params
func (o *RunsV1GetRunsArtifactsLineageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the runs v1 get runs artifacts lineage params
func (o *RunsV1GetRunsArtifactsLineageParams) WithHTTPClient(client *http.Client) *RunsV1GetRunsArtifactsLineageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the runs v1 get runs artifacts lineage params
func (o *RunsV1GetRunsArtifactsLineageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the runs v1 get runs artifacts lineage params
func (o *RunsV1GetRunsArtifactsLineageParams) WithOwner(owner string) *RunsV1GetRunsArtifactsLineageParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the runs v1 get runs artifacts lineage params
func (o *RunsV1GetRunsArtifactsLineageParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the runs v1 get runs artifacts lineage params
func (o *RunsV1GetRunsArtifactsLineageParams) WithProject(project string) *RunsV1GetRunsArtifactsLineageParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the runs v1 get runs artifacts lineage params
func (o *RunsV1GetRunsArtifactsLineageParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *RunsV1GetRunsArtifactsLineageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
