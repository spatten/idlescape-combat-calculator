package combatcalculator

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type ItemID int

// Item is an item in Idlescape
type Item struct {
	Name  string
	ID    ItemID
	Heat  int
	Price int
}

type ItemCount struct {
	Name  string
	Count int
}

type ItemList map[string]Item

// prices.json is an object with a key of "items" pointing at an array of items
type priceJSON struct {
	Items []Item
}

// LoadItems loads the items from prices.json
func LoadItems() (ItemList, error) {
	pricesPath := "./prices.json"
	pricesFile, err := os.Open(pricesPath)
	if err != nil {
		return nil, fmt.Errorf("opening prices json: %v", err)
	}
	defer pricesFile.Close()

	return parseItems(pricesFile)
}

func parseItems(pricesFile io.Reader) (ItemList, error) {
	var prices priceJSON
	items := make(ItemList)

	bytes, err := io.ReadAll(pricesFile)
	if err != nil {
		return nil, fmt.Errorf("reading file: %v", err)
	}
	if err := json.Unmarshal(bytes, &prices); err != nil {
		return nil, fmt.Errorf("unmarshalling prices json: %v", err)
	}

	for _, item := range prices.Items {
		items[item.Name] = item
	}
	maxID := len(items)
	items["Gold"] = Item{Name: "Gold", ID: ItemID(maxID), Heat: 0, Price: 1}
	helm, ok := items["Obsidian Helm"]
	if !ok {
		return nil, fmt.Errorf("obsidian helm not found")
	}
	helm.Price = 1500000
	items["Obsidian Helm"] = helm
	return items, nil
}

// CalculateTotal calculates the total worth in gold of a slice of ItemCount
func (itemList ItemList) CalculateTotal(items []ItemCount) (int, error) {
	total := 0
	for _, counter := range items {
		item, exists := itemList[counter.Name]
		if !exists {
			return 0, fmt.Errorf("item in list not found: %v", counter)
		}
		total = total + item.Price*counter.Count
	}
	return total, nil
}

// ReadItemFile reads an item file
// An item file should consist of many lines, each line having a
// item name and a count, separated by a comma. E.g.
//
// Gold, 15430000
// Sapphire, 12
// Emerald, 7
// Ruby, 5
// Black Opal, 2
//
// Lines starting with "#" are treated as comments and ignored
func ReadItemFile(path string) ([]ItemCount, error) {
	itemFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %v", err)
	}
	defer itemFile.Close()

	items, err := parseItemFile(itemFile)
	if err != nil {
		return nil, fmt.Errorf("parsing item file: %v", err)
	}

	return items, nil
}

func parseItemFile(itemFile io.Reader) ([]ItemCount, error) {
	scanner := bufio.NewScanner(itemFile)
	var items []ItemCount
	for lineNum := 0; scanner.Scan(); lineNum++ {
		line := scanner.Text()
		item, err := parseItemLine(line)
		if err != nil {
			return nil, fmt.Errorf("parsing line '%s': %v", line, err)
		}
		if item != (ItemCount{}) {
			items = append(items, item)
		}
	}
	return items, nil
}

// parseItemLine parses a line from an itemFile and returns an Item and a count
func parseItemLine(line string) (ItemCount, error) {
	// # starts a comment
	if strings.TrimSpace(line)[0] == '#' {
		return ItemCount{}, nil
	}
	splitted := strings.SplitN(line, ",", 2)
	if len(splitted) != 2 {
		return ItemCount{}, fmt.Errorf("invalid item line %s, no comma found", line)
	}
	name := splitted[0]
	count, err := strconv.Atoi(strings.TrimSpace(splitted[1]))
	if err != nil {
		return ItemCount{}, fmt.Errorf("converting %s to int in line %s: %v", splitted[1], line, err)
	}
	return ItemCount{Name: name, Count: count}, nil
}

func (items ItemList) Table(counts []ItemCount) ItemTable {
	total, _ := items.CalculateTotal(counts)
	table := make([]ItemRow, 0, len(counts))
	for _, count := range counts {
		item, ok := items[count.Name]
		if !ok {
			continue
		}
		amount := item.Price * count.Count
		percent := float32(amount) / float32(total) * 100.0
		row := ItemRow{
			Name:    count.Name,
			Count:   count.Count,
			Price:   item.Price,
			Percent: percent,
		}

		table = append(table, row)
	}
	return ItemTable(table)
}
