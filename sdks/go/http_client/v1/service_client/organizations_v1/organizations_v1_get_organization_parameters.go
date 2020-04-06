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

package organizations_v1

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

// NewOrganizationsV1GetOrganizationParams creates a new OrganizationsV1GetOrganizationParams object
// with the default values initialized.
func NewOrganizationsV1GetOrganizationParams() *OrganizationsV1GetOrganizationParams {
	var ()
	return &OrganizationsV1GetOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationsV1GetOrganizationParamsWithTimeout creates a new OrganizationsV1GetOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrganizationsV1GetOrganizationParamsWithTimeout(timeout time.Duration) *OrganizationsV1GetOrganizationParams {
	var ()
	return &OrganizationsV1GetOrganizationParams{

		timeout: timeout,
	}
}

// NewOrganizationsV1GetOrganizationParamsWithContext creates a new OrganizationsV1GetOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrganizationsV1GetOrganizationParamsWithContext(ctx context.Context) *OrganizationsV1GetOrganizationParams {
	var ()
	return &OrganizationsV1GetOrganizationParams{

		Context: ctx,
	}
}

// NewOrganizationsV1GetOrganizationParamsWithHTTPClient creates a new OrganizationsV1GetOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrganizationsV1GetOrganizationParamsWithHTTPClient(client *http.Client) *OrganizationsV1GetOrganizationParams {
	var ()
	return &OrganizationsV1GetOrganizationParams{
		HTTPClient: client,
	}
}

/*OrganizationsV1GetOrganizationParams contains all the parameters to send to the API endpoint
for the organizations v1 get organization operation typically these are written to a http.Request
*/
type OrganizationsV1GetOrganizationParams struct {

	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the organizations v1 get organization params
func (o *OrganizationsV1GetOrganizationParams) WithTimeout(timeout time.Duration) *OrganizationsV1GetOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organizations v1 get organization params
func (o *OrganizationsV1GetOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organizations v1 get organization params
func (o *OrganizationsV1GetOrganizationParams) WithContext(ctx context.Context) *OrganizationsV1GetOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organizations v1 get organization params
func (o *OrganizationsV1GetOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organizations v1 get organization params
func (o *OrganizationsV1GetOrganizationParams) WithHTTPClient(client *http.Client) *OrganizationsV1GetOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organizations v1 get organization params
func (o *OrganizationsV1GetOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the organizations v1 get organization params
func (o *OrganizationsV1GetOrganizationParams) WithOwner(owner string) *OrganizationsV1GetOrganizationParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the organizations v1 get organization params
func (o *OrganizationsV1GetOrganizationParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationsV1GetOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
