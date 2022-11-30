package validators

import (
	"testing"
)

func TestValidators(t *testing.T) {
	if EmailIsValidByRegExp("asd@asd.asd") != true {
		t.Error()
	}
	if LoginIsValidByRegExp("Aaaa") != true {
		t.Error()
	}
	if PasswordIsValidByRegExp("Aaa3") != true {
		t.Error()
	}
	if EmailIsValidByCustomValidation("asd@asd.asd") != true {
		t.Error()
	}
	if LoginIsValidByCustomValidation("Aaaaa") != true {
		t.Error()
	}
	if PasswordIsValidByCustomValidation("Aaaaa_gfewqrq") != true {
		t.Error()
	}
}

func TestValidatorsNegative(t *testing.T) {
	if EmailIsValidByRegExp("asdasd.asd") == true {
		t.Error()
	}
	if LoginIsValidByRegExp("Aa") == true {
		t.Error()
	}
	if PasswordIsValidByRegExp("Aa") == true {
		t.Error()
	}
	if EmailIsValidByCustomValidation("asdasd.asd") == true {
		t.Error()
	}
	if LoginIsValidByCustomValidation("Aaa") == true {
		t.Error()
	}
	if LoginIsValidByCustomValidation("A1/aa") == true {
		t.Error()
	}
	if PasswordIsValidByCustomValidation("Aaaaa_") == true {
		t.Error()
	}
}
