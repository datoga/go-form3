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
	"github.com/go-openapi/validate"
)

// MandateReturnSubmission mandate return submission
// swagger:model MandateReturnSubmission
type MandateReturnSubmission struct {

	// attributes
	Attributes *MandateReturnSubmissionAttributes `json:"attributes,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// organisation id
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// relationships
	Relationships *MandateReturnSubmissionRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func MandateReturnSubmissionWithDefaults(defaults client.Defaults) *MandateReturnSubmission {
	return &MandateReturnSubmission{

		Attributes: MandateReturnSubmissionAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUID("MandateReturnSubmission", "id"),

		OrganisationID: defaults.GetStrfmtUUID("MandateReturnSubmission", "organisation_id"),

		Relationships: MandateReturnSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("MandateReturnSubmission", "type"),

		Version: defaults.GetInt64Ptr("MandateReturnSubmission", "version"),
	}
}

func (m *MandateReturnSubmission) WithAttributes(attributes MandateReturnSubmissionAttributes) *MandateReturnSubmission {

	m.Attributes = &attributes

	return m
}

func (m *MandateReturnSubmission) WithoutAttributes() *MandateReturnSubmission {
	m.Attributes = nil
	return m
}

func (m *MandateReturnSubmission) WithID(id strfmt.UUID) *MandateReturnSubmission {

	m.ID = id

	return m
}

func (m *MandateReturnSubmission) WithOrganisationID(organisationID strfmt.UUID) *MandateReturnSubmission {

	m.OrganisationID = organisationID

	return m
}

func (m *MandateReturnSubmission) WithRelationships(relationships MandateReturnSubmissionRelationships) *MandateReturnSubmission {

	m.Relationships = &relationships

	return m
}

func (m *MandateReturnSubmission) WithoutRelationships() *MandateReturnSubmission {
	m.Relationships = nil
	return m
}

func (m *MandateReturnSubmission) WithType(typeVar string) *MandateReturnSubmission {

	m.Type = typeVar

	return m
}

func (m *MandateReturnSubmission) WithVersion(version int64) *MandateReturnSubmission {

	m.Version = &version

	return m
}

func (m *MandateReturnSubmission) WithoutVersion() *MandateReturnSubmission {
	m.Version = nil
	return m
}

// Validate validates this mandate return submission
func (m *MandateReturnSubmission) Validate(formats strfmt.Registry) error {
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

func (m *MandateReturnSubmission) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
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

func (m *MandateReturnSubmission) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateReturnSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateReturnSubmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *MandateReturnSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MandateReturnSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateReturnSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateReturnSubmission) UnmarshalBinary(b []byte) error {
	var res MandateReturnSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateReturnSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// MandateReturnSubmissionAttributes mandate return submission attributes
// swagger:model MandateReturnSubmissionAttributes
type MandateReturnSubmissionAttributes struct {

	// status
	Status MandateReturnSubmissionStatus `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"status_reason,omitempty"`

	// submission datetime
	// Read Only: true
	// Format: date-time
	SubmissionDatetime strfmt.DateTime `json:"submission_datetime,omitempty"`

	// transaction start datetime
	// Read Only: true
	// Format: date-time
	TransactionStartDatetime strfmt.DateTime `json:"transaction_start_datetime,omitempty"`
}

func MandateReturnSubmissionAttributesWithDefaults(defaults client.Defaults) *MandateReturnSubmissionAttributes {
	return &MandateReturnSubmissionAttributes{

		// TODO Status: MandateReturnSubmissionStatus,

		StatusReason: defaults.GetString("MandateReturnSubmissionAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTime("MandateReturnSubmissionAttributes", "submission_datetime"),

		TransactionStartDatetime: defaults.GetStrfmtDateTime("MandateReturnSubmissionAttributes", "transaction_start_datetime"),
	}
}

func (m *MandateReturnSubmissionAttributes) WithStatus(status MandateReturnSubmissionStatus) *MandateReturnSubmissionAttributes {

	m.Status = status

	return m
}

func (m *MandateReturnSubmissionAttributes) WithStatusReason(statusReason string) *MandateReturnSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *MandateReturnSubmissionAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *MandateReturnSubmissionAttributes {

	m.SubmissionDatetime = submissionDatetime

	return m
}

func (m *MandateReturnSubmissionAttributes) WithTransactionStartDatetime(transactionStartDatetime strfmt.DateTime) *MandateReturnSubmissionAttributes {

	m.TransactionStartDatetime = transactionStartDatetime

	return m
}

// Validate validates this mandate return submission attributes
func (m *MandateReturnSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionStartDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateReturnSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "status")
		}
		return err
	}

	return nil
}

func (m *MandateReturnSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateReturnSubmissionAttributes) validateTransactionStartDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionStartDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"transaction_start_datetime", "body", "date-time", m.TransactionStartDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateReturnSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateReturnSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res MandateReturnSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateReturnSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// MandateReturnSubmissionRelationships mandate return submission relationships
// swagger:model MandateReturnSubmissionRelationships
type MandateReturnSubmissionRelationships struct {

	// mandate
	Mandate *MandateReturnSubmissionRelationshipsMandate `json:"mandate,omitempty"`

	// mandate return
	MandateReturn *MandateReturnSubmissionRelationshipsMandateReturn `json:"mandate_return,omitempty"`
}

func MandateReturnSubmissionRelationshipsWithDefaults(defaults client.Defaults) *MandateReturnSubmissionRelationships {
	return &MandateReturnSubmissionRelationships{

		Mandate: MandateReturnSubmissionRelationshipsMandateWithDefaults(defaults),

		MandateReturn: MandateReturnSubmissionRelationshipsMandateReturnWithDefaults(defaults),
	}
}

func (m *MandateReturnSubmissionRelationships) WithMandate(mandate MandateReturnSubmissionRelationshipsMandate) *MandateReturnSubmissionRelationships {

	m.Mandate = &mandate

	return m
}

func (m *MandateReturnSubmissionRelationships) WithoutMandate() *MandateReturnSubmissionRelationships {
	m.Mandate = nil
	return m
}

func (m *MandateReturnSubmissionRelationships) WithMandateReturn(mandateReturn MandateReturnSubmissionRelationshipsMandateReturn) *MandateReturnSubmissionRelationships {

	m.MandateReturn = &mandateReturn

	return m
}

func (m *MandateReturnSubmissionRelationships) WithoutMandateReturn() *MandateReturnSubmissionRelationships {
	m.MandateReturn = nil
	return m
}

// Validate validates this mandate return submission relationships
func (m *MandateReturnSubmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMandate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMandateReturn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateReturnSubmissionRelationships) validateMandate(formats strfmt.Registry) error {

	if swag.IsZero(m.Mandate) { // not required
		return nil
	}

	if m.Mandate != nil {
		if err := m.Mandate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "mandate")
			}
			return err
		}
	}

	return nil
}

func (m *MandateReturnSubmissionRelationships) validateMandateReturn(formats strfmt.Registry) error {

	if swag.IsZero(m.MandateReturn) { // not required
		return nil
	}

	if m.MandateReturn != nil {
		if err := m.MandateReturn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "mandate_return")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateReturnSubmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateReturnSubmissionRelationships) UnmarshalBinary(b []byte) error {
	var res MandateReturnSubmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateReturnSubmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// MandateReturnSubmissionRelationshipsMandate mandate return submission relationships mandate
// swagger:model MandateReturnSubmissionRelationshipsMandate
type MandateReturnSubmissionRelationshipsMandate struct {

	// data
	Data []*Mandate `json:"data"`
}

func MandateReturnSubmissionRelationshipsMandateWithDefaults(defaults client.Defaults) *MandateReturnSubmissionRelationshipsMandate {
	return &MandateReturnSubmissionRelationshipsMandate{

		Data: make([]*Mandate, 0),
	}
}

func (m *MandateReturnSubmissionRelationshipsMandate) WithData(data []*Mandate) *MandateReturnSubmissionRelationshipsMandate {

	m.Data = data

	return m
}

// Validate validates this mandate return submission relationships mandate
func (m *MandateReturnSubmissionRelationshipsMandate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateReturnSubmissionRelationshipsMandate) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "mandate" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateReturnSubmissionRelationshipsMandate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateReturnSubmissionRelationshipsMandate) UnmarshalBinary(b []byte) error {
	var res MandateReturnSubmissionRelationshipsMandate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateReturnSubmissionRelationshipsMandate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// MandateReturnSubmissionRelationshipsMandateReturn mandate return submission relationships mandate return
// swagger:model MandateReturnSubmissionRelationshipsMandateReturn
type MandateReturnSubmissionRelationshipsMandateReturn struct {

	// data
	Data []*MandateReturn `json:"data"`
}

func MandateReturnSubmissionRelationshipsMandateReturnWithDefaults(defaults client.Defaults) *MandateReturnSubmissionRelationshipsMandateReturn {
	return &MandateReturnSubmissionRelationshipsMandateReturn{

		Data: make([]*MandateReturn, 0),
	}
}

func (m *MandateReturnSubmissionRelationshipsMandateReturn) WithData(data []*MandateReturn) *MandateReturnSubmissionRelationshipsMandateReturn {

	m.Data = data

	return m
}

// Validate validates this mandate return submission relationships mandate return
func (m *MandateReturnSubmissionRelationshipsMandateReturn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateReturnSubmissionRelationshipsMandateReturn) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "mandate_return" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateReturnSubmissionRelationshipsMandateReturn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateReturnSubmissionRelationshipsMandateReturn) UnmarshalBinary(b []byte) error {
	var res MandateReturnSubmissionRelationshipsMandateReturn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateReturnSubmissionRelationshipsMandateReturn) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
