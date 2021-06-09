package combatcalculator

import "fmt"

type Monster struct {
	DropTable map[string]float32
	Name      string
}

var MonsterNames = []string{
	"Small Rat",
	"Chicken",
	"Cow",
	"Goblin",
	"Imp",
	"Greater Imp",
	"Guard",
	"Black Knight",
	"Deadly Red Spider",
	"Lesser Demon",
	"Spriggan",
	"Greater Demon",
	"Corrupted Tree",
	"Infected Naga",
	"Bone Giant",
	"Fire Giant",
	"Moss Giant",
	"Ice Giant",
}

// func ScrapeMonster(monsterName string) (Monster, error) {

// }

func GenerateWikiText() {
	fmt.Printf("<!DOCTYPE html>\n<html>\n  <body>\n")
	for _, name := range MonsterNames {
		fmt.Printf("<div class=\"monster\">\n  <p>%s</p>\n{{Combat Logs|%s}}\n</div>\n\n", name, name)
	}
	fmt.Printf("</body></html>")
}
