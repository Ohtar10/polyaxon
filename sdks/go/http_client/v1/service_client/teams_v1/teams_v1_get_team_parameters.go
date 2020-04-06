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

package teams_v1

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

// NewTeamsV1GetTeamParams creates a new TeamsV1GetTeamParams object
// with the default values initialized.
func NewTeamsV1GetTeamParams() *TeamsV1GetTeamParams {
	var ()
	return &TeamsV1GetTeamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTeamsV1GetTeamParamsWithTimeout creates a new TeamsV1GetTeamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTeamsV1GetTeamParamsWithTimeout(timeout time.Duration) *TeamsV1GetTeamParams {
	var ()
	return &TeamsV1GetTeamParams{

		timeout: timeout,
	}
}

// NewTeamsV1GetTeamParamsWithContext creates a new TeamsV1GetTeamParams object
// with the default values initialized, and the ability to set a context for a request
func NewTeamsV1GetTeamParamsWithContext(ctx context.Context) *TeamsV1GetTeamParams {
	var ()
	return &TeamsV1GetTeamParams{

		Context: ctx,
	}
}

// NewTeamsV1GetTeamParamsWithHTTPClient creates a new TeamsV1GetTeamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTeamsV1GetTeamParamsWithHTTPClient(client *http.Client) *TeamsV1GetTeamParams {
	var ()
	return &TeamsV1GetTeamParams{
		HTTPClient: client,
	}
}

/*TeamsV1GetTeamParams contains all the parameters to send to the API endpoint
for the teams v1 get team operation typically these are written to a http.Request
*/
type TeamsV1GetTeamParams struct {

	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Team
	  Team under namesapce

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the teams v1 get team params
func (o *TeamsV1GetTeamParams) WithTimeout(timeout time.Duration) *TeamsV1GetTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the teams v1 get team params
func (o *TeamsV1GetTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the teams v1 get team params
func (o *TeamsV1GetTeamParams) WithContext(ctx context.Context) *TeamsV1GetTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the teams v1 get team params
func (o *TeamsV1GetTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the teams v1 get team params
func (o *TeamsV1GetTeamParams) WithHTTPClient(client *http.Client) *TeamsV1GetTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the teams v1 get team params
func (o *TeamsV1GetTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the teams v1 get team params
func (o *TeamsV1GetTeamParams) WithOwner(owner string) *TeamsV1GetTeamParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the teams v1 get team params
func (o *TeamsV1GetTeamParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithTeam adds the team to the teams v1 get team params
func (o *TeamsV1GetTeamParams) WithTeam(team string) *TeamsV1GetTeamParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the teams v1 get team params
func (o *TeamsV1GetTeamParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *TeamsV1GetTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
