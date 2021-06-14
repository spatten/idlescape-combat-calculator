package combatcalculator

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

type DropTable []DropRow

type DropRow struct {
	Chances map[int]float64
	Name    string
}

func ReadDropTable(path string) (map[string]DropTable, error) {
	tables := make(map[string]DropTable)
	lootFile, err := os.Open(path)
	if err != nil {
		return tables, fmt.Errorf("error opening loot file at %s: %v", path, err)
	}
	doc, err := html.Parse(lootFile)
	if err != nil {
		return tables, fmt.Errorf("error parsing loot file: %v", err)
	}

	mobs := cascadia.MustCompile(".monster").MatchAll(doc)
	nameSelector := cascadia.MustCompile("p")
	rowSelector := cascadia.MustCompile(".lootTable tr")
	colSelector := cascadia.MustCompile("td.lt-data-input")
	for _, mob := range mobs {
		name := nameSelector.MatchFirst(mob).FirstChild.Data
		log.Printf("\n------\n%s:\n", name)

		rows := rowSelector.MatchAll(mob)
		lootRows := make([]DropRow, 0, len(rows))
		// first row is the header
		for _, row := range rows {
			lootRow := DropRow{}
			cols := colSelector.MatchAll(row)
			if len(cols) == 0 {
				continue
			}
			lootRow.Name = strings.Trim(cols[1].FirstChild.Data, "[]")
			lootRow.Chances = make(map[int]float64)
			log.Printf("  %s\n", lootRow.Name)
			for n, col := range cols[2:] {
				thAmount := 6 - n
				lootRow.Chances[thAmount], err = parseRate(col.FirstChild.Data)
				if err != nil {
					return tables, fmt.Errorf("error parsing rate: %v", err)
				}
				log.Printf("    %d: %.6f\n", thAmount, lootRow.Chances[thAmount])
			}
			lootRows = append(lootRows, lootRow)
		}
		tables[name] = DropTable(lootRows)
	}
	return tables, nil
}

// parseRate takes the contents of a rate td and returns the rate
// A rate td can have the following formats:
// "-" -- empty. Return 0
// " 501.1 (357652803)" -- a rate possibly followed by a number in parens -- return the rate
// " 1 / 1385.8 (515)" -- a ratio possibly followed by a number in parens -- extract the numerator and the denominator and return the rate
func parseRate(rateString string) (float64, error) {
	rateString = strings.TrimSpace(rateString)
	ratioRE := regexp.MustCompile("^([0-9]+) ?/ ?([0-9]+[.]?[0-9]*)")
	if rateString == "-" {
		return 0.0, nil
	}
	ratioMatch := ratioRE.FindStringSubmatch(rateString)
	if len(ratioMatch) == 3 {
		numerator, err := strconv.ParseFloat(ratioMatch[1], 64)
		if err != nil {
			return 0.0, fmt.Errorf("error parsing numerator in rate %s: %v", rateString, err)
		}
		denominator, err := strconv.ParseFloat(ratioMatch[2], 64)
		if err != nil {
			return 0.0, fmt.Errorf("error parsing denominator in rate %s: %v", rateString, err)
		}
		return numerator / denominator, nil
	}
	rates := strings.SplitN(rateString, " ", 2)
	rate, err := strconv.ParseFloat(rates[0], 64)
	if err != nil {
		return 0.0, fmt.Errorf("error parsing rate %s: %v", rateString, err)
	}
	return rate, nil
}
