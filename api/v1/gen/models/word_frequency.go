// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// WordFrequency Word frequency
// swagger:model WordFrequency
type WordFrequency struct {

	// count
	Count int64 `json:"count,omitempty"`

	// word
	Word string `json:"word,omitempty"`
}

// Validate validates this word frequency
func (m *WordFrequency) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WordFrequency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WordFrequency) UnmarshalBinary(b []byte) error {
	var res WordFrequency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}