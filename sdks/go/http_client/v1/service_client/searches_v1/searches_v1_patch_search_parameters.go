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

package searches_v1

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

// NewSearchesV1PatchSearchParams creates a new SearchesV1PatchSearchParams object
// with the default values initialized.
func NewSearchesV1PatchSearchParams() *SearchesV1PatchSearchParams {
	var ()
	return &SearchesV1PatchSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchesV1PatchSearchParamsWithTimeout creates a new SearchesV1PatchSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchesV1PatchSearchParamsWithTimeout(timeout time.Duration) *SearchesV1PatchSearchParams {
	var ()
	return &SearchesV1PatchSearchParams{

		timeout: timeout,
	}
}

// NewSearchesV1PatchSearchParamsWithContext creates a new SearchesV1PatchSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchesV1PatchSearchParamsWithContext(ctx context.Context) *SearchesV1PatchSearchParams {
	var ()
	return &SearchesV1PatchSearchParams{

		Context: ctx,
	}
}

// NewSearchesV1PatchSearchParamsWithHTTPClient creates a new SearchesV1PatchSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchesV1PatchSearchParamsWithHTTPClient(client *http.Client) *SearchesV1PatchSearchParams {
	var ()
	return &SearchesV1PatchSearchParams{
		HTTPClient: client,
	}
}

/*SearchesV1PatchSearchParams contains all the parameters to send to the API endpoint
for the searches v1 patch search operation typically these are written to a http.Request
*/
type SearchesV1PatchSearchParams struct {

	/*Body
	  Search body

	*/
	Body *service_model.V1Search
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*SearchUUID
	  UUID

	*/
	SearchUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the searches v1 patch search params
func (o *SearchesV1PatchSearchParams) WithTimeout(timeout time.Duration) *SearchesV1PatchSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the searches v1 patch search params
func (o *SearchesV1PatchSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the searches v1 patch search params
func (o *SearchesV1PatchSearchParams) WithContext(ctx context.Context) *SearchesV1PatchSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the searches v1 patch search params
func (o *SearchesV1PatchSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the searches v1 patch search params
func (o *SearchesV1PatchSearchParams) WithHTTPClient(client *http.Client) *SearchesV1PatchSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the searches v1 patch search params
func (o *SearchesV1PatchSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the searches v1 patch search params
func (o *SearchesV1PatchSearchParams) WithBody(body *service_model.V1Search) *SearchesV1PatchSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the searches v1 patch search params
func (o *SearchesV1PatchSearchParams) SetBody(body *service_model.V1Search) {
	o.Body = body
}

// WithOwner adds the owner to the searches v1 patch search params
func (o *SearchesV1PatchSearchParams) WithOwner(owner string) *SearchesV1PatchSearchParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the searches v1 patch search params
func (o *SearchesV1PatchSearchParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithSearchUUID adds the searchUUID to the searches v1 patch search params
func (o *SearchesV1PatchSearchParams) WithSearchUUID(searchUUID string) *SearchesV1PatchSearchParams {
	o.SetSearchUUID(searchUUID)
	return o
}

// SetSearchUUID adds the searchUuid to the searches v1 patch search params
func (o *SearchesV1PatchSearchParams) SetSearchUUID(searchUUID string) {
	o.SearchUUID = searchUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SearchesV1PatchSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param search.uuid
	if err := r.SetPathParam("search.uuid", o.SearchUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
