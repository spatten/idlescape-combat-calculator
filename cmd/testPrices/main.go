package main

import (
	combatcalculator "combatCalculator"
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

	total := itemList.CalculateTotal(items)
	p := message.NewPrinter(message.MatchLanguage("en"))
	p.Printf("total: %d\n", total)
}
