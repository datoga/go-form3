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

// ReportAdmissionDetailsResponse report admission details response
// swagger:model ReportAdmissionDetailsResponse
type ReportAdmissionDetailsResponse struct {

	// data
	// Required: true
	Data *ReportAdmission `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func ReportAdmissionDetailsResponseWithDefaults(defaults client.Defaults) *ReportAdmissionDetailsResponse {
	return &ReportAdmissionDetailsResponse{

		Data: ReportAdmissionWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *ReportAdmissionDetailsResponse) WithData(data ReportAdmission) *ReportAdmissionDetailsResponse {

	m.Data = &data

	return m
}

func (m *ReportAdmissionDetailsResponse) WithoutData() *ReportAdmissionDetailsResponse {
	m.Data = nil
	return m
}

func (m *ReportAdmissionDetailsResponse) WithLinks(links Links) *ReportAdmissionDetailsResponse {

	m.Links = &links

	return m
}

func (m *ReportAdmissionDetailsResponse) WithoutLinks() *ReportAdmissionDetailsResponse {
	m.Links = nil
	return m
}

// Validate validates this report admission details response
func (m *ReportAdmissionDetailsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportAdmissionDetailsResponse) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *ReportAdmissionDetailsResponse) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportAdmissionDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportAdmissionDetailsResponse) UnmarshalBinary(b []byte) error {
	var res ReportAdmissionDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReportAdmissionDetailsResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
