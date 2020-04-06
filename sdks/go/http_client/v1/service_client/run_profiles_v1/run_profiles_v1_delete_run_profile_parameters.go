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

package run_profiles_v1

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

// NewRunProfilesV1DeleteRunProfileParams creates a new RunProfilesV1DeleteRunProfileParams object
// with the default values initialized.
func NewRunProfilesV1DeleteRunProfileParams() *RunProfilesV1DeleteRunProfileParams {
	var ()
	return &RunProfilesV1DeleteRunProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRunProfilesV1DeleteRunProfileParamsWithTimeout creates a new RunProfilesV1DeleteRunProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRunProfilesV1DeleteRunProfileParamsWithTimeout(timeout time.Duration) *RunProfilesV1DeleteRunProfileParams {
	var ()
	return &RunProfilesV1DeleteRunProfileParams{

		timeout: timeout,
	}
}

// NewRunProfilesV1DeleteRunProfileParamsWithContext creates a new RunProfilesV1DeleteRunProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewRunProfilesV1DeleteRunProfileParamsWithContext(ctx context.Context) *RunProfilesV1DeleteRunProfileParams {
	var ()
	return &RunProfilesV1DeleteRunProfileParams{

		Context: ctx,
	}
}

// NewRunProfilesV1DeleteRunProfileParamsWithHTTPClient creates a new RunProfilesV1DeleteRunProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRunProfilesV1DeleteRunProfileParamsWithHTTPClient(client *http.Client) *RunProfilesV1DeleteRunProfileParams {
	var ()
	return &RunProfilesV1DeleteRunProfileParams{
		HTTPClient: client,
	}
}

/*RunProfilesV1DeleteRunProfileParams contains all the parameters to send to the API endpoint
for the run profiles v1 delete run profile operation typically these are written to a http.Request
*/
type RunProfilesV1DeleteRunProfileParams struct {

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

// WithTimeout adds the timeout to the run profiles v1 delete run profile params
func (o *RunProfilesV1DeleteRunProfileParams) WithTimeout(timeout time.Duration) *RunProfilesV1DeleteRunProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the run profiles v1 delete run profile params
func (o *RunProfilesV1DeleteRunProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the run profiles v1 delete run profile params
func (o *RunProfilesV1DeleteRunProfileParams) WithContext(ctx context.Context) *RunProfilesV1DeleteRunProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the run profiles v1 delete run profile params
func (o *RunProfilesV1DeleteRunProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the run profiles v1 delete run profile params
func (o *RunProfilesV1DeleteRunProfileParams) WithHTTPClient(client *http.Client) *RunProfilesV1DeleteRunProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the run profiles v1 delete run profile params
func (o *RunProfilesV1DeleteRunProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the run profiles v1 delete run profile params
func (o *RunProfilesV1DeleteRunProfileParams) WithOwner(owner string) *RunProfilesV1DeleteRunProfileParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the run profiles v1 delete run profile params
func (o *RunProfilesV1DeleteRunProfileParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the run profiles v1 delete run profile params
func (o *RunProfilesV1DeleteRunProfileParams) WithUUID(uuid string) *RunProfilesV1DeleteRunProfileParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the run profiles v1 delete run profile params
func (o *RunProfilesV1DeleteRunProfileParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *RunProfilesV1DeleteRunProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
