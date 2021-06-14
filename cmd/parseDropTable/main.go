package main

import (
	combatcalculator "combatCalculator"
	"log"
	"os"
)

func main() {
	path := os.Args[1]
	dropTable, err := combatcalculator.ReadLootTable(path)
	if err != nil {
		log.Fatalf("Error while parsing drop table: %v", err)
	}

	log.Printf("\n--------------------------------------\n")
	for key, table := range dropTable {
		log.Printf("%s:\n%v\n\n--------\n\n", key, table)
	}
}
