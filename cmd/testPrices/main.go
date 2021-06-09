package main

import (
	combatcalculator "combatCalculator"
	"fmt"
	"log"
	"os"

	"golang.org/x/text/message"
)

func main() {
	itemList, err := combatcalculator.LoadItems()
	if err != nil {
		log.Fatalf("Error while loading prices: %v\n", err)
	}

	// items := map[string]int{
	// 	"Iron Boots": 12,
	// 	"Iron Ore":   100,
	// }
	path := os.Args[1]
	items, err := combatcalculator.ReadItemFile(path)
	if err != nil {
		log.Fatalf("Error while loading item list: %v", err)
	}
	fmt.Printf("Items: %v\n", items)

	total, err := itemList.CalculateTotal(items)
	if err != nil {
		log.Fatalf("Error while calculating total for item list: %v", err)
	}

	p := message.NewPrinter(message.MatchLanguage("en"))
	p.Printf("total: %d\n", total)
}
