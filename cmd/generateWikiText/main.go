package main

import combatcalculator "combatCalculator"

// USAGE:
// run this:
// go run ./cmd/generateWikiText | pbcopy
// Then go here and paste it into the input:
// https://idlescape.wiki/p/Special:ExpandTemplates
// copy the result into dropTables.html
// pbpaste > dropTables.html

func main() {
	combatcalculator.GenerateWikiText()
}
