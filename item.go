package combatcalculator

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

	bytes, err := ioutil.ReadAll(pricesFile)
	if err != nil {
		return nil, fmt.Errorf("reading file: %v", err)
	}
	if err := json.Unmarshal(bytes, &prices); err != nil {
		return nil, fmt.Errorf("unmarshalling prices json: %v", err)
	}

	for _, item := range prices.Items {
		items[item.Name] = item
	}
	return items, nil
}

func (itemList ItemList) CalculateTotal(items map[string]int) gold {
	total := gold(0)
	for id, count := range items {
		item, exists := itemList[id]
		if exists {
			total = total + item.Price*gold(count)
		}
	}
	return total
}
