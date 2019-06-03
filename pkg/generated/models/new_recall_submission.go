// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewRecallSubmission new recall submission
// swagger:model NewRecallSubmission
type NewRecallSubmission struct {

	// attributes
	Attributes interface{} `json:"attributes,omitempty"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *RecallSubmissionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func NewRecallSubmissionWithDefaults(defaults client.Defaults) *NewRecallSubmission {
	return &NewRecallSubmission{

		// TODO Attributes: interface{},

		ID: defaults.GetStrfmtUUIDPtr("NewRecallSubmission", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("NewRecallSubmission", "organisation_id"),

		Relationships: RecallSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("NewRecallSubmission", "type"),

		Version: defaults.GetInt64Ptr("NewRecallSubmission", "version"),
	}
}

func (m *NewRecallSubmission) WithAttributes(attributes interface{}) *NewRecallSubmission {

	m.Attributes = attributes

	return m
}

func (m *NewRecallSubmission) WithID(id strfmt.UUID) *NewRecallSubmission {

	m.ID = &id

	return m
}

func (m *NewRecallSubmission) WithoutID() *NewRecallSubmission {
	m.ID = nil
	return m
}

func (m *NewRecallSubmission) WithOrganisationID(organisationID strfmt.UUID) *NewRecallSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *NewRecallSubmission) WithoutOrganisationID() *NewRecallSubmission {
	m.OrganisationID = nil
	return m
}

func (m *NewRecallSubmission) WithRelationships(relationships RecallSubmissionRelationships) *NewRecallSubmission {

	m.Relationships = &relationships

	return m
}

func (m *NewRecallSubmission) WithoutRelationships() *NewRecallSubmission {
	m.Relationships = nil
	return m
}

func (m *NewRecallSubmission) WithType(typeVar string) *NewRecallSubmission {

	m.Type = typeVar

	return m
}

func (m *NewRecallSubmission) WithVersion(version int64) *NewRecallSubmission {

	m.Version = &version

	return m
}

func (m *NewRecallSubmission) WithoutVersion() *NewRecallSubmission {
	m.Version = nil
	return m
}

// Validate validates this new recall submission
func (m *NewRecallSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewRecallSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewRecallSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewRecallSubmission) validateRelationships(formats strfmt.Registry) error {

	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	if m.Relationships != nil {
		if err := m.Relationships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships")
			}
			return err
		}
	}

	return nil
}

func (m *NewRecallSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *NewRecallSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewRecallSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewRecallSubmission) UnmarshalBinary(b []byte) error {
	var res NewRecallSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewRecallSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
