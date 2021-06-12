package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodec(t *testing.T) {
	inputs := []int64{0, 62, 1000}
	outputs := []string{"0000", "0010", "00G8"}

	for i := range inputs {
		bs := make([]byte, 4)
		Base62Encode(inputs[i], bs)
		assert.Equal(t, outputs[i], string(bs))
		assert.Equal(t, inputs[i], Base62Decode(bs))
	}
}

func TestCodec32(t *testing.T) {
	inputs := []int64{0, 32, 360}
	outputs := []string{"0000", "0010", "00b8"}

	for i := range inputs {
		bs := make([]byte, 4)
		Base32Encode(inputs[i], bs)
		assert.Equal(t, outputs[i], string(bs))
		assert.Equal(t, inputs[i], Base32Decode(bs))
	}
}
