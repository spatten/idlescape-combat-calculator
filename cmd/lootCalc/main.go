package main

import (
	combatcalculator "combatCalculator"
	"log"
	"os"
	"sort"

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

	total, err := itemList.CalculateTotal(items)
	if err != nil {
		log.Fatalf("Error while calculating total for item list: %v", err)
	}

	p := message.NewPrinter(message.MatchLanguage("en"))

	p.Println()

	table := itemList.Table(items)
	sort.Sort(table)
	p.Printf(table.String())

	p.Printf("\ntotal: %d\n", total)
}
