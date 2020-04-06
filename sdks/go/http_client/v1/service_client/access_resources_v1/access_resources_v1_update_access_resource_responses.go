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

package access_resources_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// AccessResourcesV1UpdateAccessResourceReader is a Reader for the AccessResourcesV1UpdateAccessResource structure.
type AccessResourcesV1UpdateAccessResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccessResourcesV1UpdateAccessResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccessResourcesV1UpdateAccessResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewAccessResourcesV1UpdateAccessResourceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAccessResourcesV1UpdateAccessResourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccessResourcesV1UpdateAccessResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAccessResourcesV1UpdateAccessResourceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccessResourcesV1UpdateAccessResourceOK creates a AccessResourcesV1UpdateAccessResourceOK with default headers values
func NewAccessResourcesV1UpdateAccessResourceOK() *AccessResourcesV1UpdateAccessResourceOK {
	return &AccessResourcesV1UpdateAccessResourceOK{}
}

/*AccessResourcesV1UpdateAccessResourceOK handles this case with default header values.

A successful response.
*/
type AccessResourcesV1UpdateAccessResourceOK struct {
	Payload *service_model.V1AccessResource
}

func (o *AccessResourcesV1UpdateAccessResourceOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/access_resources/{access_resource.uuid}][%d] accessResourcesV1UpdateAccessResourceOK  %+v", 200, o.Payload)
}

func (o *AccessResourcesV1UpdateAccessResourceOK) GetPayload() *service_model.V1AccessResource {
	return o.Payload
}

func (o *AccessResourcesV1UpdateAccessResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1AccessResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessResourcesV1UpdateAccessResourceNoContent creates a AccessResourcesV1UpdateAccessResourceNoContent with default headers values
func NewAccessResourcesV1UpdateAccessResourceNoContent() *AccessResourcesV1UpdateAccessResourceNoContent {
	return &AccessResourcesV1UpdateAccessResourceNoContent{}
}

/*AccessResourcesV1UpdateAccessResourceNoContent handles this case with default header values.

No content.
*/
type AccessResourcesV1UpdateAccessResourceNoContent struct {
	Payload interface{}
}

func (o *AccessResourcesV1UpdateAccessResourceNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/access_resources/{access_resource.uuid}][%d] accessResourcesV1UpdateAccessResourceNoContent  %+v", 204, o.Payload)
}

func (o *AccessResourcesV1UpdateAccessResourceNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *AccessResourcesV1UpdateAccessResourceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessResourcesV1UpdateAccessResourceForbidden creates a AccessResourcesV1UpdateAccessResourceForbidden with default headers values
func NewAccessResourcesV1UpdateAccessResourceForbidden() *AccessResourcesV1UpdateAccessResourceForbidden {
	return &AccessResourcesV1UpdateAccessResourceForbidden{}
}

/*AccessResourcesV1UpdateAccessResourceForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type AccessResourcesV1UpdateAccessResourceForbidden struct {
	Payload interface{}
}

func (o *AccessResourcesV1UpdateAccessResourceForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/access_resources/{access_resource.uuid}][%d] accessResourcesV1UpdateAccessResourceForbidden  %+v", 403, o.Payload)
}

func (o *AccessResourcesV1UpdateAccessResourceForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *AccessResourcesV1UpdateAccessResourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessResourcesV1UpdateAccessResourceNotFound creates a AccessResourcesV1UpdateAccessResourceNotFound with default headers values
func NewAccessResourcesV1UpdateAccessResourceNotFound() *AccessResourcesV1UpdateAccessResourceNotFound {
	return &AccessResourcesV1UpdateAccessResourceNotFound{}
}

/*AccessResourcesV1UpdateAccessResourceNotFound handles this case with default header values.

Resource does not exist.
*/
type AccessResourcesV1UpdateAccessResourceNotFound struct {
	Payload interface{}
}

func (o *AccessResourcesV1UpdateAccessResourceNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/access_resources/{access_resource.uuid}][%d] accessResourcesV1UpdateAccessResourceNotFound  %+v", 404, o.Payload)
}

func (o *AccessResourcesV1UpdateAccessResourceNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *AccessResourcesV1UpdateAccessResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessResourcesV1UpdateAccessResourceDefault creates a AccessResourcesV1UpdateAccessResourceDefault with default headers values
func NewAccessResourcesV1UpdateAccessResourceDefault(code int) *AccessResourcesV1UpdateAccessResourceDefault {
	return &AccessResourcesV1UpdateAccessResourceDefault{
		_statusCode: code,
	}
}

/*AccessResourcesV1UpdateAccessResourceDefault handles this case with default header values.

An unexpected error response
*/
type AccessResourcesV1UpdateAccessResourceDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the access resources v1 update access resource default response
func (o *AccessResourcesV1UpdateAccessResourceDefault) Code() int {
	return o._statusCode
}

func (o *AccessResourcesV1UpdateAccessResourceDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/access_resources/{access_resource.uuid}][%d] AccessResourcesV1_UpdateAccessResource default  %+v", o._statusCode, o.Payload)
}

func (o *AccessResourcesV1UpdateAccessResourceDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *AccessResourcesV1UpdateAccessResourceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
