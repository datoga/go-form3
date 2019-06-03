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

// DirectDebit direct debit
// swagger:model DirectDebit
type DirectDebit struct {

	// attributes
	// Required: true
	Attributes *DirectDebitAttributes `json:"attributes"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// modified on
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *DirectDebitRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version of the resource
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func DirectDebitWithDefaults(defaults client.Defaults) *DirectDebit {
	return &DirectDebit{

		Attributes: DirectDebitAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("DirectDebit", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("DirectDebit", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("DirectDebit", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("DirectDebit", "organisation_id"),

		Relationships: DirectDebitRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("DirectDebit", "type"),

		Version: defaults.GetInt64Ptr("DirectDebit", "version"),
	}
}

func (m *DirectDebit) WithAttributes(attributes DirectDebitAttributes) *DirectDebit {

	m.Attributes = &attributes

	return m
}

func (m *DirectDebit) WithoutAttributes() *DirectDebit {
	m.Attributes = nil
	return m
}

func (m *DirectDebit) WithCreatedOn(createdOn strfmt.DateTime) *DirectDebit {

	m.CreatedOn = &createdOn

	return m
}

func (m *DirectDebit) WithoutCreatedOn() *DirectDebit {
	m.CreatedOn = nil
	return m
}

func (m *DirectDebit) WithID(id strfmt.UUID) *DirectDebit {

	m.ID = &id

	return m
}

func (m *DirectDebit) WithoutID() *DirectDebit {
	m.ID = nil
	return m
}

func (m *DirectDebit) WithModifiedOn(modifiedOn strfmt.DateTime) *DirectDebit {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *DirectDebit) WithoutModifiedOn() *DirectDebit {
	m.ModifiedOn = nil
	return m
}

func (m *DirectDebit) WithOrganisationID(organisationID strfmt.UUID) *DirectDebit {

	m.OrganisationID = &organisationID

	return m
}

func (m *DirectDebit) WithoutOrganisationID() *DirectDebit {
	m.OrganisationID = nil
	return m
}

func (m *DirectDebit) WithRelationships(relationships DirectDebitRelationships) *DirectDebit {

	m.Relationships = &relationships

	return m
}

func (m *DirectDebit) WithoutRelationships() *DirectDebit {
	m.Relationships = nil
	return m
}

func (m *DirectDebit) WithType(typeVar string) *DirectDebit {

	m.Type = typeVar

	return m
}

func (m *DirectDebit) WithVersion(version int64) *DirectDebit {

	m.Version = &version

	return m
}

func (m *DirectDebit) WithoutVersion() *DirectDebit {
	m.Version = nil
	return m
}

// Validate validates this direct debit
func (m *DirectDebit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
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

func (m *DirectDebit) validateAttributes(formats strfmt.Registry) error {

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

func (m *DirectDebit) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebit) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebit) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebit) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebit) validateRelationships(formats strfmt.Registry) error {

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

func (m *DirectDebit) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebit) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebit) UnmarshalBinary(b []byte) error {
	var res DirectDebit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitAttributes direct debit attributes
// swagger:model DirectDebitAttributes
type DirectDebitAttributes struct {

	// Amount of money moved between the instructing agent and instructed agent
	// Pattern: ^[0-9.]{0,20}$
	Amount string `json:"amount,omitempty"`

	// beneficiary party
	BeneficiaryParty *DirectDebitAttributesBeneficiaryParty `json:"beneficiary_party,omitempty"`

	// Unique identifier for organisations collecting payments
	ClearingID string `json:"clearing_id,omitempty"`

	// Currency of the transaction amount. Currency code as defined in [ISO 4217](http://www.iso.org/iso/home/standards/currency_codes.htm)
	Currency string `json:"currency,omitempty"`

	// debtor party
	DebtorParty *DirectDebitAttributesDebtorParty `json:"debtor_party,omitempty"`

	// Numeric reference field, see scheme specific descriptions for usage
	NumericReference string `json:"numeric_reference,omitempty"`

	// Clearing infrastructure through which the operation instruction is to be processed. Default for given organisation ID is used if left empty. Has to be a valid [scheme identifier](http://draft-api-docs.form3.tech/api.html#enumerations-schemes).
	PaymentScheme string `json:"payment_scheme,omitempty"`

	// Date on which the operation is to be debited from the debtor account. Formatted according to ISO 8601 format: YYYY-MM-DD.
	// Format: date
	ProcessingDate strfmt.Date `json:"processing_date,omitempty"`

	// Payment reference for beneficiary use
	Reference string `json:"reference,omitempty"`

	// The [scheme-specific payment type](#enumerations-scheme-payment-types)
	SchemePaymentType string `json:"scheme_payment_type,omitempty"`

	// Date on which the operation is processed by the scheme. Only used if different from `processing_date`.
	// Format: date
	SchemeProcessingDate strfmt.Date `json:"scheme_processing_date,omitempty"`

	// This reference is used by the receiving party to identify whether the related DDI would have been electronic (AUDDIS) or paper‐based.
	// Enum: [AUDDIS MIGRATING]
	SchemeStatus string `json:"scheme_status,omitempty"`

	// The scheme-specific unique transaction ID. Populated by the scheme.
	UniqueSchemeID string `json:"unique_scheme_id,omitempty"`
}

func DirectDebitAttributesWithDefaults(defaults client.Defaults) *DirectDebitAttributes {
	return &DirectDebitAttributes{

		Amount: defaults.GetString("DirectDebitAttributes", "amount"),

		BeneficiaryParty: DirectDebitAttributesBeneficiaryPartyWithDefaults(defaults),

		ClearingID: defaults.GetString("DirectDebitAttributes", "clearing_id"),

		Currency: defaults.GetString("DirectDebitAttributes", "currency"),

		DebtorParty: DirectDebitAttributesDebtorPartyWithDefaults(defaults),

		NumericReference: defaults.GetString("DirectDebitAttributes", "numeric_reference"),

		PaymentScheme: defaults.GetString("DirectDebitAttributes", "payment_scheme"),

		ProcessingDate: defaults.GetStrfmtDate("DirectDebitAttributes", "processing_date"),

		Reference: defaults.GetString("DirectDebitAttributes", "reference"),

		SchemePaymentType: defaults.GetString("DirectDebitAttributes", "scheme_payment_type"),

		SchemeProcessingDate: defaults.GetStrfmtDate("DirectDebitAttributes", "scheme_processing_date"),

		SchemeStatus: defaults.GetString("DirectDebitAttributes", "scheme_status"),

		UniqueSchemeID: defaults.GetString("DirectDebitAttributes", "unique_scheme_id"),
	}
}

func (m *DirectDebitAttributes) WithAmount(amount string) *DirectDebitAttributes {

	m.Amount = amount

	return m
}

func (m *DirectDebitAttributes) WithBeneficiaryParty(beneficiaryParty DirectDebitAttributesBeneficiaryParty) *DirectDebitAttributes {

	m.BeneficiaryParty = &beneficiaryParty

	return m
}

func (m *DirectDebitAttributes) WithoutBeneficiaryParty() *DirectDebitAttributes {
	m.BeneficiaryParty = nil
	return m
}

func (m *DirectDebitAttributes) WithClearingID(clearingID string) *DirectDebitAttributes {

	m.ClearingID = clearingID

	return m
}

func (m *DirectDebitAttributes) WithCurrency(currency string) *DirectDebitAttributes {

	m.Currency = currency

	return m
}

func (m *DirectDebitAttributes) WithDebtorParty(debtorParty DirectDebitAttributesDebtorParty) *DirectDebitAttributes {

	m.DebtorParty = &debtorParty

	return m
}

func (m *DirectDebitAttributes) WithoutDebtorParty() *DirectDebitAttributes {
	m.DebtorParty = nil
	return m
}

func (m *DirectDebitAttributes) WithNumericReference(numericReference string) *DirectDebitAttributes {

	m.NumericReference = numericReference

	return m
}

func (m *DirectDebitAttributes) WithPaymentScheme(paymentScheme string) *DirectDebitAttributes {

	m.PaymentScheme = paymentScheme

	return m
}

func (m *DirectDebitAttributes) WithProcessingDate(processingDate strfmt.Date) *DirectDebitAttributes {

	m.ProcessingDate = processingDate

	return m
}

func (m *DirectDebitAttributes) WithReference(reference string) *DirectDebitAttributes {

	m.Reference = reference

	return m
}

func (m *DirectDebitAttributes) WithSchemePaymentType(schemePaymentType string) *DirectDebitAttributes {

	m.SchemePaymentType = schemePaymentType

	return m
}

func (m *DirectDebitAttributes) WithSchemeProcessingDate(schemeProcessingDate strfmt.Date) *DirectDebitAttributes {

	m.SchemeProcessingDate = schemeProcessingDate

	return m
}

func (m *DirectDebitAttributes) WithSchemeStatus(schemeStatus string) *DirectDebitAttributes {

	m.SchemeStatus = schemeStatus

	return m
}

func (m *DirectDebitAttributes) WithUniqueSchemeID(uniqueSchemeID string) *DirectDebitAttributes {

	m.UniqueSchemeID = uniqueSchemeID

	return m
}

// Validate validates this direct debit attributes
func (m *DirectDebitAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBeneficiaryParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtorParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemeProcessingDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemeStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitAttributes) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"amount", "body", string(m.Amount), `^[0-9.]{0,20}$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitAttributes) validateBeneficiaryParty(formats strfmt.Registry) error {

	if swag.IsZero(m.BeneficiaryParty) { // not required
		return nil
	}

	if m.BeneficiaryParty != nil {
		if err := m.BeneficiaryParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "beneficiary_party")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitAttributes) validateDebtorParty(formats strfmt.Registry) error {

	if swag.IsZero(m.DebtorParty) { // not required
		return nil
	}

	if m.DebtorParty != nil {
		if err := m.DebtorParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "debtor_party")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitAttributes) validateProcessingDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessingDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"processing_date", "body", "date", m.ProcessingDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitAttributes) validateSchemeProcessingDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SchemeProcessingDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"scheme_processing_date", "body", "date", m.SchemeProcessingDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var directDebitAttributesTypeSchemeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AUDDIS","MIGRATING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		directDebitAttributesTypeSchemeStatusPropEnum = append(directDebitAttributesTypeSchemeStatusPropEnum, v)
	}
}

const (

	// DirectDebitAttributesSchemeStatusAUDDIS captures enum value "AUDDIS"
	DirectDebitAttributesSchemeStatusAUDDIS string = "AUDDIS"

	// DirectDebitAttributesSchemeStatusMIGRATING captures enum value "MIGRATING"
	DirectDebitAttributesSchemeStatusMIGRATING string = "MIGRATING"
)

// prop value enum
func (m *DirectDebitAttributes) validateSchemeStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, directDebitAttributesTypeSchemeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DirectDebitAttributes) validateSchemeStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.SchemeStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateSchemeStatusEnum("attributes"+"."+"scheme_status", "body", m.SchemeStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitAttributes) UnmarshalBinary(b []byte) error {
	var res DirectDebitAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitAttributesBeneficiaryParty direct debit attributes beneficiary party
// swagger:model DirectDebitAttributesBeneficiaryParty
type DirectDebitAttributesBeneficiaryParty struct {

	// Name of beneficiary as given with account
	AccountName string `json:"account_name"`

	// Beneficiary account number
	AccountNumber string `json:"account_number,omitempty"`

	// account number code
	AccountNumberCode AccountNumberCode `json:"account_number_code,omitempty"`

	// The type of the account given with `beneficiary_party.account_number`. Single digit number. Only required if requested by the beneficiary party. Defaults to 0.
	AccountType int64 `json:"account_type,omitempty"`

	// account with
	AccountWith *AccountHoldingEntity `json:"account_with,omitempty"`

	// Beneficiary address
	Address []string `json:"address"`

	// Country of the beneficiary address, ISO 3166 format country code
	Country string `json:"country,omitempty"`
}

func DirectDebitAttributesBeneficiaryPartyWithDefaults(defaults client.Defaults) *DirectDebitAttributesBeneficiaryParty {
	return &DirectDebitAttributesBeneficiaryParty{

		AccountName: defaults.GetString("DirectDebitAttributesBeneficiaryParty", "account_name"),

		AccountNumber: defaults.GetString("DirectDebitAttributesBeneficiaryParty", "account_number"),

		// TODO AccountNumberCode: AccountNumberCode,

		AccountType: defaults.GetInt64("DirectDebitAttributesBeneficiaryParty", "account_type"),

		AccountWith: AccountHoldingEntityWithDefaults(defaults),

		Address: make([]string, 0),

		Country: defaults.GetString("DirectDebitAttributesBeneficiaryParty", "country"),
	}
}

func (m *DirectDebitAttributesBeneficiaryParty) WithAccountName(accountName string) *DirectDebitAttributesBeneficiaryParty {

	m.AccountName = accountName

	return m
}

func (m *DirectDebitAttributesBeneficiaryParty) WithAccountNumber(accountNumber string) *DirectDebitAttributesBeneficiaryParty {

	m.AccountNumber = accountNumber

	return m
}

func (m *DirectDebitAttributesBeneficiaryParty) WithAccountNumberCode(accountNumberCode AccountNumberCode) *DirectDebitAttributesBeneficiaryParty {

	m.AccountNumberCode = accountNumberCode

	return m
}

func (m *DirectDebitAttributesBeneficiaryParty) WithAccountType(accountType int64) *DirectDebitAttributesBeneficiaryParty {

	m.AccountType = accountType

	return m
}

func (m *DirectDebitAttributesBeneficiaryParty) WithAccountWith(accountWith AccountHoldingEntity) *DirectDebitAttributesBeneficiaryParty {

	m.AccountWith = &accountWith

	return m
}

func (m *DirectDebitAttributesBeneficiaryParty) WithoutAccountWith() *DirectDebitAttributesBeneficiaryParty {
	m.AccountWith = nil
	return m
}

func (m *DirectDebitAttributesBeneficiaryParty) WithAddress(address []string) *DirectDebitAttributesBeneficiaryParty {

	m.Address = address

	return m
}

func (m *DirectDebitAttributesBeneficiaryParty) WithCountry(country string) *DirectDebitAttributesBeneficiaryParty {

	m.Country = country

	return m
}

// Validate validates this direct debit attributes beneficiary party
func (m *DirectDebitAttributesBeneficiaryParty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumberCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountWith(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitAttributesBeneficiaryParty) validateAccountNumberCode(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumberCode) { // not required
		return nil
	}

	if err := m.AccountNumberCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "beneficiary_party" + "." + "account_number_code")
		}
		return err
	}

	return nil
}

func (m *DirectDebitAttributesBeneficiaryParty) validateAccountWith(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountWith) { // not required
		return nil
	}

	if m.AccountWith != nil {
		if err := m.AccountWith.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "beneficiary_party" + "." + "account_with")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitAttributesBeneficiaryParty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitAttributesBeneficiaryParty) UnmarshalBinary(b []byte) error {
	var res DirectDebitAttributesBeneficiaryParty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitAttributesBeneficiaryParty) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitAttributesDebtorParty direct debit attributes debtor party
// swagger:model DirectDebitAttributesDebtorParty
type DirectDebitAttributesDebtorParty struct {

	// Name of debtor as given with account
	AccountName string `json:"account_name"`

	// Debtor account number. Allows upper case and numeric characters.
	AccountNumber string `json:"account_number,omitempty"`

	// account number code
	AccountNumberCode AccountNumberCode `json:"account_number_code,omitempty"`

	// account with
	AccountWith *AccountHoldingEntity `json:"account_with,omitempty"`

	// Debtor address
	Address []string `json:"address"`

	// Country of debtor address. ISO 3166 format country code"
	Country string `json:"country,omitempty"`
}

func DirectDebitAttributesDebtorPartyWithDefaults(defaults client.Defaults) *DirectDebitAttributesDebtorParty {
	return &DirectDebitAttributesDebtorParty{

		AccountName: defaults.GetString("DirectDebitAttributesDebtorParty", "account_name"),

		AccountNumber: defaults.GetString("DirectDebitAttributesDebtorParty", "account_number"),

		// TODO AccountNumberCode: AccountNumberCode,

		AccountWith: AccountHoldingEntityWithDefaults(defaults),

		Address: make([]string, 0),

		Country: defaults.GetString("DirectDebitAttributesDebtorParty", "country"),
	}
}

func (m *DirectDebitAttributesDebtorParty) WithAccountName(accountName string) *DirectDebitAttributesDebtorParty {

	m.AccountName = accountName

	return m
}

func (m *DirectDebitAttributesDebtorParty) WithAccountNumber(accountNumber string) *DirectDebitAttributesDebtorParty {

	m.AccountNumber = accountNumber

	return m
}

func (m *DirectDebitAttributesDebtorParty) WithAccountNumberCode(accountNumberCode AccountNumberCode) *DirectDebitAttributesDebtorParty {

	m.AccountNumberCode = accountNumberCode

	return m
}

func (m *DirectDebitAttributesDebtorParty) WithAccountWith(accountWith AccountHoldingEntity) *DirectDebitAttributesDebtorParty {

	m.AccountWith = &accountWith

	return m
}

func (m *DirectDebitAttributesDebtorParty) WithoutAccountWith() *DirectDebitAttributesDebtorParty {
	m.AccountWith = nil
	return m
}

func (m *DirectDebitAttributesDebtorParty) WithAddress(address []string) *DirectDebitAttributesDebtorParty {

	m.Address = address

	return m
}

func (m *DirectDebitAttributesDebtorParty) WithCountry(country string) *DirectDebitAttributesDebtorParty {

	m.Country = country

	return m
}

// Validate validates this direct debit attributes debtor party
func (m *DirectDebitAttributesDebtorParty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumberCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountWith(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitAttributesDebtorParty) validateAccountNumberCode(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumberCode) { // not required
		return nil
	}

	if err := m.AccountNumberCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "debtor_party" + "." + "account_number_code")
		}
		return err
	}

	return nil
}

func (m *DirectDebitAttributesDebtorParty) validateAccountWith(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountWith) { // not required
		return nil
	}

	if m.AccountWith != nil {
		if err := m.AccountWith.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "debtor_party" + "." + "account_with")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitAttributesDebtorParty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitAttributesDebtorParty) UnmarshalBinary(b []byte) error {
	var res DirectDebitAttributesDebtorParty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitAttributesDebtorParty) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
