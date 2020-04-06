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

// ProjectSearchesV1CreateProjectSearchReader is a Reader for the ProjectSearchesV1CreateProjectSearch structure.
type ProjectSearchesV1CreateProjectSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectSearchesV1CreateProjectSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectSearchesV1CreateProjectSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewProjectSearchesV1CreateProjectSearchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewProjectSearchesV1CreateProjectSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectSearchesV1CreateProjectSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewProjectSearchesV1CreateProjectSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectSearchesV1CreateProjectSearchOK creates a ProjectSearchesV1CreateProjectSearchOK with default headers values
func NewProjectSearchesV1CreateProjectSearchOK() *ProjectSearchesV1CreateProjectSearchOK {
	return &ProjectSearchesV1CreateProjectSearchOK{}
}

/*ProjectSearchesV1CreateProjectSearchOK handles this case with default header values.

A successful response.
*/
type ProjectSearchesV1CreateProjectSearchOK struct {
	Payload *service_model.V1Search
}

func (o *ProjectSearchesV1CreateProjectSearchOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/searches][%d] projectSearchesV1CreateProjectSearchOK  %+v", 200, o.Payload)
}

func (o *ProjectSearchesV1CreateProjectSearchOK) GetPayload() *service_model.V1Search {
	return o.Payload
}

func (o *ProjectSearchesV1CreateProjectSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Search)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1CreateProjectSearchNoContent creates a ProjectSearchesV1CreateProjectSearchNoContent with default headers values
func NewProjectSearchesV1CreateProjectSearchNoContent() *ProjectSearchesV1CreateProjectSearchNoContent {
	return &ProjectSearchesV1CreateProjectSearchNoContent{}
}

/*ProjectSearchesV1CreateProjectSearchNoContent handles this case with default header values.

No content.
*/
type ProjectSearchesV1CreateProjectSearchNoContent struct {
	Payload interface{}
}

func (o *ProjectSearchesV1CreateProjectSearchNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/searches][%d] projectSearchesV1CreateProjectSearchNoContent  %+v", 204, o.Payload)
}

func (o *ProjectSearchesV1CreateProjectSearchNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectSearchesV1CreateProjectSearchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1CreateProjectSearchForbidden creates a ProjectSearchesV1CreateProjectSearchForbidden with default headers values
func NewProjectSearchesV1CreateProjectSearchForbidden() *ProjectSearchesV1CreateProjectSearchForbidden {
	return &ProjectSearchesV1CreateProjectSearchForbidden{}
}

/*ProjectSearchesV1CreateProjectSearchForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ProjectSearchesV1CreateProjectSearchForbidden struct {
	Payload interface{}
}

func (o *ProjectSearchesV1CreateProjectSearchForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/searches][%d] projectSearchesV1CreateProjectSearchForbidden  %+v", 403, o.Payload)
}

func (o *ProjectSearchesV1CreateProjectSearchForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectSearchesV1CreateProjectSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1CreateProjectSearchNotFound creates a ProjectSearchesV1CreateProjectSearchNotFound with default headers values
func NewProjectSearchesV1CreateProjectSearchNotFound() *ProjectSearchesV1CreateProjectSearchNotFound {
	return &ProjectSearchesV1CreateProjectSearchNotFound{}
}

/*ProjectSearchesV1CreateProjectSearchNotFound handles this case with default header values.

Resource does not exist.
*/
type ProjectSearchesV1CreateProjectSearchNotFound struct {
	Payload interface{}
}

func (o *ProjectSearchesV1CreateProjectSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/searches][%d] projectSearchesV1CreateProjectSearchNotFound  %+v", 404, o.Payload)
}

func (o *ProjectSearchesV1CreateProjectSearchNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectSearchesV1CreateProjectSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1CreateProjectSearchDefault creates a ProjectSearchesV1CreateProjectSearchDefault with default headers values
func NewProjectSearchesV1CreateProjectSearchDefault(code int) *ProjectSearchesV1CreateProjectSearchDefault {
	return &ProjectSearchesV1CreateProjectSearchDefault{
		_statusCode: code,
	}
}

/*ProjectSearchesV1CreateProjectSearchDefault handles this case with default header values.

An unexpected error response
*/
type ProjectSearchesV1CreateProjectSearchDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the project searches v1 create project search default response
func (o *ProjectSearchesV1CreateProjectSearchDefault) Code() int {
	return o._statusCode
}

func (o *ProjectSearchesV1CreateProjectSearchDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/searches][%d] ProjectSearchesV1_CreateProjectSearch default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectSearchesV1CreateProjectSearchDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ProjectSearchesV1CreateProjectSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
