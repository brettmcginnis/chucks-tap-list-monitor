package main

import (
	"fmt"
	"github.com/brettmcginnis/chucks-tap-list-monitor/src/site"
)

func main() {
	fmt.Println("Hello, world!")
	beers := site.GreenWood.List()

	//url := os.Getenv("SLACK_WEBHOOK")
	//if url == "" {
	//	return fmt.Errorf("SLACK_WEBHOOK environment variable not set")
	//}

	for _, beer := range beers {
		fmt.Printf("beer %d:\n", beer.Tap)
		fmt.Printf("  beer: %s\n", beer.Beer)
		fmt.Printf("  ABV: %s\n", beer.ABV)
		fmt.Printf("  Origin: %s\n", beer.Origin)
		fmt.Printf("  Type: %s\n", beer.Type)
		fmt.Printf("  Price per oz: $%.2f\n", beer.PriceOz)
		fmt.Println()
	}
}
