package utils_test

import (
	"testing"

	"ardeolib.sapions.com/utils"
	"github.com/stretchr/testify/assert"
)

func TestValidateUserPswd(t *testing.T) {
	pswd := "Ab1aaasdsdsds0@#$%"
	ok := utils.ValidateUserPswd(&pswd)
	assert.True(t, ok)

	pswd = "Ab1234567890"
	ok = utils.ValidateUserPswd(&pswd)
	assert.False(t, ok)

	pswd = "A1234567890@#$%1"
	ok = utils.ValidateUserPswd(&pswd)
	assert.False(t, ok)

	pswd = "a1234567890@#$%12"
	ok = utils.ValidateUserPswd(&pswd)
	assert.False(t, ok)

	pswd = "AB1234567890@#$%AB"
	ok = utils.ValidateUserPswd(&pswd)
	assert.False(t, ok)

}

func TestValidateUserEmailFormat(t *testing.T) {
	email := "test@test.com"
	ok := utils.ValidateUserEmailFormat(email)
	assert.True(t, ok)

	email = "test@test"
	ok = utils.ValidateUserEmailFormat(email)
	assert.False(t, ok)

	email = "test.test.com"
	ok = utils.ValidateUserEmailFormat(email)
	assert.False(t, ok)

	email = "test@test."
	ok = utils.ValidateUserEmailFormat(email)
	assert.False(t, ok)

}
