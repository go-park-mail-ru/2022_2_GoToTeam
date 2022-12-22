package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmailIsValid(t *testing.T) {
	res := EmailIsValidByRegExp("asd@asd.asd")
	assert.NotEqual(t, true, res)
}

func TestLoginIsValidByRegExp(t *testing.T) {
	res := LoginIsValidByRegExp("Aaaa")
	assert.Equal(t, true, res)
}

func TestPasswordIsValidByRegExp(t *testing.T) {
	res := PasswordIsValidByRegExp("Aaa3")
	assert.Equal(t, true, res)
}

func TestEmailIsValidByCustomValidation(t *testing.T) {
	res := EmailIsValidByCustomValidation("asd@asd.asd")
	assert.Equal(t, true, res)
}

func TestPasswordIsValidByCustomValidation(t *testing.T) {
	res := PasswordIsValidByCustomValidation("Aaaaa_gfewqrq")
	assert.Equal(t, true, res)
}

func TestLoginIsValidByCustomValidation(t *testing.T) {
	res := LoginIsValidByCustomValidation("Aaaaa")
	assert.Equal(t, true, res)
}

func TestEmailIsValidNegative(t *testing.T) {
	res := EmailIsValidByRegExp("asdasd.asd")
	assert.Equal(t, false, res)
}

func TestLoginIsValidByRegExpNegative(t *testing.T) {
	res := LoginIsValidByRegExp("Aa")
	assert.Equal(t, false, res)
}

func TestPasswordIsValidByRegExpNegative(t *testing.T) {
	res := PasswordIsValidByRegExp("Aa")
	assert.Equal(t, false, res)
}

func TestEmailIsValidByCustomValidationNegative(t *testing.T) {
	res := EmailIsValidByCustomValidation("asdasd.asd")
	assert.Equal(t, false, res)
}

func TestPasswordIsValidByCustomValidationNegative(t *testing.T) {
	res := PasswordIsValidByCustomValidation("Aaaaa_")
	assert.Equal(t, false, res)
}

func TestLoginIsValidByCustomValidationNegative(t *testing.T) {
	res := LoginIsValidByCustomValidation("A1/aa")
	assert.Equal(t, false, res)
}

func TestLoginIsValidByCustomValidationNegative2(t *testing.T) {
	res := LoginIsValidByCustomValidation("Aaa")
	assert.Equal(t, false, res)
}
