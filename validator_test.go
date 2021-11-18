package innValidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_ValidTin2(t *testing.T) {
	t.Parallel()

	testData := []struct {
		tin    string
		result bool
	}{
		{"366512608416", true},
		{"366512608416", true},
		{"", false},
		{"366512608414", false},
		{"3665126084162", false},
		{"36651260841", false},
	}
	var valid bool
	for _, tt := range testData {
		if len(tt.tin) == 12 {
			valid = validTinNumber(tt.tin, 2, 1) && validTinNumber(tt.tin, 1, 0)
		} else {
			valid = validTinNumber(tt.tin, 1, 2)
		}

		assert.Equal(t, tt.result, valid)
	}

}
