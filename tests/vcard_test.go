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
