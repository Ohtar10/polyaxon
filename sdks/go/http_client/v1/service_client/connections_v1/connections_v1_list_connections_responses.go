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

package connections_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ConnectionsV1ListConnectionsReader is a Reader for the ConnectionsV1ListConnections structure.
type ConnectionsV1ListConnectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectionsV1ListConnectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectionsV1ListConnectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewConnectionsV1ListConnectionsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewConnectionsV1ListConnectionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewConnectionsV1ListConnectionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewConnectionsV1ListConnectionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConnectionsV1ListConnectionsOK creates a ConnectionsV1ListConnectionsOK with default headers values
func NewConnectionsV1ListConnectionsOK() *ConnectionsV1ListConnectionsOK {
	return &ConnectionsV1ListConnectionsOK{}
}

/*ConnectionsV1ListConnectionsOK handles this case with default header values.

A successful response.
*/
type ConnectionsV1ListConnectionsOK struct {
	Payload *service_model.V1ListConnectionsResponse
}

func (o *ConnectionsV1ListConnectionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections][%d] connectionsV1ListConnectionsOK  %+v", 200, o.Payload)
}

func (o *ConnectionsV1ListConnectionsOK) GetPayload() *service_model.V1ListConnectionsResponse {
	return o.Payload
}

func (o *ConnectionsV1ListConnectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListConnectionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectionsV1ListConnectionsNoContent creates a ConnectionsV1ListConnectionsNoContent with default headers values
func NewConnectionsV1ListConnectionsNoContent() *ConnectionsV1ListConnectionsNoContent {
	return &ConnectionsV1ListConnectionsNoContent{}
}

/*ConnectionsV1ListConnectionsNoContent handles this case with default header values.

No content.
*/
type ConnectionsV1ListConnectionsNoContent struct {
	Payload interface{}
}

func (o *ConnectionsV1ListConnectionsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections][%d] connectionsV1ListConnectionsNoContent  %+v", 204, o.Payload)
}

func (o *ConnectionsV1ListConnectionsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ConnectionsV1ListConnectionsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectionsV1ListConnectionsForbidden creates a ConnectionsV1ListConnectionsForbidden with default headers values
func NewConnectionsV1ListConnectionsForbidden() *ConnectionsV1ListConnectionsForbidden {
	return &ConnectionsV1ListConnectionsForbidden{}
}

/*ConnectionsV1ListConnectionsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ConnectionsV1ListConnectionsForbidden struct {
	Payload interface{}
}

func (o *ConnectionsV1ListConnectionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections][%d] connectionsV1ListConnectionsForbidden  %+v", 403, o.Payload)
}

func (o *ConnectionsV1ListConnectionsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ConnectionsV1ListConnectionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectionsV1ListConnectionsNotFound creates a ConnectionsV1ListConnectionsNotFound with default headers values
func NewConnectionsV1ListConnectionsNotFound() *ConnectionsV1ListConnectionsNotFound {
	return &ConnectionsV1ListConnectionsNotFound{}
}

/*ConnectionsV1ListConnectionsNotFound handles this case with default header values.

Resource does not exist.
*/
type ConnectionsV1ListConnectionsNotFound struct {
	Payload interface{}
}

func (o *ConnectionsV1ListConnectionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections][%d] connectionsV1ListConnectionsNotFound  %+v", 404, o.Payload)
}

func (o *ConnectionsV1ListConnectionsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ConnectionsV1ListConnectionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectionsV1ListConnectionsDefault creates a ConnectionsV1ListConnectionsDefault with default headers values
func NewConnectionsV1ListConnectionsDefault(code int) *ConnectionsV1ListConnectionsDefault {
	return &ConnectionsV1ListConnectionsDefault{
		_statusCode: code,
	}
}

/*ConnectionsV1ListConnectionsDefault handles this case with default header values.

An unexpected error response
*/
type ConnectionsV1ListConnectionsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the connections v1 list connections default response
func (o *ConnectionsV1ListConnectionsDefault) Code() int {
	return o._statusCode
}

func (o *ConnectionsV1ListConnectionsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections][%d] ConnectionsV1_ListConnections default  %+v", o._statusCode, o.Payload)
}

func (o *ConnectionsV1ListConnectionsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ConnectionsV1ListConnectionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
