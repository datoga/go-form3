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

// AccountAttributesIdentificationAddress account attributes identification address
// swagger:model AccountAttributesIdentificationAddress
type AccountAttributesIdentificationAddress struct {

	// city
	// Max Length: 35
	// Min Length: 1
	City string `json:"city,omitempty"`

	// country
	// Pattern: ^[A-Z]{2}$
	Country string `json:"country,omitempty"`

	// street number
	// Max Length: 140
	// Min Length: 1
	StreetNumber string `json:"street_number,omitempty"`
}

func AccountAttributesIdentificationAddressWithDefaults(defaults client.Defaults) *AccountAttributesIdentificationAddress {
	return &AccountAttributesIdentificationAddress{

		City: defaults.GetString("AccountAttributesIdentificationAddress", "city"),

		Country: defaults.GetString("AccountAttributesIdentificationAddress", "country"),

		StreetNumber: defaults.GetString("AccountAttributesIdentificationAddress", "street_number"),
	}
}

func (m *AccountAttributesIdentificationAddress) WithCity(city string) *AccountAttributesIdentificationAddress {

	m.City = city

	return m
}

func (m *AccountAttributesIdentificationAddress) WithCountry(country string) *AccountAttributesIdentificationAddress {

	m.Country = country

	return m
}

func (m *AccountAttributesIdentificationAddress) WithStreetNumber(streetNumber string) *AccountAttributesIdentificationAddress {

	m.StreetNumber = streetNumber

	return m
}

// Validate validates this account attributes identification address
func (m *AccountAttributesIdentificationAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreetNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountAttributesIdentificationAddress) validateCity(formats strfmt.Registry) error {

	if swag.IsZero(m.City) { // not required
		return nil
	}

	if err := validate.MinLength("city", "body", string(m.City), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("city", "body", string(m.City), 35); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributesIdentificationAddress) validateCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.Country) { // not required
		return nil
	}

	if err := validate.Pattern("country", "body", string(m.Country), `^[A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributesIdentificationAddress) validateStreetNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.StreetNumber) { // not required
		return nil
	}

	if err := validate.MinLength("street_number", "body", string(m.StreetNumber), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("street_number", "body", string(m.StreetNumber), 140); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAttributesIdentificationAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAttributesIdentificationAddress) UnmarshalBinary(b []byte) error {
	var res AccountAttributesIdentificationAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountAttributesIdentificationAddress) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
