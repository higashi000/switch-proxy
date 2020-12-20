package loadresolv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadResolv(t *testing.T) {
	search, err := LoadResolv("testdata/test")
	assert.Equal(t, nil, err)
	assert.Equal(t, "", search)
}

func TestLoadResolv1(t *testing.T) {
	search, err := LoadResolv("testdata/test1")
	assert.Equal(t, nil, err)
	assert.Equal(t, "search test", search)
}
