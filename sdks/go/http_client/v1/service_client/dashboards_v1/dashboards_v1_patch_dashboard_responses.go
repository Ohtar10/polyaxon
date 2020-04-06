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

package dashboards_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// DashboardsV1PatchDashboardReader is a Reader for the DashboardsV1PatchDashboard structure.
type DashboardsV1PatchDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DashboardsV1PatchDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDashboardsV1PatchDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDashboardsV1PatchDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDashboardsV1PatchDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDashboardsV1PatchDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDashboardsV1PatchDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDashboardsV1PatchDashboardOK creates a DashboardsV1PatchDashboardOK with default headers values
func NewDashboardsV1PatchDashboardOK() *DashboardsV1PatchDashboardOK {
	return &DashboardsV1PatchDashboardOK{}
}

/*DashboardsV1PatchDashboardOK handles this case with default header values.

A successful response.
*/
type DashboardsV1PatchDashboardOK struct {
	Payload *service_model.V1Dashboard
}

func (o *DashboardsV1PatchDashboardOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/dashboards/{dashboard.uuid}][%d] dashboardsV1PatchDashboardOK  %+v", 200, o.Payload)
}

func (o *DashboardsV1PatchDashboardOK) GetPayload() *service_model.V1Dashboard {
	return o.Payload
}

func (o *DashboardsV1PatchDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardsV1PatchDashboardNoContent creates a DashboardsV1PatchDashboardNoContent with default headers values
func NewDashboardsV1PatchDashboardNoContent() *DashboardsV1PatchDashboardNoContent {
	return &DashboardsV1PatchDashboardNoContent{}
}

/*DashboardsV1PatchDashboardNoContent handles this case with default header values.

No content.
*/
type DashboardsV1PatchDashboardNoContent struct {
	Payload interface{}
}

func (o *DashboardsV1PatchDashboardNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/dashboards/{dashboard.uuid}][%d] dashboardsV1PatchDashboardNoContent  %+v", 204, o.Payload)
}

func (o *DashboardsV1PatchDashboardNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DashboardsV1PatchDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardsV1PatchDashboardForbidden creates a DashboardsV1PatchDashboardForbidden with default headers values
func NewDashboardsV1PatchDashboardForbidden() *DashboardsV1PatchDashboardForbidden {
	return &DashboardsV1PatchDashboardForbidden{}
}

/*DashboardsV1PatchDashboardForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type DashboardsV1PatchDashboardForbidden struct {
	Payload interface{}
}

func (o *DashboardsV1PatchDashboardForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/dashboards/{dashboard.uuid}][%d] dashboardsV1PatchDashboardForbidden  %+v", 403, o.Payload)
}

func (o *DashboardsV1PatchDashboardForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DashboardsV1PatchDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardsV1PatchDashboardNotFound creates a DashboardsV1PatchDashboardNotFound with default headers values
func NewDashboardsV1PatchDashboardNotFound() *DashboardsV1PatchDashboardNotFound {
	return &DashboardsV1PatchDashboardNotFound{}
}

/*DashboardsV1PatchDashboardNotFound handles this case with default header values.

Resource does not exist.
*/
type DashboardsV1PatchDashboardNotFound struct {
	Payload interface{}
}

func (o *DashboardsV1PatchDashboardNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/dashboards/{dashboard.uuid}][%d] dashboardsV1PatchDashboardNotFound  %+v", 404, o.Payload)
}

func (o *DashboardsV1PatchDashboardNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DashboardsV1PatchDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardsV1PatchDashboardDefault creates a DashboardsV1PatchDashboardDefault with default headers values
func NewDashboardsV1PatchDashboardDefault(code int) *DashboardsV1PatchDashboardDefault {
	return &DashboardsV1PatchDashboardDefault{
		_statusCode: code,
	}
}

/*DashboardsV1PatchDashboardDefault handles this case with default header values.

An unexpected error response
*/
type DashboardsV1PatchDashboardDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the dashboards v1 patch dashboard default response
func (o *DashboardsV1PatchDashboardDefault) Code() int {
	return o._statusCode
}

func (o *DashboardsV1PatchDashboardDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/dashboards/{dashboard.uuid}][%d] DashboardsV1_PatchDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *DashboardsV1PatchDashboardDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DashboardsV1PatchDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
