 User with {{.Name}} level access {{.Scenario}}
  `GIVEN` I am a user with {{.Name}} access level
  `WHEN` I {{.Scenario}}
  `THEN` I {{.Result}}
  {{if gt (len .AndSlice) 0}}{{ range  .AndSlice }}`AND` {{.}}
  {{end}}{{end}}