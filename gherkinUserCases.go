package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
)

type PermissionLevel int
type UserLevel struct {
	Name     string
	Level    PermissionLevel
	Scenario string
	Result   string
}

const (
	PHYSICIAN PermissionLevel = iota
	FACILITY
	REGION
	DIVISION
	COMPANY
	ADMIN
)

const gherkinTemplate = ` User with {{.Name}} level access {{.Scenario}}
  GIVEN I am a user with {{.Name}} access level
  WHEN I {{.Scenario}}
  THEN I {{.Result}}
`

func main() {
	var scenario string
	var result string
	accessLevelPtr := flag.Int("minLevel", 0, "The minimum access level")
	flag.StringVar(&scenario, "scenario", "describe the scenario here", "fills in the SCENARIO and WHEN part")
	flag.StringVar(&result, "result", "write what should happen here", "fills in the THEN part")
	flag.Parse()
	userTypes := []UserLevel{
		{"PHYSICIAN", PHYSICIAN, scenario, result},
		{"FACILITY", FACILITY, scenario, result},
		{"REGION", REGION, scenario, result},
		{"DIVISION", DIVISION, scenario, result},
		{"COMPANY", COMPANY, scenario, result},
		{"ADMIN", ADMIN, scenario, result},
	}

	t := template.Must(template.New("gherkin").Parse(gherkinTemplate))
	skipped := 0
	for i, v := range userTypes {
		if int(v.Level) >= *accessLevelPtr {
			fmt.Printf("Scenario %d: ", i+1-skipped)
			err := t.Execute(os.Stdout, v)
			if err != nil {
				log.Println("Template Error:", err)
			}
		} else {
			skipped = skipped + 1
		}
	}
}
