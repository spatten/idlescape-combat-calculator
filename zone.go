package combatcalculator

import (
	"fmt"

	"golang.org/x/text/message"
)

type Zone []ZoneMonster

type ZoneMonster struct {
	Name          string
	DropTable     DropTable
	EncounterRate float64
}

type ZoneData struct {
	Name     string
	Monsters []ZoneMonster
}

var farm = ZoneData{
	Name: "Farm",
	Monsters: []ZoneMonster{
		{
			Name:          "Small Rat",
			EncounterRate: 0.30,
		},
		{
			Name:          "Chicken",
			EncounterRate: 0.30,
		},
		{
			Name:          "Cow",
			EncounterRate: 0.30,
		},
		{
			Name:          "Goblin",
			EncounterRate: 0.10,
		},
	},
}

var caves = ZoneData{
	Name: "Caves",
	Monsters: []ZoneMonster{
		{
			Name:          "Goblin",
			EncounterRate: 0.45,
		},
		{
			Name:          "Imp",
			EncounterRate: 0.30,
		},
		{
			Name:          "Greater Imp",
			EncounterRate: 0.15,
		},
	},
}

var city = ZoneData{
	Name: "City",
	Monsters: []ZoneMonster{
		{
			Name:          "Guard",
			EncounterRate: 0.80,
		},
		{
			Name:          "Black Knight",
			EncounterRate: 0.20,
		},
	},
}

var lavamaze = ZoneData{
	Name: "Lava Maze",
	Monsters: []ZoneMonster{
		{
			Name:          "Deadly Red Spider",
			EncounterRate: 0.40,
		},
		{
			Name:          "Lesser Demon",
			EncounterRate: 0.60,
		},
	},
}

var corruptedlands = ZoneData{
	Name: "Corrupted Lands",
	Monsters: []ZoneMonster{
		{
			Name:          "Corrupted Tree",
			EncounterRate: 0.45,
		},
		{
			Name:          "Infected Naga",
			EncounterRate: 0.45,
		},
		{
			Name:          "Bone Giant",
			EncounterRate: 0.10,
		},
	},
}

var giants = ZoneData{
	Name: "Valley of Giants",
	Monsters: []ZoneMonster{
		{
			Name:          "Fire Giant",
			EncounterRate: 1.0 / 3.0,
		},
		{
			Name:          "Moss Giant",
			EncounterRate: 1.0 / 3.0,
		},
		{
			Name:          "Ice Giant",
			EncounterRate: 1.0 / 3.0,
		},
	},
}

var DefaultZones = []ZoneData{
	farm,
	caves,
	city,
	lavamaze,
	corruptedlands,
	giants,
}

func AddDropTablesToZones(zones []ZoneData, dropTables map[string]DropTable) (map[string]ZoneData, error) {
	zoneData := make(map[string]ZoneData)
	for _, zone := range zones {
		for n, monster := range zone.Monsters {
			table, ok := dropTables[monster.Name]
			if !ok {
				return nil, fmt.Errorf("no droptable found for monster %s", monster.Name)
			}
			zone.Monsters[n].DropTable = table
		}
		zoneData[zone.Name] = zone
	}
	return zoneData, nil
}

func CalcDropsForZone(zone ZoneData, zoneKPH float64, itemList ItemList) (map[string]float64, error) {
	res := make(map[string]float64)
	var total float64
	for _, monster := range zone.Monsters {
		monsterKPH := monster.EncounterRate * zoneKPH
		fmt.Printf("%s (%.2f kills):\n", monster.Name, monster.EncounterRate*zoneKPH)
		gold, err := CalcDropsForMonster(monsterKPH, monster.DropTable, itemList)
		if err != nil {
			return res, fmt.Errorf("calculating drops for %s: %v", monster.Name, err)
		}
		res[monster.Name] = gold
		total += gold
	}
	res["total"] = total
	return res, nil
}

func CalcDropsForMonster(kills float64, dropTable DropTable, itemList ItemList) (float64, error) {
	total := 0.0
	p := message.NewPrinter(message.MatchLanguage("en"))
	for _, row := range dropTable {
		item, ok := itemList[row.Name]
		if !ok {
			return 0.0, fmt.Errorf("%s not found in pricelist", row.Name)
		}
		amount := float64(item.Price) * row.Chances[6] * kills

		p.Printf("  %s: %d (price) x %.6f (chance) x %.0f (kph) = %.2f\n", row.Name, item.Price, row.Chances[6], kills, amount)
		total += amount
	}
	p.Printf("  Total: %.2f\n", total)
	return total, nil
}
