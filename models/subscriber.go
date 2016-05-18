package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Subscriber subscriber

swagger:model Subscriber
*/
type Subscriber struct {

	/* address line1
	 */
	AddressLine1 string `json:"addressLine1,omitempty"`

	/* address line2
	 */
	AddressLine2 string `json:"addressLine2,omitempty"`

	/* address line3
	 */
	AddressLine3 string `json:"addressLine3,omitempty"`

	/* city
	 */
	City string `json:"city,omitempty"`

	/* country

	Required: true
	*/
	Country *string `json:"country"`

	/* The subscribed email address.

	Required: true
	Pattern: ^.*\@.*\..*
	*/
	EmailAddress *string `json:"emailAddress"`

	/* The first name of the subscriber.  This field is required

	Required: true
	*/
	FirstName *string `json:"firstName"`

	/* Internal id generated by Guest service.  Only exists on subscriptions for HHonors member.
	 */
	GuestID int64 `json:"guestId,omitempty"`

	/* Language code for this resource
	 */
	LanguageCode string `json:"languageCode,omitempty"`

	/* The last name of the subscriber. This field is ONLY required for guest
	 */
	LastName string `json:"lastName,omitempty"`

	/* postal code

	Required: true
	*/
	PostalCode *string `json:"postalCode"`

	/* For US base states, use the two char abbreviations specified in ISO 3166-2
	 */
	State string `json:"state,omitempty"`

	/* subscriptions
	 */
	Subscriptions *Subscriptions `json:"subscriptions,omitempty"`
}

// Validate validates this subscriber
func (m *Subscriber) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEmailAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePostalCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubscriptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subscriber) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

func (m *Subscriber) validateEmailAddress(formats strfmt.Registry) error {

	if err := validate.Required("emailAddress", "body", m.EmailAddress); err != nil {
		return err
	}

	if err := validate.Pattern("emailAddress", "body", string(*m.EmailAddress), `^.*\@.*\..*`); err != nil {
		return err
	}

	return nil
}

func (m *Subscriber) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *Subscriber) validatePostalCode(formats strfmt.Registry) error {

	if err := validate.Required("postalCode", "body", m.PostalCode); err != nil {
		return err
	}

	return nil
}

func (m *Subscriber) validateSubscriptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Subscriptions) { // not required
		return nil
	}

	if m.Subscriptions != nil {

		if err := m.Subscriptions.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}