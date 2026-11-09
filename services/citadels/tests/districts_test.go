package tests

import (
	"testing"

	"github.com/Zadigo/citadels/internal/districts"
	"github.com/stretchr/testify/assert"
)

func TestReadYamlFile(t *testing.T) {
	items, err := districts.ReadYamlFile()
	assert.NoError(t, err)
	assert.Len(t, items.Districts.Unique, 14)
	assert.IsType(t, items.Districts.Unique, []districts.District{})
}
