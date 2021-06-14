package main

import (
	combatcalculator "combatCalculator"
	"log"
	"os"
	"strconv"

	"golang.org/x/text/message"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("USAGE: go run ./cmd/parseDropTable <path to drop table html> <name of zone> <kph>")
	}
	path := os.Args[1]
	dropTable, err := combatcalculator.ReadDropTable(path)
	if err != nil {
		log.Fatalf("Error while parsing drop table: %v", err)
	}

	log.Printf("\n--------------------------------------\n")
	for key, table := range dropTable {
		log.Printf("%s:\n%v\n\n--------\n\n", key, table)
	}

	itemList, err := combatcalculator.LoadItems()
	if err != nil {
		log.Fatalf("Error while loading prices: %v\n", err)
	}

	zones, err := combatcalculator.AddDropTablesToZones(combatcalculator.DefaultZones, dropTable)
	if err != nil {
		log.Fatalf("Error while adding drop tables to zones: %v", err)
	}

	zoneName := os.Args[2]

	kph, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatalf("Error while converting third arg, %s, to float: %v", os.Args[3], err)
	}

	zone, ok := zones[zoneName]
	if !ok {
		log.Fatalf("No zone found for %s", zoneName)
	}
	gold, err := combatcalculator.CalcDropsForZone(zone, kph, itemList)
	if err != nil {
		log.Fatalf("Error while calculating drops: %v", err)
	}

	p := message.NewPrinter(message.MatchLanguage("en"))
	p.Printf("Drops for %s:\n%v\n", zoneName, gold)
	p.Printf("total: %.0f\n", gold["total"])
}
