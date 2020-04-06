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

package agents_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// AgentsV1GetAgentStatusesReader is a Reader for the AgentsV1GetAgentStatuses structure.
type AgentsV1GetAgentStatusesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AgentsV1GetAgentStatusesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAgentsV1GetAgentStatusesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewAgentsV1GetAgentStatusesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAgentsV1GetAgentStatusesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAgentsV1GetAgentStatusesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAgentsV1GetAgentStatusesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAgentsV1GetAgentStatusesOK creates a AgentsV1GetAgentStatusesOK with default headers values
func NewAgentsV1GetAgentStatusesOK() *AgentsV1GetAgentStatusesOK {
	return &AgentsV1GetAgentStatusesOK{}
}

/*AgentsV1GetAgentStatusesOK handles this case with default header values.

A successful response.
*/
type AgentsV1GetAgentStatusesOK struct {
	Payload *service_model.V1Status
}

func (o *AgentsV1GetAgentStatusesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}/statuses][%d] agentsV1GetAgentStatusesOK  %+v", 200, o.Payload)
}

func (o *AgentsV1GetAgentStatusesOK) GetPayload() *service_model.V1Status {
	return o.Payload
}

func (o *AgentsV1GetAgentStatusesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAgentsV1GetAgentStatusesNoContent creates a AgentsV1GetAgentStatusesNoContent with default headers values
func NewAgentsV1GetAgentStatusesNoContent() *AgentsV1GetAgentStatusesNoContent {
	return &AgentsV1GetAgentStatusesNoContent{}
}

/*AgentsV1GetAgentStatusesNoContent handles this case with default header values.

No content.
*/
type AgentsV1GetAgentStatusesNoContent struct {
	Payload interface{}
}

func (o *AgentsV1GetAgentStatusesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}/statuses][%d] agentsV1GetAgentStatusesNoContent  %+v", 204, o.Payload)
}

func (o *AgentsV1GetAgentStatusesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *AgentsV1GetAgentStatusesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAgentsV1GetAgentStatusesForbidden creates a AgentsV1GetAgentStatusesForbidden with default headers values
func NewAgentsV1GetAgentStatusesForbidden() *AgentsV1GetAgentStatusesForbidden {
	return &AgentsV1GetAgentStatusesForbidden{}
}

/*AgentsV1GetAgentStatusesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type AgentsV1GetAgentStatusesForbidden struct {
	Payload interface{}
}

func (o *AgentsV1GetAgentStatusesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}/statuses][%d] agentsV1GetAgentStatusesForbidden  %+v", 403, o.Payload)
}

func (o *AgentsV1GetAgentStatusesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *AgentsV1GetAgentStatusesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAgentsV1GetAgentStatusesNotFound creates a AgentsV1GetAgentStatusesNotFound with default headers values
func NewAgentsV1GetAgentStatusesNotFound() *AgentsV1GetAgentStatusesNotFound {
	return &AgentsV1GetAgentStatusesNotFound{}
}

/*AgentsV1GetAgentStatusesNotFound handles this case with default header values.

Resource does not exist.
*/
type AgentsV1GetAgentStatusesNotFound struct {
	Payload interface{}
}

func (o *AgentsV1GetAgentStatusesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}/statuses][%d] agentsV1GetAgentStatusesNotFound  %+v", 404, o.Payload)
}

func (o *AgentsV1GetAgentStatusesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *AgentsV1GetAgentStatusesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAgentsV1GetAgentStatusesDefault creates a AgentsV1GetAgentStatusesDefault with default headers values
func NewAgentsV1GetAgentStatusesDefault(code int) *AgentsV1GetAgentStatusesDefault {
	return &AgentsV1GetAgentStatusesDefault{
		_statusCode: code,
	}
}

/*AgentsV1GetAgentStatusesDefault handles this case with default header values.

An unexpected error response
*/
type AgentsV1GetAgentStatusesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the agents v1 get agent statuses default response
func (o *AgentsV1GetAgentStatusesDefault) Code() int {
	return o._statusCode
}

func (o *AgentsV1GetAgentStatusesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}/statuses][%d] AgentsV1_GetAgentStatuses default  %+v", o._statusCode, o.Payload)
}

func (o *AgentsV1GetAgentStatusesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *AgentsV1GetAgentStatusesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
