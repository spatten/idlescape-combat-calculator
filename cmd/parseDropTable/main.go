package main

import (
	combatcalculator "combatCalculator"
	"log"
	"os"
)

func main() {
	path := os.Args[1]
	dropTable, err := combatcalculator.ReadDropTable(path)
	if err != nil {
		log.Fatalf("Error while parsing drop table: %v", err)
	}

	log.Printf("\n--------------------------------------\n")
	for key, table := range dropTable {
		log.Printf("%s:\n%v\n\n--------\n\n", key, table)
	}

	zones, err := combatcalculator.AddDropTablesToZones(combatcalculator.DefaultZones, dropTable)
	if err != nil {
		log.Fatalf("Error while adding drop tables to zones: %v", err)
	}

	for name, zone := range zones {
		log.Printf("%s: %v", name, zone)
	}
}
