package tests

import (
	"testing"

	"github.com/Zadigo/citadels/internal/districts"
	"github.com/stretchr/testify/assert"
)

func TestReadYamlFile(t *testing.T) {
	items, err := districts.ReadYamlFile()

	assert.NoError(t, err)
	assert.Len(t, items.Unique, 14)
	assert.Len(t, items.Nobel, 3)
	assert.Len(t, items.Trade, 6)
	assert.Len(t, items.Military, 4)
	assert.Len(t, items.Religious, 4)
}

// func BenchmarkReadYamlFile(b *testing.B) {
// 	b.ResetTimer()

// 	_, _ = districts.ReadYamlFile()
// 	// for b.Loop() {
// 	// }
// }

func TestGenerateDistricts(t *testing.T) {
	d, err := districts.GenerateDistricts()
	assert.NoError(t, err)
	assert.NotEmpty(t, d)
	assert.Len(t, d, districts.UniqueDistrictCount)
}
