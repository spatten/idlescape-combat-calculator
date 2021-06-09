package main

import (
	combatcalculator "combatCalculator"
	"log"
)

func main() {
	itemList, err := combatcalculator.LoadItems()
	if err != nil {
		log.Fatalf("Error while loading prices: %v\n", err)
	}

	items := map[string]int{
		"Iron Boots": 12,
		"Iron Ore":   100,
	}
	total := itemList.CalculateTotal(items)
	log.Printf("total: %v", total)
}
