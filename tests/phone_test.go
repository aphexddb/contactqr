package tests

import (
	"testing"

	"github.com/aphexddb/contactqr/phone"
)

func TestParsePhoneNationalFormatting(t *testing.T) {

	// map of test data -> expected data
	formatNationalTests := make(map[string]string)

	// National formatting tests
	formatNationalTests["13142225555"] = "(314) 222-5555"
	formatNationalTests["3142225555"] = "(314) 222-5555"
	formatNationalTests["314.222.5555"] = "(314) 222-5555"
	formatNationalTests["314-222-5555"] = "(314) 222-5555"
	formatNationalTests["314 222 5555"] = "(314) 222-5555"
	formatNationalTests["+1 314 222 5555"] = "(314) 222-5555"
	formatNationalTests["1 314 222 5555"] = "(314) 222-5555"
	formatNationalTests["314222 5555"] = "(314) 222-5555"
	formatNationalTests["314 2225555"] = "(314) 222-5555"
	for test, expected := range formatNationalTests {
		t.Logf("Testing valid number '%s' -> format '%s'", test, expected)

		number, err := phone.ParsePhone(test)
		if err != nil {
			t.Errorf("Test failed: '%s' -> '%s', got: %s", test, expected, number)
		}
		if err != nil {
			t.Logf("Parse error: %s", err)
		} else if number != expected {
			t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, number)
		}
	}

}
func TestParsePhoneE164Formatting(t *testing.T) {
	// map of test data -> expected data
	formatE164Tests := make(map[string]string)

	// E164 formatting tests
	formatE164Tests["13142225555"] = "+13142225555"
	formatE164Tests["3142225555"] = "+13142225555"
	formatE164Tests["314.222.5555"] = "+13142225555"
	formatE164Tests["314-222-5555"] = "+13142225555"
	formatE164Tests["314 222 5555"] = "+13142225555"
	formatE164Tests["+1 314 222 5555"] = "+13142225555"
	formatE164Tests["1 314 222 5555"] = "+13142225555"
	formatE164Tests["314222 5555"] = "+13142225555"
	formatE164Tests["314 2225555"] = "+13142225555"
	for test, expected := range formatE164Tests {
		t.Logf("Testing valid number '%s' -> format '%s'", test, expected)

		number, err := phone.ParsePhoneE164(test)
		if err != nil {
			t.Errorf("Test failed: '%s' -> '%s', got: %s", test, expected, number)
		}
		if err != nil {
			t.Logf("Parse error: %s", err)
		} else if number != expected {
			t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, number)
		}
	}
}

func TestParsePhoneInvalid(t *testing.T) {

	// parse tests
	invalidTests := make(map[string]bool)
	invalidTests["314876085"] = false
	invalidTests["314 555 022"] = false

	for invalidNumber, expectedValid := range invalidTests {
		t.Logf("Testing invalid number '%s' -> boolean %v", invalidNumber, expectedValid)
		number, err := phone.ParsePhone(invalidNumber)
		if err != nil {
			t.Logf("Invalid number was caught: %s", err)
		} else {
			t.Errorf("Invalid number was accepted! [%s]", number)
		}
	}

}
