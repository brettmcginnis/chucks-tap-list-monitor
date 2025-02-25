package beer

import (
	"reflect"
)

type Beer struct {
	Tap            int
	BreweryAndName string
	ABV            string
	Type           string
	Color          string
}

func Diff(old, new []Beer) (added, removed, changed []Beer) {
	newMap := toMap(new)
	oldMap := toMap(old)

	for _, beer2 := range newMap {
		if beer1, exists := oldMap[beer2.BreweryAndName]; !exists {
			added = append(added, beer2)
		} else if !reflect.DeepEqual(beer1, beer2) {
			changed = append(changed, beer2)
		}
	}

	for _, beer1 := range oldMap {
		if _, exists := newMap[beer1.BreweryAndName]; !exists {
			removed = append(removed, beer1)
		}
	}

	return added, removed, changed
}

func toMap(beers []Beer) map[string]Beer {
	m := make(map[string]Beer)
	for _, beer := range beers {
		m[beer.BreweryAndName] = beer
	}
	return m
}
