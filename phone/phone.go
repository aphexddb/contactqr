package phone

import (
	"fmt"

	"github.com/ttacon/libphonenumber"
)

// ParsePhone looks for a valid number string and returns a libphonenumber phone number,
// throws an error if otherwise
func parse(number string) (*libphonenumber.PhoneNumber, error) {
	parsedNum, parseErr := libphonenumber.Parse(number, "US")
	if parseErr != nil {
		return nil, fmt.Errorf("Unable to parse phone number '%s': %s", number, parseErr)
	}

	ok := libphonenumber.IsValidNumber(parsedNum)
	if !ok {
		return nil, fmt.Errorf("Invalid Number '%s'", number)
	}

	return parsedNum, nil
}

// ParsePhone returns a number formmated in National format
func ParsePhone(number string) (string, error) {
	if number == "" {
		return "", fmt.Errorf("Phone number missing")
	}

	parsedNum, err := parse(number)
	formattedNum := libphonenumber.Format(parsedNum, libphonenumber.NATIONAL)
	return formattedNum, err
}

// ParsePhoneE164 returns a number formmated in E164
func ParsePhoneE164(number string) (string, error) {
	parsedNum, err := parse(number)
	formattedNum := libphonenumber.Format(parsedNum, libphonenumber.E164)
	return formattedNum, err
}
