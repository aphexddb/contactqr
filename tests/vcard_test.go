package tests

import (
	"testing"

	"github.com/aphexddb/contactqr/vcard"
	"github.com/stretchr/testify/assert"
)

func TestVCardPhoneFormat(t *testing.T) {
	vc, err := vcard.New(
		vcard.First("Jane"),
		vcard.Last("Doe"),
		vcard.CellPhone("3145551234"),
	)

	assert.Nil(t, err)
	assert.Equal(t, "+13145551234", vc.CellPhone())
}

func TestVCardAddressFormat(t *testing.T) {
	vc, err := vcard.New(
		vcard.First("Jane"),
		vcard.Last("Doe"),
		vcard.HomeAddress(vcard.Address{
			Street:     "123 Main Street, Suite #500",
			City:       "Los Angeles",
			State:      "CA",
			PostalCode: "90046",
		}),
	)

	assert.Nil(t, err)
	assert.Equal(t, "123 Main Street, Suite #500;Los Angeles;CA;90046;USA", vc.HomeAddress())
}

func TestVCardNew(t *testing.T) {
	vc, err := vcard.New(
		vcard.First("Jane"),
		vcard.Last("Doe"),
		vcard.Org("Another Castle Games"),
		vcard.Title("Mushroom Keeper"),
		vcard.Email("jane.doe@gmail.com"),
		vcard.CellPhone("3145551234"),
		vcard.HomeAddress(vcard.Address{
			Street:     "123 Main Street, Suite #500",
			City:       "Los Angeles",
			State:      "CA",
			PostalCode: "90046",
		}),
		vcard.FacebookProfileURL("http://www.facebook.com/janedoe"),
		vcard.TwitterHandle("janedoe"),
		vcard.URL("https://www.mywebsite.com"),
		vcard.Note("Hello world!"),
	)

	assert.Nil(t, err)
	assert.Equal(t, "Jane", vc.First())
	assert.Equal(t, "Doe", vc.Last())
	assert.Equal(t, "Another Castle Games", vc.Org())
	assert.Equal(t, "Mushroom Keeper", vc.Title())
	assert.Equal(t, "jane.doe@gmail.com", vc.Email())
	assert.Contains(t, vc.CellPhone(), "3145551234")
	assert.Contains(t, vc.HomeAddress(), ";USA")
	assert.Equal(t, "http://www.facebook.com/janedoe", vc.FacebookProfileURL())
	assert.Equal(t, "janedoe", vc.TwitterHandle())
	assert.Equal(t, "https://www.mywebsite.com", vc.URL())
	assert.Equal(t, "Hello world!", vc.Note())
}

func TestVCardMaxStringLengthValidation(t *testing.T) {
	_, err := vcard.New(
		vcard.First("---------------------------------------------------------------------------------------------------------------------------------"),
		vcard.Last("---------------------------------------------------------------------------------------------------------------------------------"),
		vcard.Org("---------------------------------------------------------------------------------------------------------------------------------"),
		vcard.Title("---------------------------------------------------------------------------------------------------------------------------------"),
		vcard.Email("---------------------------------------------------------------------------------------------------------------------------------"),
		vcard.CellPhone("5555"),
		vcard.HomeAddress(vcard.Address{
			Street:     "---------------------------------------------------------------------------------------------------------------------------------",
			City:       "---------------------------------------------------------------------------------------------------------------------------------",
			State:      "---------------------------------------------------------------------------------------------------------------------------------",
			PostalCode: "---------------------------------------------------------------------------------------------------------------------------------",
		}),
		vcard.FacebookProfileURL("---------------------------------------------------------------------------------------------------------------------------------"),
		vcard.TwitterHandle("---------------------------------------------------------------------------------------------------------------------------------"),
		vcard.URL("---------------------------------------------------------------------------------------------------------------------------------"),
		vcard.Note("----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------"),
	)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid Number '5555', State must be two letters a-z, First name is too long, Last name is too long, Organization is too long, Title is too long, Email is too long, Facebook URL is too long, Twitter name is too long, URL is too long, Note is too long", err.Error())
}

func TestVCardRequiredFields(t *testing.T) {
	_, err := vcard.New(
		vcard.First(""),
		vcard.Last(""),
		vcard.Org(""),
		vcard.Title(""),
		vcard.Email(""),
		vcard.CellPhone(""),
		vcard.HomeAddress(vcard.Address{
			Street:     "",
			City:       "",
			State:      "",
			PostalCode: "",
		}),
		vcard.FacebookProfileURL(""),
		vcard.TwitterHandle(""),
		vcard.URL(""),
		vcard.Note(""),
	)

	assert.Equal(t, "You need a name", err.Error())
}

func TestVCardFormatString(t *testing.T) {
	vc, err := vcard.New(
		vcard.First("Jane"),
		vcard.Last("Doe"),
		vcard.Org("Another Castle Games"),
		vcard.Title("Mushroom Keeper"),
		vcard.Email("jane.doe@gmail.com"),
		vcard.CellPhone("3145551234"),
		vcard.HomeAddress(vcard.Address{
			Street:     "123 Main Street, Suite #500",
			City:       "Los Angeles",
			State:      "CA",
			PostalCode: "90046",
		}),
		vcard.FacebookProfileURL("http://www.facebook.com/janedoe"),
		vcard.TwitterHandle("janedoe"),
		vcard.URL("https://www.mywebsite.com"),
		vcard.Note("Hello world!"),
	)

	assert.Nil(t, err)
	assert.Equal(t, `BEGIN:VCARD
VERSION:4.0
N:Doe;Jane;;;
FN:Jane Doe
ORG:Another Castle Games;
TITLE:Mushroom Keeper
EMAIL;PREF=1;TYPE=home:jane.doe@gmail.com
TEL;type=CELL:+13145551234
ADR;PREF=1;TYPE=home:;;123 Main Street, Suite #500;Los Angeles;CA;90046;USA
X-SOCIALPROFILE;PREF=1;TYPE=facebook:http://www.facebook.com/janedoe
X-SOCIALPROFILE;PREF=2;TYPE=twitter;x-user=janedoe:x-apple:janedoe
URL;PREF=1;TYPE=internet:https://www.mywebsite.com
NOTE;PREF=1:Hello world!
END:VCARD`,
		vc.String())
}

func TestVCardImageQRCode(t *testing.T) {
	vc, err := vcard.New(
		vcard.First("Jane"),
		vcard.Last("Doe"),
		vcard.TwitterHandle("janedoe"),
	)

	assert.Nil(t, err)
	assert.Equal(t, "data:image/png;base64,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",
		vc.QRCode(200, 200))
}
