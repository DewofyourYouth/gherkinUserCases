package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

// PermissionLevel is an int that represent the level of access to chs.care
type PermissionLevel int

// UserLevel is an type (struct) that represents a user story in gherkin
type UserLevel struct {
	Name     string
	Level    PermissionLevel
	Scenario string
	Result   string
	AndSlice []string
}

// PermissionLevels for chs.care
const (
	PHYSICIAN PermissionLevel = iota
	FACILITY
	REGION
	DIVISION
	COMPANY
	ADMIN
)

const gherkinTemplate = ` User with {{.Name}} level access {{.Scenario}}
  ` + "`GIVEN`" + ` I am a user with {{.Name}} access level
  ` + "`WHEN`" + ` I {{.Scenario}}
  ` + "`THEN`" + ` I {{.Result}}
  {{if gt (len .AndSlice) 0}}{{ range  .AndSlice }}` + "`AND`" + ` {{.}} 
  {{end}}{{end}}

`

func main() {
	var scenario string
	var result string
	var andString string
	var xAnd []string
	flag.StringVar(&andString, "andStr", "", "additional requirements separated by &s")
	accessLevelPtr := flag.Int("minLevel", 0, "The minimum access level")
	startNumPtr := flag.Int("startNum", 1, "scenario starts here")
	flag.StringVar(&scenario, "scenario", "describe the scenario here", "fills in the SCENARIO and WHEN part")
	flag.StringVar(&result, "result", "write what should happen here", "fills in the THEN part")
	flag.Parse()
	if andString != "" {
		if strings.Contains(andString, "&") {
			xAnd = strings.Split(andString, "&")
		} else {
			xAnd = append(xAnd, andString)
		}
	}
	userTypes := []UserLevel{
		{"PHYSICIAN", PHYSICIAN, scenario, result, xAnd},
		{"FACILITY", FACILITY, scenario, result, xAnd},
		{"REGION", REGION, scenario, result, xAnd},
		{"DIVISION", DIVISION, scenario, result, xAnd},
		{"COMPANY", COMPANY, scenario, result, xAnd},
		{"ADMIN", ADMIN, scenario, result, xAnd},
	}
	t := template.Must(template.New("gherkin").Parse(gherkinTemplate))
	for i, v := range userTypes {
		if int(v.Level) >= *accessLevelPtr {
			fmt.Printf("`SCENARIO %d`: ", i+*startNumPtr-*accessLevelPtr)
			err := t.Execute(os.Stdout, v)
			if err != nil {
				log.Println("Template Error:", err)
			}
		}
	}
}
