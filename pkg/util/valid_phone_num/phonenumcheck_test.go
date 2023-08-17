package validphonenum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPhoneNumCN(t *testing.T) {
	phoneNumbers := []string{
		"12345678901",
		"13800138000",
		"99999999999",
		"abc12345678",
	}
	for _, v := range phoneNumbers {
		assert.True(t, IsPhoneNumCN(v))
	}
}
