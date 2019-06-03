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

// CurrencyAndAmount currency and amount
// swagger:model CurrencyAndAmount
type CurrencyAndAmount struct {

	// Amount that is being returned/kept. The sum of `recall-amount.amount` and `charges-amount.amount` have to equal the original payment amount.
	// Pattern: ^[0-9.]{0,20}$
	Amount string `json:"amount,omitempty"`

	// ISO currency code for `amount`
	Currency string `json:"currency,omitempty"`
}

func CurrencyAndAmountWithDefaults(defaults client.Defaults) *CurrencyAndAmount {
	return &CurrencyAndAmount{

		Amount: defaults.GetString("CurrencyAndAmount", "amount"),

		Currency: defaults.GetString("CurrencyAndAmount", "currency"),
	}
}

func (m *CurrencyAndAmount) WithAmount(amount string) *CurrencyAndAmount {

	m.Amount = amount

	return m
}

func (m *CurrencyAndAmount) WithCurrency(currency string) *CurrencyAndAmount {

	m.Currency = currency

	return m
}

// Validate validates this currency and amount
func (m *CurrencyAndAmount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CurrencyAndAmount) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	if err := validate.Pattern("amount", "body", string(m.Amount), `^[0-9.]{0,20}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CurrencyAndAmount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CurrencyAndAmount) UnmarshalBinary(b []byte) error {
	var res CurrencyAndAmount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *CurrencyAndAmount) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
