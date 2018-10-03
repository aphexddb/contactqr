package vcard

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/aphexddb/contactqr/phone"
)

// Address represents a postal address
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

// VCard repreents vCard data
type VCard struct {
	first              string
	last               string
	org                string
	title              string
	email              string
	cellPhone          string
	homeAddress        string
	facebookProfileURL string
	twitterHandle      string
	url                string
	note               string
	invalidFields      []string
	vCardValue         string
}

// Option defines a functional option for VCard
type Option func(*VCard)

// vCard version
const vCardVersion = "4.0"

// Format formats an address in vCard format
func (a Address) Format() (string, error) {
	invalidFields := []string{}

	// Basic data validation
	if len(a.State) != 2 {
		return "", fmt.Errorf("State must be two letters a-z")
	}

	// do some sane string length checks
	if len(a.Street) > 128 {
		invalidFields = append(invalidFields, "Street is too long")
	}
	if len(a.City) > 64 {
		invalidFields = append(invalidFields, "City is too long")
	}
	if len(a.PostalCode) > 16 {
		invalidFields = append(invalidFields, "Postal Code is too long")
	}

	// if any validations failed, return an error
	if len(invalidFields) > 0 {
		return "", fmt.Errorf(strings.Join(invalidFields, ", "))
	}

	return fmt.Sprintf("%s;%s;%s;%s;USA", a.Street, a.City, strings.ToUpper(a.State), a.PostalCode), nil
}

// validate validates vCard data
func (vc *VCard) validate() error {

	// you should at least set your name
	if len(vc.first) == 0 && len(vc.last) == 0 {
		vc.invalidFields = append(vc.invalidFields, "You need a name")
	}

	// do some sane string length checks
	if len(vc.first) > 32 {
		vc.invalidFields = append(vc.invalidFields, "First name is too long")
	}
	if len(vc.last) > 32 {
		vc.invalidFields = append(vc.invalidFields, "Last name is too long")
	}
	if len(vc.org) > 64 {
		vc.invalidFields = append(vc.invalidFields, "Organization is too long")
	}
	if len(vc.title) > 64 {
		vc.invalidFields = append(vc.invalidFields, "Title is too long")
	}
	if len(vc.email) > 128 {
		vc.invalidFields = append(vc.invalidFields, "Email is too long")
	}
	if len(vc.facebookProfileURL) > 128 {
		vc.invalidFields = append(vc.invalidFields, "Facebook URL is too long")
	}
	if len(vc.twitterHandle) > 64 {
		vc.invalidFields = append(vc.invalidFields, "Twitter name is too long")
	}
	if len(vc.url) > 128 {
		vc.invalidFields = append(vc.invalidFields, "URL is too long")
	}
	if len(vc.note) > 255 {
		vc.invalidFields = append(vc.invalidFields, "Note is too long")
	}

	// if any validations failed, return an error
	if len(vc.invalidFields) > 0 {
		return fmt.Errorf(strings.Join(vc.invalidFields, ", "))
	}

	return nil
}

// First is a getter for the first name
func (vc *VCard) First() string {
	return vc.first
}

// Last is a getter for the last name
func (vc *VCard) Last() string {
	return vc.last
}

// Org is a getter for the org
func (vc *VCard) Org() string {
	return vc.org
}

// Title is a getter for the title
func (vc *VCard) Title() string {
	return vc.title
}

// Email is a getter for the email
func (vc *VCard) Email() string {
	return vc.email
}

// CellPhone is a getter for the cell phone
func (vc *VCard) CellPhone() string {
	return vc.cellPhone
}

// HomeAddress is a getter for the home address
func (vc *VCard) HomeAddress() string {
	return vc.homeAddress
}

// FacebookProfileURL is a getter for the FB profile url
func (vc *VCard) FacebookProfileURL() string {
	return vc.facebookProfileURL
}

// TwitterHandle is a getter for the Twitter handle
func (vc *VCard) TwitterHandle() string {
	return vc.twitterHandle
}

// URL is a getter for the URL
func (vc *VCard) URL() string {
	return vc.url
}

// Note is a getter for the note
func (vc *VCard) Note() string {
	return vc.note
}

// QRCode returns a QRCode in Base64
func (vc *VCard) QRCode(width, height int) string {
	return TextToBase64QRCodeImage(vc.vCardValue, width, height)
}

// String implements to string for VCard
func (vc *VCard) String() string {
	return vc.vCardValue
}

// generate creates the vcard value
// TODO: fully implement remaining fields in the vCard RFC
func (vc *VCard) generate() string {
	var buffer bytes.Buffer

	// Write each line of the vcard as defined in RFC6350 (https://tools.ietf.org/html/rfc6350).
	buffer.WriteString("BEGIN:VCARD\n")
	buffer.WriteString(fmt.Sprintf("VERSION:%s\n", vCardVersion))
	buffer.WriteString(fmt.Sprintf("N:%s;%s;;;\n", vc.last, vc.first))
	buffer.WriteString(fmt.Sprintf("FN:%s %s\n", vc.first, vc.last))
	buffer.WriteString(fmt.Sprintf("ORG:%s;\n", vc.org))
	buffer.WriteString(fmt.Sprintf("TITLE:%s\n", vc.title))
	// ROLE:Executive
	buffer.WriteString(fmt.Sprintf("EMAIL;PREF=1;TYPE=home:%s\n", vc.email))
	// EMAIL;PREF=2;TYPE=work:foo@bar.com
	buffer.WriteString(fmt.Sprintf("TEL;type=CELL:%s\n", vc.cellPhone))
	// TEL;PREF=2;TYPE=work:(123) 123-1234
	// TEL;VALUE=uri;PREF=3:tel:1234567890
	buffer.WriteString(fmt.Sprintf("ADR;PREF=1;TYPE=home:;;%s\n", vc.homeAddress))
	// ADR;PREF=2;TYPE=work:;;123 Forbes Ave\, Apt 1;San Francisco;CA;12345;USA
	buffer.WriteString(fmt.Sprintf("X-SOCIALPROFILE;PREF=1;TYPE=facebook:%s\n", vc.facebookProfileURL))
	buffer.WriteString(fmt.Sprintf("X-SOCIALPROFILE;PREF=2;TYPE=twitter;x-user=%s:x-apple:%s\n", vc.twitterHandle, vc.twitterHandle))
	buffer.WriteString(fmt.Sprintf("URL;PREF=1;TYPE=internet:%s\n", vc.url))
	// URL;PREF=2;TYPE=personal:http://www.johndoe.com
	// PHOTO;PREF=1;TYPE=work:https://upload.wikimedia.org/wikipedia/en/8/80/Wikipedia-logo-v2.svg
	// PHOTO;PREF=2;TYPE=home:data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQAQMAAAAlPW0iAAAABlBMVEUAAAD///+l2Z/dAAAAM0lEQVR4nGP4/5/h/1+G/58ZDrAz3D/McH8yw83NDDeNGe4Ug9C9zwz3gVLMDA/A6P9/AFGGFyjOXZtQAAAAAElFTkSuQmCC
	// BDAY:20000131
	// GENDER:M
	buffer.WriteString(fmt.Sprintf("NOTE;PREF=1:%s\n", vc.note))
	// NOTE;PREF=2:Another note.
	buffer.WriteString("END:VCARD")

	return buffer.String()
}

// First sets the first name on a vCard
func First(first string) func(*VCard) {
	return func(vc *VCard) {
		vc.first = first
	}
}

// Last sets the last name on a vCard
func Last(last string) func(*VCard) {
	return func(vc *VCard) {
		vc.last = last
	}
}

// Org sets the org name on a vCard
func Org(org string) func(*VCard) {
	return func(vc *VCard) {
		vc.org = org
	}
}

// Title sets the title on a vCard
func Title(title string) func(*VCard) {
	return func(vc *VCard) {
		vc.title = title
	}
}

// Email sets the email on a vCard
func Email(email string) func(*VCard) {
	return func(vc *VCard) {
		vc.email = email
	}
}

// CellPhone sets the cell phone on a vCard
func CellPhone(number string) func(*VCard) {
	return func(vc *VCard) {
		// ignore blank number
		if len(number) == 0 {
			return
		}

		formattedNumber, err := phone.ParsePhoneE164(number)
		if err != nil {
			vc.invalidFields = append(vc.invalidFields, err.Error())
			return
		}
		vc.cellPhone = formattedNumber
	}
}

// HomeAddress sets the home address on a vCard
func HomeAddress(address Address) func(*VCard) {
	return func(vc *VCard) {
		// ignore empty address
		if len(address.Street) == 0 {
			return
		}

		formattedAddress, err := address.Format()
		if err != nil {
			vc.invalidFields = append(vc.invalidFields, err.Error())
			return
		}
		vc.homeAddress = formattedAddress
	}
}

// FacebookProfileURL sets the facebook profile on a vCard
func FacebookProfileURL(url string) func(*VCard) {
	return func(vc *VCard) {
		vc.facebookProfileURL = url
	}
}

// TwitterHandle sets the twitter handle on a vCard
func TwitterHandle(url string) func(*VCard) {
	return func(vc *VCard) {
		vc.twitterHandle = url
	}
}

// URL sets the URL on a vCard
func URL(url string) func(*VCard) {
	return func(vc *VCard) {
		vc.url = url
	}
}

// Note sets the note on a vCard
func Note(note string) func(*VCard) {
	return func(vc *VCard) {
		vc.note = note
	}
}

// New creates a new VCard
func New(builders ...Option) (*VCard, error) {
	vcard := &VCard{
		first:              "",
		last:               "",
		org:                "",
		title:              "",
		email:              "",
		cellPhone:          "",
		homeAddress:        "",
		facebookProfileURL: "",
		twitterHandle:      "",
		url:                "",
		note:               "",
		invalidFields:      []string{},
	}

	for _, builder := range builders {
		builder(vcard)
	}

	err := vcard.validate()
	if err != nil {
		return nil, err
	}

	vcard.vCardValue = vcard.generate()

	return vcard, nil
}
