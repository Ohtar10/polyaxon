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

// ProjectSearchesV1PromoteProjectSearchReader is a Reader for the ProjectSearchesV1PromoteProjectSearch structure.
type ProjectSearchesV1PromoteProjectSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectSearchesV1PromoteProjectSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectSearchesV1PromoteProjectSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewProjectSearchesV1PromoteProjectSearchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewProjectSearchesV1PromoteProjectSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectSearchesV1PromoteProjectSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewProjectSearchesV1PromoteProjectSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectSearchesV1PromoteProjectSearchOK creates a ProjectSearchesV1PromoteProjectSearchOK with default headers values
func NewProjectSearchesV1PromoteProjectSearchOK() *ProjectSearchesV1PromoteProjectSearchOK {
	return &ProjectSearchesV1PromoteProjectSearchOK{}
}

/*ProjectSearchesV1PromoteProjectSearchOK handles this case with default header values.

A successful response.
*/
type ProjectSearchesV1PromoteProjectSearchOK struct {
}

func (o *ProjectSearchesV1PromoteProjectSearchOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/searches/{uuid}/promote][%d] projectSearchesV1PromoteProjectSearchOK ", 200)
}

func (o *ProjectSearchesV1PromoteProjectSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectSearchesV1PromoteProjectSearchNoContent creates a ProjectSearchesV1PromoteProjectSearchNoContent with default headers values
func NewProjectSearchesV1PromoteProjectSearchNoContent() *ProjectSearchesV1PromoteProjectSearchNoContent {
	return &ProjectSearchesV1PromoteProjectSearchNoContent{}
}

/*ProjectSearchesV1PromoteProjectSearchNoContent handles this case with default header values.

No content.
*/
type ProjectSearchesV1PromoteProjectSearchNoContent struct {
	Payload interface{}
}

func (o *ProjectSearchesV1PromoteProjectSearchNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/searches/{uuid}/promote][%d] projectSearchesV1PromoteProjectSearchNoContent  %+v", 204, o.Payload)
}

func (o *ProjectSearchesV1PromoteProjectSearchNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectSearchesV1PromoteProjectSearchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1PromoteProjectSearchForbidden creates a ProjectSearchesV1PromoteProjectSearchForbidden with default headers values
func NewProjectSearchesV1PromoteProjectSearchForbidden() *ProjectSearchesV1PromoteProjectSearchForbidden {
	return &ProjectSearchesV1PromoteProjectSearchForbidden{}
}

/*ProjectSearchesV1PromoteProjectSearchForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ProjectSearchesV1PromoteProjectSearchForbidden struct {
	Payload interface{}
}

func (o *ProjectSearchesV1PromoteProjectSearchForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/searches/{uuid}/promote][%d] projectSearchesV1PromoteProjectSearchForbidden  %+v", 403, o.Payload)
}

func (o *ProjectSearchesV1PromoteProjectSearchForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectSearchesV1PromoteProjectSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1PromoteProjectSearchNotFound creates a ProjectSearchesV1PromoteProjectSearchNotFound with default headers values
func NewProjectSearchesV1PromoteProjectSearchNotFound() *ProjectSearchesV1PromoteProjectSearchNotFound {
	return &ProjectSearchesV1PromoteProjectSearchNotFound{}
}

/*ProjectSearchesV1PromoteProjectSearchNotFound handles this case with default header values.

Resource does not exist.
*/
type ProjectSearchesV1PromoteProjectSearchNotFound struct {
	Payload interface{}
}

func (o *ProjectSearchesV1PromoteProjectSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/searches/{uuid}/promote][%d] projectSearchesV1PromoteProjectSearchNotFound  %+v", 404, o.Payload)
}

func (o *ProjectSearchesV1PromoteProjectSearchNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectSearchesV1PromoteProjectSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1PromoteProjectSearchDefault creates a ProjectSearchesV1PromoteProjectSearchDefault with default headers values
func NewProjectSearchesV1PromoteProjectSearchDefault(code int) *ProjectSearchesV1PromoteProjectSearchDefault {
	return &ProjectSearchesV1PromoteProjectSearchDefault{
		_statusCode: code,
	}
}

/*ProjectSearchesV1PromoteProjectSearchDefault handles this case with default header values.

An unexpected error response
*/
type ProjectSearchesV1PromoteProjectSearchDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the project searches v1 promote project search default response
func (o *ProjectSearchesV1PromoteProjectSearchDefault) Code() int {
	return o._statusCode
}

func (o *ProjectSearchesV1PromoteProjectSearchDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/searches/{uuid}/promote][%d] ProjectSearchesV1_PromoteProjectSearch default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectSearchesV1PromoteProjectSearchDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ProjectSearchesV1PromoteProjectSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
