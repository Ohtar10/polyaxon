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

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1Run Run specification
// swagger:model v1Run
type V1Run struct {

	// Optional if this entity was bookmarked
	Bookmarked bool `json:"bookmarked,omitempty"`

	// Optional if this run was restarted/copied/resumed
	CloningStrategy string `json:"cloning_strategy,omitempty"`

	// Optional content of the entity's spec
	Content string `json:"content,omitempty"`

	// Optional time when the entityt was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Optional if the entity has been deleted
	Deleted bool `json:"deleted,omitempty"`

	// Optional description
	Description string `json:"description,omitempty"`

	// Optional last time the entity was started
	// Format: date-time
	FinishedAt strfmt.DateTime `json:"finished_at,omitempty"`

	// Optional inputs of this entity
	Inputs map[string]string `json:"inputs,omitempty"`

	// Is clone
	IsClone bool `json:"is_clone,omitempty"`

	// Optional flag to tell if this entity is managed by the platform
	IsManaged string `json:"is_managed,omitempty"`

	// Is resume
	IsResume bool `json:"is_resume,omitempty"`

	// Optional run meta info
	MetaInfo *V1RunMetaInfo `json:"meta_info,omitempty"`

	// Optional name
	Name string `json:"name,omitempty"`

	// Optional uuid of the original run
	Original string `json:"original,omitempty"`

	// Optional name of the original run
	OriginalName string `json:"original_name,omitempty"`

	// Optional outputs of this entity
	Outputs map[string]string `json:"outputs,omitempty"`

	// Required name of owner of this entity
	Owner string `json:"owner,omitempty"`

	// Optional uuid of the pipeline
	Pipeline string `json:"pipeline,omitempty"`

	// Optional name of the pipeline
	PipelineName string `json:"pipeline_name,omitempty"`

	// Required project name
	Project string `json:"project,omitempty"`

	// Optional a readme text describing this entity
	Readme string `json:"readme,omitempty"`

	// Optional run environment tracked
	RunEnv map[string]string `json:"run_env,omitempty"`

	// Optional last time the entity was started
	// Format: date-time
	StartedAt strfmt.DateTime `json:"started_at,omitempty"`

	// Optional latest status of this entity
	Status string `json:"status,omitempty"`

	// Optional Tags of this entity
	Tags []string `json:"tags"`

	// Optional last time the entity was updated
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// Required name of user started this entity
	User string `json:"user,omitempty"`

	// UUID
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this v1 run
func (m *V1Run) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetaInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Run) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Run) validateFinishedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.FinishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("finished_at", "body", "date-time", m.FinishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Run) validateMetaInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.MetaInfo) { // not required
		return nil
	}

	if m.MetaInfo != nil {
		if err := m.MetaInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta_info")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Run) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Run) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Run) UnmarshalBinary(b []byte) error {
	var res V1Run
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}