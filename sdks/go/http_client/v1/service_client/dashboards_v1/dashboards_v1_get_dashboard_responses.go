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

// DashboardsV1GetDashboardReader is a Reader for the DashboardsV1GetDashboard structure.
type DashboardsV1GetDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DashboardsV1GetDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDashboardsV1GetDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDashboardsV1GetDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDashboardsV1GetDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDashboardsV1GetDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDashboardsV1GetDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDashboardsV1GetDashboardOK creates a DashboardsV1GetDashboardOK with default headers values
func NewDashboardsV1GetDashboardOK() *DashboardsV1GetDashboardOK {
	return &DashboardsV1GetDashboardOK{}
}

/*DashboardsV1GetDashboardOK handles this case with default header values.

A successful response.
*/
type DashboardsV1GetDashboardOK struct {
	Payload *service_model.V1Dashboard
}

func (o *DashboardsV1GetDashboardOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/dashboards/{uuid}][%d] dashboardsV1GetDashboardOK  %+v", 200, o.Payload)
}

func (o *DashboardsV1GetDashboardOK) GetPayload() *service_model.V1Dashboard {
	return o.Payload
}

func (o *DashboardsV1GetDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardsV1GetDashboardNoContent creates a DashboardsV1GetDashboardNoContent with default headers values
func NewDashboardsV1GetDashboardNoContent() *DashboardsV1GetDashboardNoContent {
	return &DashboardsV1GetDashboardNoContent{}
}

/*DashboardsV1GetDashboardNoContent handles this case with default header values.

No content.
*/
type DashboardsV1GetDashboardNoContent struct {
	Payload interface{}
}

func (o *DashboardsV1GetDashboardNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/dashboards/{uuid}][%d] dashboardsV1GetDashboardNoContent  %+v", 204, o.Payload)
}

func (o *DashboardsV1GetDashboardNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DashboardsV1GetDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardsV1GetDashboardForbidden creates a DashboardsV1GetDashboardForbidden with default headers values
func NewDashboardsV1GetDashboardForbidden() *DashboardsV1GetDashboardForbidden {
	return &DashboardsV1GetDashboardForbidden{}
}

/*DashboardsV1GetDashboardForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type DashboardsV1GetDashboardForbidden struct {
	Payload interface{}
}

func (o *DashboardsV1GetDashboardForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/dashboards/{uuid}][%d] dashboardsV1GetDashboardForbidden  %+v", 403, o.Payload)
}

func (o *DashboardsV1GetDashboardForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DashboardsV1GetDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardsV1GetDashboardNotFound creates a DashboardsV1GetDashboardNotFound with default headers values
func NewDashboardsV1GetDashboardNotFound() *DashboardsV1GetDashboardNotFound {
	return &DashboardsV1GetDashboardNotFound{}
}

/*DashboardsV1GetDashboardNotFound handles this case with default header values.

Resource does not exist.
*/
type DashboardsV1GetDashboardNotFound struct {
	Payload interface{}
}

func (o *DashboardsV1GetDashboardNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/dashboards/{uuid}][%d] dashboardsV1GetDashboardNotFound  %+v", 404, o.Payload)
}

func (o *DashboardsV1GetDashboardNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DashboardsV1GetDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardsV1GetDashboardDefault creates a DashboardsV1GetDashboardDefault with default headers values
func NewDashboardsV1GetDashboardDefault(code int) *DashboardsV1GetDashboardDefault {
	return &DashboardsV1GetDashboardDefault{
		_statusCode: code,
	}
}

/*DashboardsV1GetDashboardDefault handles this case with default header values.

An unexpected error response
*/
type DashboardsV1GetDashboardDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the dashboards v1 get dashboard default response
func (o *DashboardsV1GetDashboardDefault) Code() int {
	return o._statusCode
}

func (o *DashboardsV1GetDashboardDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/dashboards/{uuid}][%d] DashboardsV1_GetDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *DashboardsV1GetDashboardDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DashboardsV1GetDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
