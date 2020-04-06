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

package project_dashboards_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ProjectDashboardsV1ListProjectDashboardsReader is a Reader for the ProjectDashboardsV1ListProjectDashboards structure.
type ProjectDashboardsV1ListProjectDashboardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectDashboardsV1ListProjectDashboardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectDashboardsV1ListProjectDashboardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewProjectDashboardsV1ListProjectDashboardsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewProjectDashboardsV1ListProjectDashboardsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectDashboardsV1ListProjectDashboardsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewProjectDashboardsV1ListProjectDashboardsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectDashboardsV1ListProjectDashboardsOK creates a ProjectDashboardsV1ListProjectDashboardsOK with default headers values
func NewProjectDashboardsV1ListProjectDashboardsOK() *ProjectDashboardsV1ListProjectDashboardsOK {
	return &ProjectDashboardsV1ListProjectDashboardsOK{}
}

/*ProjectDashboardsV1ListProjectDashboardsOK handles this case with default header values.

A successful response.
*/
type ProjectDashboardsV1ListProjectDashboardsOK struct {
	Payload *service_model.V1ListDashboardsResponse
}

func (o *ProjectDashboardsV1ListProjectDashboardsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/dashboards][%d] projectDashboardsV1ListProjectDashboardsOK  %+v", 200, o.Payload)
}

func (o *ProjectDashboardsV1ListProjectDashboardsOK) GetPayload() *service_model.V1ListDashboardsResponse {
	return o.Payload
}

func (o *ProjectDashboardsV1ListProjectDashboardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListDashboardsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectDashboardsV1ListProjectDashboardsNoContent creates a ProjectDashboardsV1ListProjectDashboardsNoContent with default headers values
func NewProjectDashboardsV1ListProjectDashboardsNoContent() *ProjectDashboardsV1ListProjectDashboardsNoContent {
	return &ProjectDashboardsV1ListProjectDashboardsNoContent{}
}

/*ProjectDashboardsV1ListProjectDashboardsNoContent handles this case with default header values.

No content.
*/
type ProjectDashboardsV1ListProjectDashboardsNoContent struct {
	Payload interface{}
}

func (o *ProjectDashboardsV1ListProjectDashboardsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/dashboards][%d] projectDashboardsV1ListProjectDashboardsNoContent  %+v", 204, o.Payload)
}

func (o *ProjectDashboardsV1ListProjectDashboardsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectDashboardsV1ListProjectDashboardsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectDashboardsV1ListProjectDashboardsForbidden creates a ProjectDashboardsV1ListProjectDashboardsForbidden with default headers values
func NewProjectDashboardsV1ListProjectDashboardsForbidden() *ProjectDashboardsV1ListProjectDashboardsForbidden {
	return &ProjectDashboardsV1ListProjectDashboardsForbidden{}
}

/*ProjectDashboardsV1ListProjectDashboardsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ProjectDashboardsV1ListProjectDashboardsForbidden struct {
	Payload interface{}
}

func (o *ProjectDashboardsV1ListProjectDashboardsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/dashboards][%d] projectDashboardsV1ListProjectDashboardsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectDashboardsV1ListProjectDashboardsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectDashboardsV1ListProjectDashboardsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectDashboardsV1ListProjectDashboardsNotFound creates a ProjectDashboardsV1ListProjectDashboardsNotFound with default headers values
func NewProjectDashboardsV1ListProjectDashboardsNotFound() *ProjectDashboardsV1ListProjectDashboardsNotFound {
	return &ProjectDashboardsV1ListProjectDashboardsNotFound{}
}

/*ProjectDashboardsV1ListProjectDashboardsNotFound handles this case with default header values.

Resource does not exist.
*/
type ProjectDashboardsV1ListProjectDashboardsNotFound struct {
	Payload interface{}
}

func (o *ProjectDashboardsV1ListProjectDashboardsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/dashboards][%d] projectDashboardsV1ListProjectDashboardsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectDashboardsV1ListProjectDashboardsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectDashboardsV1ListProjectDashboardsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectDashboardsV1ListProjectDashboardsDefault creates a ProjectDashboardsV1ListProjectDashboardsDefault with default headers values
func NewProjectDashboardsV1ListProjectDashboardsDefault(code int) *ProjectDashboardsV1ListProjectDashboardsDefault {
	return &ProjectDashboardsV1ListProjectDashboardsDefault{
		_statusCode: code,
	}
}

/*ProjectDashboardsV1ListProjectDashboardsDefault handles this case with default header values.

An unexpected error response
*/
type ProjectDashboardsV1ListProjectDashboardsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the project dashboards v1 list project dashboards default response
func (o *ProjectDashboardsV1ListProjectDashboardsDefault) Code() int {
	return o._statusCode
}

func (o *ProjectDashboardsV1ListProjectDashboardsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/dashboards][%d] ProjectDashboardsV1_ListProjectDashboards default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectDashboardsV1ListProjectDashboardsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ProjectDashboardsV1ListProjectDashboardsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
