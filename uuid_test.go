package uuid

import (
	"sort"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func TestUUID(t *testing.T) {
	uuids := make([]string, 0)
	for i := 0; i < 100; i++ {
		id := UUID()
		assert.Equal(t, 18, len(id))
		uuids = append(uuids, id)
	}

	oldStr := strings.Join(uuids, "|")
	sort.Strings(uuids)
	newStr := strings.Join(uuids, "|")
	assert.Equal(t, oldStr, newStr)
}

func TestRandID(t *testing.T) {
	assert.Equal(t, 8, len(Rand(8)))
	assert.Equal(t, 14, len(RandID()))
}

func BenchmarkMyUUID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UUID()
	}
}

func BenchmarkGoogleUUID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		uuid.New().String()
	}
}
