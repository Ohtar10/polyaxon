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

package projects_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ProjectsV1ListArchivedProjectsReader is a Reader for the ProjectsV1ListArchivedProjects structure.
type ProjectsV1ListArchivedProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsV1ListArchivedProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsV1ListArchivedProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewProjectsV1ListArchivedProjectsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewProjectsV1ListArchivedProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsV1ListArchivedProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewProjectsV1ListArchivedProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectsV1ListArchivedProjectsOK creates a ProjectsV1ListArchivedProjectsOK with default headers values
func NewProjectsV1ListArchivedProjectsOK() *ProjectsV1ListArchivedProjectsOK {
	return &ProjectsV1ListArchivedProjectsOK{}
}

/*ProjectsV1ListArchivedProjectsOK handles this case with default header values.

A successful response.
*/
type ProjectsV1ListArchivedProjectsOK struct {
	Payload *service_model.V1ListProjectsResponse
}

func (o *ProjectsV1ListArchivedProjectsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/projects][%d] projectsV1ListArchivedProjectsOK  %+v", 200, o.Payload)
}

func (o *ProjectsV1ListArchivedProjectsOK) GetPayload() *service_model.V1ListProjectsResponse {
	return o.Payload
}

func (o *ProjectsV1ListArchivedProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListProjectsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsV1ListArchivedProjectsNoContent creates a ProjectsV1ListArchivedProjectsNoContent with default headers values
func NewProjectsV1ListArchivedProjectsNoContent() *ProjectsV1ListArchivedProjectsNoContent {
	return &ProjectsV1ListArchivedProjectsNoContent{}
}

/*ProjectsV1ListArchivedProjectsNoContent handles this case with default header values.

No content.
*/
type ProjectsV1ListArchivedProjectsNoContent struct {
	Payload interface{}
}

func (o *ProjectsV1ListArchivedProjectsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/projects][%d] projectsV1ListArchivedProjectsNoContent  %+v", 204, o.Payload)
}

func (o *ProjectsV1ListArchivedProjectsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsV1ListArchivedProjectsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsV1ListArchivedProjectsForbidden creates a ProjectsV1ListArchivedProjectsForbidden with default headers values
func NewProjectsV1ListArchivedProjectsForbidden() *ProjectsV1ListArchivedProjectsForbidden {
	return &ProjectsV1ListArchivedProjectsForbidden{}
}

/*ProjectsV1ListArchivedProjectsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ProjectsV1ListArchivedProjectsForbidden struct {
	Payload interface{}
}

func (o *ProjectsV1ListArchivedProjectsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/projects][%d] projectsV1ListArchivedProjectsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsV1ListArchivedProjectsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsV1ListArchivedProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsV1ListArchivedProjectsNotFound creates a ProjectsV1ListArchivedProjectsNotFound with default headers values
func NewProjectsV1ListArchivedProjectsNotFound() *ProjectsV1ListArchivedProjectsNotFound {
	return &ProjectsV1ListArchivedProjectsNotFound{}
}

/*ProjectsV1ListArchivedProjectsNotFound handles this case with default header values.

Resource does not exist.
*/
type ProjectsV1ListArchivedProjectsNotFound struct {
	Payload interface{}
}

func (o *ProjectsV1ListArchivedProjectsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/projects][%d] projectsV1ListArchivedProjectsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsV1ListArchivedProjectsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsV1ListArchivedProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsV1ListArchivedProjectsDefault creates a ProjectsV1ListArchivedProjectsDefault with default headers values
func NewProjectsV1ListArchivedProjectsDefault(code int) *ProjectsV1ListArchivedProjectsDefault {
	return &ProjectsV1ListArchivedProjectsDefault{
		_statusCode: code,
	}
}

/*ProjectsV1ListArchivedProjectsDefault handles this case with default header values.

An unexpected error response
*/
type ProjectsV1ListArchivedProjectsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the projects v1 list archived projects default response
func (o *ProjectsV1ListArchivedProjectsDefault) Code() int {
	return o._statusCode
}

func (o *ProjectsV1ListArchivedProjectsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/projects][%d] ProjectsV1_ListArchivedProjects default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectsV1ListArchivedProjectsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ProjectsV1ListArchivedProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
