package site

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomeBeer(t *testing.T) {
	assert.True(t, len(GreenWood.List()) > 0)
}
