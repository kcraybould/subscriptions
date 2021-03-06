package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ErrorNotification error notification

swagger:model ErrorNotification
*/
type ErrorNotification struct {

	/* code
	 */
	Code string `json:"code,omitempty"`

	/* fields
	 */
	Fields []string `json:"fields,omitempty"`

	/* message

	Required: true
	*/
	Message *string `json:"message"`
}

// Validate validates this error notification
func (m *ErrorNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFields(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorNotification) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	return nil
}

func (m *ErrorNotification) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}
