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

package versions_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// VersionsV1GetVersionsReader is a Reader for the VersionsV1GetVersions structure.
type VersionsV1GetVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VersionsV1GetVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVersionsV1GetVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewVersionsV1GetVersionsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewVersionsV1GetVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVersionsV1GetVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewVersionsV1GetVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVersionsV1GetVersionsOK creates a VersionsV1GetVersionsOK with default headers values
func NewVersionsV1GetVersionsOK() *VersionsV1GetVersionsOK {
	return &VersionsV1GetVersionsOK{}
}

/*VersionsV1GetVersionsOK handles this case with default header values.

A successful response.
*/
type VersionsV1GetVersionsOK struct {
	Payload *service_model.V1Versions
}

func (o *VersionsV1GetVersionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/version][%d] versionsV1GetVersionsOK  %+v", 200, o.Payload)
}

func (o *VersionsV1GetVersionsOK) GetPayload() *service_model.V1Versions {
	return o.Payload
}

func (o *VersionsV1GetVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Versions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVersionsV1GetVersionsNoContent creates a VersionsV1GetVersionsNoContent with default headers values
func NewVersionsV1GetVersionsNoContent() *VersionsV1GetVersionsNoContent {
	return &VersionsV1GetVersionsNoContent{}
}

/*VersionsV1GetVersionsNoContent handles this case with default header values.

No content.
*/
type VersionsV1GetVersionsNoContent struct {
	Payload interface{}
}

func (o *VersionsV1GetVersionsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/version][%d] versionsV1GetVersionsNoContent  %+v", 204, o.Payload)
}

func (o *VersionsV1GetVersionsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *VersionsV1GetVersionsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVersionsV1GetVersionsForbidden creates a VersionsV1GetVersionsForbidden with default headers values
func NewVersionsV1GetVersionsForbidden() *VersionsV1GetVersionsForbidden {
	return &VersionsV1GetVersionsForbidden{}
}

/*VersionsV1GetVersionsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type VersionsV1GetVersionsForbidden struct {
	Payload interface{}
}

func (o *VersionsV1GetVersionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/version][%d] versionsV1GetVersionsForbidden  %+v", 403, o.Payload)
}

func (o *VersionsV1GetVersionsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *VersionsV1GetVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVersionsV1GetVersionsNotFound creates a VersionsV1GetVersionsNotFound with default headers values
func NewVersionsV1GetVersionsNotFound() *VersionsV1GetVersionsNotFound {
	return &VersionsV1GetVersionsNotFound{}
}

/*VersionsV1GetVersionsNotFound handles this case with default header values.

Resource does not exist.
*/
type VersionsV1GetVersionsNotFound struct {
	Payload interface{}
}

func (o *VersionsV1GetVersionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/version][%d] versionsV1GetVersionsNotFound  %+v", 404, o.Payload)
}

func (o *VersionsV1GetVersionsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *VersionsV1GetVersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVersionsV1GetVersionsDefault creates a VersionsV1GetVersionsDefault with default headers values
func NewVersionsV1GetVersionsDefault(code int) *VersionsV1GetVersionsDefault {
	return &VersionsV1GetVersionsDefault{
		_statusCode: code,
	}
}

/*VersionsV1GetVersionsDefault handles this case with default header values.

An unexpected error response
*/
type VersionsV1GetVersionsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the versions v1 get versions default response
func (o *VersionsV1GetVersionsDefault) Code() int {
	return o._statusCode
}

func (o *VersionsV1GetVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/version][%d] VersionsV1_GetVersions default  %+v", o._statusCode, o.Payload)
}

func (o *VersionsV1GetVersionsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *VersionsV1GetVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
