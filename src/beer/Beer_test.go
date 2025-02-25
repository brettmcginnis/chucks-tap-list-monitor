package beer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoved(t *testing.T) {
	oldList := []Beer{
		{1, "beer1", "5.0", "IPA", "yellow"},
	}
	var newList []Beer

	added, removed, changed := Diff(oldList, newList)

	assert.Equal(t, 0, len(added))
	assert.Equal(t, 1, len(removed))
	assert.Equal(t, 0, len(changed))
}

func TestAdded(t *testing.T) {
	var oldList []Beer

	newList := []Beer{
		{1, "beer1", "5.0", "IPA", "yellow"},
	}

	added, removed, changed := Diff(oldList, newList)

	assert.Equal(t, 1, len(added))
	assert.Equal(t, 0, len(removed))
	assert.Equal(t, 0, len(changed))
}

func TestAddedAndRemoved(t *testing.T) {
	oldList := []Beer{
		{1, "beer1", "5.0", "IPA", "yellow"},
	}
	newList := []Beer{
		{1, "beer2", "5.0", "IPA", "yellow"},
	}

	added, removed, changed := Diff(oldList, newList)

	assert.Equal(t, 1, len(added))
	assert.Equal(t, 1, len(removed))
	assert.Equal(t, 0, len(changed))
}
