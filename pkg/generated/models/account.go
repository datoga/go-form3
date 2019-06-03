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

// Account account
// swagger:model Account
type Account struct {

	// attributes
	// Required: true
	Attributes *AccountAttributes `json:"attributes"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// Name of the resource type
	// Pattern: ^[A-Za-z]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func AccountWithDefaults(defaults client.Defaults) *Account {
	return &Account{

		Attributes: AccountAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUIDPtr("Account", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("Account", "organisation_id"),

		Type: defaults.GetString("Account", "type"),

		Version: defaults.GetInt64Ptr("Account", "version"),
	}
}

func (m *Account) WithAttributes(attributes AccountAttributes) *Account {

	m.Attributes = &attributes

	return m
}

func (m *Account) WithoutAttributes() *Account {
	m.Attributes = nil
	return m
}

func (m *Account) WithID(id strfmt.UUID) *Account {

	m.ID = &id

	return m
}

func (m *Account) WithoutID() *Account {
	m.ID = nil
	return m
}

func (m *Account) WithOrganisationID(organisationID strfmt.UUID) *Account {

	m.OrganisationID = &organisationID

	return m
}

func (m *Account) WithoutOrganisationID() *Account {
	m.OrganisationID = nil
	return m
}

func (m *Account) WithType(typeVar string) *Account {

	m.Type = typeVar

	return m
}

func (m *Account) WithVersion(version int64) *Account {

	m.Version = &version

	return m
}

func (m *Account) WithoutVersion() *Account {
	m.Version = nil
	return m
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
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

func (m *Account) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *Account) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Account) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Account) UnmarshalBinary(b []byte) error {
	var res Account
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *Account) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
