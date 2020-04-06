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

package project_searches_v1

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

// NewProjectSearchesV1PatchProjectSearchParams creates a new ProjectSearchesV1PatchProjectSearchParams object
// with the default values initialized.
func NewProjectSearchesV1PatchProjectSearchParams() *ProjectSearchesV1PatchProjectSearchParams {
	var ()
	return &ProjectSearchesV1PatchProjectSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectSearchesV1PatchProjectSearchParamsWithTimeout creates a new ProjectSearchesV1PatchProjectSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectSearchesV1PatchProjectSearchParamsWithTimeout(timeout time.Duration) *ProjectSearchesV1PatchProjectSearchParams {
	var ()
	return &ProjectSearchesV1PatchProjectSearchParams{

		timeout: timeout,
	}
}

// NewProjectSearchesV1PatchProjectSearchParamsWithContext creates a new ProjectSearchesV1PatchProjectSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectSearchesV1PatchProjectSearchParamsWithContext(ctx context.Context) *ProjectSearchesV1PatchProjectSearchParams {
	var ()
	return &ProjectSearchesV1PatchProjectSearchParams{

		Context: ctx,
	}
}

// NewProjectSearchesV1PatchProjectSearchParamsWithHTTPClient creates a new ProjectSearchesV1PatchProjectSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectSearchesV1PatchProjectSearchParamsWithHTTPClient(client *http.Client) *ProjectSearchesV1PatchProjectSearchParams {
	var ()
	return &ProjectSearchesV1PatchProjectSearchParams{
		HTTPClient: client,
	}
}

/*ProjectSearchesV1PatchProjectSearchParams contains all the parameters to send to the API endpoint
for the project searches v1 patch project search operation typically these are written to a http.Request
*/
type ProjectSearchesV1PatchProjectSearchParams struct {

	/*Body
	  Search body

	*/
	Body *service_model.V1Search
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project under namesapce

	*/
	Project string
	/*SearchUUID
	  UUID

	*/
	SearchUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project searches v1 patch project search params
func (o *ProjectSearchesV1PatchProjectSearchParams) WithTimeout(timeout time.Duration) *ProjectSearchesV1PatchProjectSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project searches v1 patch project search params
func (o *ProjectSearchesV1PatchProjectSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project searches v1 patch project search params
func (o *ProjectSearchesV1PatchProjectSearchParams) WithContext(ctx context.Context) *ProjectSearchesV1PatchProjectSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project searches v1 patch project search params
func (o *ProjectSearchesV1PatchProjectSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project searches v1 patch project search params
func (o *ProjectSearchesV1PatchProjectSearchParams) WithHTTPClient(client *http.Client) *ProjectSearchesV1PatchProjectSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project searches v1 patch project search params
func (o *ProjectSearchesV1PatchProjectSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the project searches v1 patch project search params
func (o *ProjectSearchesV1PatchProjectSearchParams) WithBody(body *service_model.V1Search) *ProjectSearchesV1PatchProjectSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the project searches v1 patch project search params
func (o *ProjectSearchesV1PatchProjectSearchParams) SetBody(body *service_model.V1Search) {
	o.Body = body
}

// WithOwner adds the owner to the project searches v1 patch project search params
func (o *ProjectSearchesV1PatchProjectSearchParams) WithOwner(owner string) *ProjectSearchesV1PatchProjectSearchParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the project searches v1 patch project search params
func (o *ProjectSearchesV1PatchProjectSearchParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the project searches v1 patch project search params
func (o *ProjectSearchesV1PatchProjectSearchParams) WithProject(project string) *ProjectSearchesV1PatchProjectSearchParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the project searches v1 patch project search params
func (o *ProjectSearchesV1PatchProjectSearchParams) SetProject(project string) {
	o.Project = project
}

// WithSearchUUID adds the searchUUID to the project searches v1 patch project search params
func (o *ProjectSearchesV1PatchProjectSearchParams) WithSearchUUID(searchUUID string) *ProjectSearchesV1PatchProjectSearchParams {
	o.SetSearchUUID(searchUUID)
	return o
}

// SetSearchUUID adds the searchUuid to the project searches v1 patch project search params
func (o *ProjectSearchesV1PatchProjectSearchParams) SetSearchUUID(searchUUID string) {
	o.SearchUUID = searchUUID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectSearchesV1PatchProjectSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param search.uuid
	if err := r.SetPathParam("search.uuid", o.SearchUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
