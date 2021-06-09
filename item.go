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

type gold uint32
type heat uint32
type ItemID int

// Item is an item in Idlescape
type Item struct {
	Name  string
	ID    ItemID
	Heat  heat
	Price gold
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

func LoadItems() (ItemList, error) {
	pricesPath := "./prices.json"
	pricesFile, err := os.Open(pricesPath)
	if err != nil {
		return nil, fmt.Errorf("opening prices json: %v", err)
	}
	defer pricesFile.Close()

	return ParseItems(pricesFile)
}

func ParseItems(pricesFile io.Reader) (ItemList, error) {
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
	items["Gold"] = Item{Name: "Gold", ID: 0, Heat: 0, Price: 1}
	return items, nil
}

func (itemList ItemList) CalculateTotal(items []ItemCount) gold {
	total := gold(0)
	for _, counter := range items {
		item, exists := itemList[counter.Name]
		if exists {
			total = total + item.Price*gold(counter.Count)
		}
	}
	return total
}

func ReadItemFile(path string) ([]ItemCount, error) {
	itemFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %v", err)
	}
	defer itemFile.Close()

	items, err := ParseItemFile(itemFile)
	if err != nil {
		return nil, fmt.Errorf("parsing item file: %v", err)
	}

	return items, nil
}

func ParseItemFile(itemFile io.Reader) ([]ItemCount, error) {
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
