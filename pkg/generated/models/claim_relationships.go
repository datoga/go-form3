// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClaimRelationships claim relationships
// swagger:model ClaimRelationships
type ClaimRelationships struct {

	// claim reversal
	ClaimReversal *ClaimRelationshipsClaimReversal `json:"claim_reversal,omitempty"`

	// claim submission
	ClaimSubmission *ClaimRelationshipsClaimSubmission `json:"claim_submission,omitempty"`
}

func ClaimRelationshipsWithDefaults(defaults client.Defaults) *ClaimRelationships {
	return &ClaimRelationships{

		ClaimReversal: ClaimRelationshipsClaimReversalWithDefaults(defaults),

		ClaimSubmission: ClaimRelationshipsClaimSubmissionWithDefaults(defaults),
	}
}

func (m *ClaimRelationships) WithClaimReversal(claimReversal ClaimRelationshipsClaimReversal) *ClaimRelationships {

	m.ClaimReversal = &claimReversal

	return m
}

func (m *ClaimRelationships) WithoutClaimReversal() *ClaimRelationships {
	m.ClaimReversal = nil
	return m
}

func (m *ClaimRelationships) WithClaimSubmission(claimSubmission ClaimRelationshipsClaimSubmission) *ClaimRelationships {

	m.ClaimSubmission = &claimSubmission

	return m
}

func (m *ClaimRelationships) WithoutClaimSubmission() *ClaimRelationships {
	m.ClaimSubmission = nil
	return m
}

// Validate validates this claim relationships
func (m *ClaimRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClaimReversal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClaimSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimRelationships) validateClaimReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.ClaimReversal) { // not required
		return nil
	}

	if m.ClaimReversal != nil {
		if err := m.ClaimReversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("claim_reversal")
			}
			return err
		}
	}

	return nil
}

func (m *ClaimRelationships) validateClaimSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.ClaimSubmission) { // not required
		return nil
	}

	if m.ClaimSubmission != nil {
		if err := m.ClaimSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("claim_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimRelationships) UnmarshalBinary(b []byte) error {
	var res ClaimRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ClaimRelationshipsClaimReversal claim relationships claim reversal
// swagger:model ClaimRelationshipsClaimReversal
type ClaimRelationshipsClaimReversal struct {

	// data
	Data []*ClaimReversal `json:"data"`
}

func ClaimRelationshipsClaimReversalWithDefaults(defaults client.Defaults) *ClaimRelationshipsClaimReversal {
	return &ClaimRelationshipsClaimReversal{

		Data: make([]*ClaimReversal, 0),
	}
}

func (m *ClaimRelationshipsClaimReversal) WithData(data []*ClaimReversal) *ClaimRelationshipsClaimReversal {

	m.Data = data

	return m
}

// Validate validates this claim relationships claim reversal
func (m *ClaimRelationshipsClaimReversal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimRelationshipsClaimReversal) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("claim_reversal" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimRelationshipsClaimReversal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimRelationshipsClaimReversal) UnmarshalBinary(b []byte) error {
	var res ClaimRelationshipsClaimReversal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimRelationshipsClaimReversal) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ClaimRelationshipsClaimSubmission claim relationships claim submission
// swagger:model ClaimRelationshipsClaimSubmission
type ClaimRelationshipsClaimSubmission struct {

	// data
	Data []*ClaimSubmission `json:"data"`
}

func ClaimRelationshipsClaimSubmissionWithDefaults(defaults client.Defaults) *ClaimRelationshipsClaimSubmission {
	return &ClaimRelationshipsClaimSubmission{

		Data: make([]*ClaimSubmission, 0),
	}
}

func (m *ClaimRelationshipsClaimSubmission) WithData(data []*ClaimSubmission) *ClaimRelationshipsClaimSubmission {

	m.Data = data

	return m
}

// Validate validates this claim relationships claim submission
func (m *ClaimRelationshipsClaimSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimRelationshipsClaimSubmission) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("claim_submission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimRelationshipsClaimSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimRelationshipsClaimSubmission) UnmarshalBinary(b []byte) error {
	var res ClaimRelationshipsClaimSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimRelationshipsClaimSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
