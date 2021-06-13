package combatcalculator

import (
	"sort"

	"golang.org/x/text/message"
)

type ItemTable []ItemRow

var _ sort.Interface = &ItemTable{}

type ItemRow struct {
	Name    string
	Count   int
	Price   int
	Percent float32
}

func (it ItemTable) Len() int {
	return len(it)
}

func (it ItemTable) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

// Sort by percent, descending, then by Count
func (it ItemTable) Less(i, j int) bool {
	it1 := it[i]
	it2 := it[j]

	if it1 == it2 {
		return false
	}

	if it1.Percent != it2.Percent {
		return it1.Percent > it2.Percent
	}

	if it1.Count != it2.Count {
		return it1.Count > it2.Count
	}

	return false
}

func (it ItemTable) String() string {
	p := message.NewPrinter(message.MatchLanguage("en"))
	table := p.Sprintf("%20s\t%s\t%s\t%s\t%s\n", "Item", "Count", "Gold/Item", "Gold Total", "Percent")
	for _, row := range it {
		table = table + row.String() + "\n"
	}
	return table
}

func (row ItemRow) Amount() int {
	return row.Price * row.Count
}

func (row ItemRow) String() string {
	p := message.NewPrinter(message.MatchLanguage("en"))
	return p.Sprintf("%20s\t%d\t%10d\t%10d\t%.2f", row.Name, row.Count, row.Price, row.Amount(), row.Percent)
}
