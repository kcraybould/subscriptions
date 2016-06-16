package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Unsubscribe unsubscribe

swagger:model Unsubscribe
*/
type Unsubscribe struct {

	/* The subscribed email address.

	Required: true
	Pattern: ^.*\@.*\..*
	*/
	EmailAddress *string `json:"emailAddress"`
}

// Validate validates this unsubscribe
func (m *Unsubscribe) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmailAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Unsubscribe) validateEmailAddress(formats strfmt.Registry) error {

	if err := validate.Required("emailAddress", "body", m.EmailAddress); err != nil {
		return err
	}

	if err := validate.Pattern("emailAddress", "body", string(*m.EmailAddress), `^.*\@.*\..*`); err != nil {
		return err
	}

	return nil
}
