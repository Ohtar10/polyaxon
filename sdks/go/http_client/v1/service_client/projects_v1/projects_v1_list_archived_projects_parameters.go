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

package projects_v1

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
	"github.com/go-openapi/swag"
)

// NewProjectsV1ListArchivedProjectsParams creates a new ProjectsV1ListArchivedProjectsParams object
// with the default values initialized.
func NewProjectsV1ListArchivedProjectsParams() *ProjectsV1ListArchivedProjectsParams {
	var ()
	return &ProjectsV1ListArchivedProjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsV1ListArchivedProjectsParamsWithTimeout creates a new ProjectsV1ListArchivedProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsV1ListArchivedProjectsParamsWithTimeout(timeout time.Duration) *ProjectsV1ListArchivedProjectsParams {
	var ()
	return &ProjectsV1ListArchivedProjectsParams{

		timeout: timeout,
	}
}

// NewProjectsV1ListArchivedProjectsParamsWithContext creates a new ProjectsV1ListArchivedProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsV1ListArchivedProjectsParamsWithContext(ctx context.Context) *ProjectsV1ListArchivedProjectsParams {
	var ()
	return &ProjectsV1ListArchivedProjectsParams{

		Context: ctx,
	}
}

// NewProjectsV1ListArchivedProjectsParamsWithHTTPClient creates a new ProjectsV1ListArchivedProjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsV1ListArchivedProjectsParamsWithHTTPClient(client *http.Client) *ProjectsV1ListArchivedProjectsParams {
	var ()
	return &ProjectsV1ListArchivedProjectsParams{
		HTTPClient: client,
	}
}

/*ProjectsV1ListArchivedProjectsParams contains all the parameters to send to the API endpoint
for the projects v1 list archived projects operation typically these are written to a http.Request
*/
type ProjectsV1ListArchivedProjectsParams struct {

	/*Limit
	  Limit size.

	*/
	Limit *int32
	/*Offset
	  Pagination offset.

	*/
	Offset *int32
	/*Query
	  Query filter the search search.

	*/
	Query *string
	/*Sort
	  Sort to order the search.

	*/
	Sort *string
	/*User
	  User

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects v1 list archived projects params
func (o *ProjectsV1ListArchivedProjectsParams) WithTimeout(timeout time.Duration) *ProjectsV1ListArchivedProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects v1 list archived projects params
func (o *ProjectsV1ListArchivedProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects v1 list archived projects params
func (o *ProjectsV1ListArchivedProjectsParams) WithContext(ctx context.Context) *ProjectsV1ListArchivedProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects v1 list archived projects params
func (o *ProjectsV1ListArchivedProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects v1 list archived projects params
func (o *ProjectsV1ListArchivedProjectsParams) WithHTTPClient(client *http.Client) *ProjectsV1ListArchivedProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects v1 list archived projects params
func (o *ProjectsV1ListArchivedProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the projects v1 list archived projects params
func (o *ProjectsV1ListArchivedProjectsParams) WithLimit(limit *int32) *ProjectsV1ListArchivedProjectsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the projects v1 list archived projects params
func (o *ProjectsV1ListArchivedProjectsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the projects v1 list archived projects params
func (o *ProjectsV1ListArchivedProjectsParams) WithOffset(offset *int32) *ProjectsV1ListArchivedProjectsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the projects v1 list archived projects params
func (o *ProjectsV1ListArchivedProjectsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithQuery adds the query to the projects v1 list archived projects params
func (o *ProjectsV1ListArchivedProjectsParams) WithQuery(query *string) *ProjectsV1ListArchivedProjectsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the projects v1 list archived projects params
func (o *ProjectsV1ListArchivedProjectsParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the projects v1 list archived projects params
func (o *ProjectsV1ListArchivedProjectsParams) WithSort(sort *string) *ProjectsV1ListArchivedProjectsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the projects v1 list archived projects params
func (o *ProjectsV1ListArchivedProjectsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithUser adds the user to the projects v1 list archived projects params
func (o *ProjectsV1ListArchivedProjectsParams) WithUser(user string) *ProjectsV1ListArchivedProjectsParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the projects v1 list archived projects params
func (o *ProjectsV1ListArchivedProjectsParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsV1ListArchivedProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Query != nil {

		// query param query
		var qrQuery string
		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {
			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
