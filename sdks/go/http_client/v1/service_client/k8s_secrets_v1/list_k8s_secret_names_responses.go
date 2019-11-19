// Copyright 2019 Polyaxon, Inc.
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

package k8s_secrets_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListK8sSecretNamesReader is a Reader for the ListK8sSecretNames structure.
type ListK8sSecretNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListK8sSecretNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListK8sSecretNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListK8sSecretNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListK8sSecretNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListK8sSecretNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListK8sSecretNamesOK creates a ListK8sSecretNamesOK with default headers values
func NewListK8sSecretNamesOK() *ListK8sSecretNamesOK {
	return &ListK8sSecretNamesOK{}
}

/*ListK8sSecretNamesOK handles this case with default header values.

A successful response.
*/
type ListK8sSecretNamesOK struct {
	Payload *service_model.V1ListK8sResourcesResponse
}

func (o *ListK8sSecretNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_secrets/names][%d] listK8sSecretNamesOK  %+v", 200, o.Payload)
}

func (o *ListK8sSecretNamesOK) GetPayload() *service_model.V1ListK8sResourcesResponse {
	return o.Payload
}

func (o *ListK8sSecretNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListK8sResourcesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListK8sSecretNamesNoContent creates a ListK8sSecretNamesNoContent with default headers values
func NewListK8sSecretNamesNoContent() *ListK8sSecretNamesNoContent {
	return &ListK8sSecretNamesNoContent{}
}

/*ListK8sSecretNamesNoContent handles this case with default header values.

No content.
*/
type ListK8sSecretNamesNoContent struct {
	Payload interface{}
}

func (o *ListK8sSecretNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_secrets/names][%d] listK8sSecretNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListK8sSecretNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListK8sSecretNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListK8sSecretNamesForbidden creates a ListK8sSecretNamesForbidden with default headers values
func NewListK8sSecretNamesForbidden() *ListK8sSecretNamesForbidden {
	return &ListK8sSecretNamesForbidden{}
}

/*ListK8sSecretNamesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListK8sSecretNamesForbidden struct {
	Payload interface{}
}

func (o *ListK8sSecretNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_secrets/names][%d] listK8sSecretNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListK8sSecretNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListK8sSecretNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListK8sSecretNamesNotFound creates a ListK8sSecretNamesNotFound with default headers values
func NewListK8sSecretNamesNotFound() *ListK8sSecretNamesNotFound {
	return &ListK8sSecretNamesNotFound{}
}

/*ListK8sSecretNamesNotFound handles this case with default header values.

Resource does not exist.
*/
type ListK8sSecretNamesNotFound struct {
	Payload interface{}
}

func (o *ListK8sSecretNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_secrets/names][%d] listK8sSecretNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListK8sSecretNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListK8sSecretNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}