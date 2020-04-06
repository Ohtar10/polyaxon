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

package project_searches_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ProjectSearchesV1UpdateProjectSearchReader is a Reader for the ProjectSearchesV1UpdateProjectSearch structure.
type ProjectSearchesV1UpdateProjectSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectSearchesV1UpdateProjectSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectSearchesV1UpdateProjectSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewProjectSearchesV1UpdateProjectSearchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewProjectSearchesV1UpdateProjectSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectSearchesV1UpdateProjectSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewProjectSearchesV1UpdateProjectSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectSearchesV1UpdateProjectSearchOK creates a ProjectSearchesV1UpdateProjectSearchOK with default headers values
func NewProjectSearchesV1UpdateProjectSearchOK() *ProjectSearchesV1UpdateProjectSearchOK {
	return &ProjectSearchesV1UpdateProjectSearchOK{}
}

/*ProjectSearchesV1UpdateProjectSearchOK handles this case with default header values.

A successful response.
*/
type ProjectSearchesV1UpdateProjectSearchOK struct {
	Payload *service_model.V1Search
}

func (o *ProjectSearchesV1UpdateProjectSearchOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/searches/{search.uuid}][%d] projectSearchesV1UpdateProjectSearchOK  %+v", 200, o.Payload)
}

func (o *ProjectSearchesV1UpdateProjectSearchOK) GetPayload() *service_model.V1Search {
	return o.Payload
}

func (o *ProjectSearchesV1UpdateProjectSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Search)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1UpdateProjectSearchNoContent creates a ProjectSearchesV1UpdateProjectSearchNoContent with default headers values
func NewProjectSearchesV1UpdateProjectSearchNoContent() *ProjectSearchesV1UpdateProjectSearchNoContent {
	return &ProjectSearchesV1UpdateProjectSearchNoContent{}
}

/*ProjectSearchesV1UpdateProjectSearchNoContent handles this case with default header values.

No content.
*/
type ProjectSearchesV1UpdateProjectSearchNoContent struct {
	Payload interface{}
}

func (o *ProjectSearchesV1UpdateProjectSearchNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/searches/{search.uuid}][%d] projectSearchesV1UpdateProjectSearchNoContent  %+v", 204, o.Payload)
}

func (o *ProjectSearchesV1UpdateProjectSearchNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectSearchesV1UpdateProjectSearchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1UpdateProjectSearchForbidden creates a ProjectSearchesV1UpdateProjectSearchForbidden with default headers values
func NewProjectSearchesV1UpdateProjectSearchForbidden() *ProjectSearchesV1UpdateProjectSearchForbidden {
	return &ProjectSearchesV1UpdateProjectSearchForbidden{}
}

/*ProjectSearchesV1UpdateProjectSearchForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ProjectSearchesV1UpdateProjectSearchForbidden struct {
	Payload interface{}
}

func (o *ProjectSearchesV1UpdateProjectSearchForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/searches/{search.uuid}][%d] projectSearchesV1UpdateProjectSearchForbidden  %+v", 403, o.Payload)
}

func (o *ProjectSearchesV1UpdateProjectSearchForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectSearchesV1UpdateProjectSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1UpdateProjectSearchNotFound creates a ProjectSearchesV1UpdateProjectSearchNotFound with default headers values
func NewProjectSearchesV1UpdateProjectSearchNotFound() *ProjectSearchesV1UpdateProjectSearchNotFound {
	return &ProjectSearchesV1UpdateProjectSearchNotFound{}
}

/*ProjectSearchesV1UpdateProjectSearchNotFound handles this case with default header values.

Resource does not exist.
*/
type ProjectSearchesV1UpdateProjectSearchNotFound struct {
	Payload interface{}
}

func (o *ProjectSearchesV1UpdateProjectSearchNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/searches/{search.uuid}][%d] projectSearchesV1UpdateProjectSearchNotFound  %+v", 404, o.Payload)
}

func (o *ProjectSearchesV1UpdateProjectSearchNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectSearchesV1UpdateProjectSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1UpdateProjectSearchDefault creates a ProjectSearchesV1UpdateProjectSearchDefault with default headers values
func NewProjectSearchesV1UpdateProjectSearchDefault(code int) *ProjectSearchesV1UpdateProjectSearchDefault {
	return &ProjectSearchesV1UpdateProjectSearchDefault{
		_statusCode: code,
	}
}

/*ProjectSearchesV1UpdateProjectSearchDefault handles this case with default header values.

An unexpected error response
*/
type ProjectSearchesV1UpdateProjectSearchDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the project searches v1 update project search default response
func (o *ProjectSearchesV1UpdateProjectSearchDefault) Code() int {
	return o._statusCode
}

func (o *ProjectSearchesV1UpdateProjectSearchDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/searches/{search.uuid}][%d] ProjectSearchesV1_UpdateProjectSearch default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectSearchesV1UpdateProjectSearchDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ProjectSearchesV1UpdateProjectSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
