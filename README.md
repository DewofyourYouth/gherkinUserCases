# GHERKIN USER CASES

This is a script that pushes repetitive silly stuff that I am required to format
in a specific way to Jira cards. It started with gherkin user cases, but will 
probably include more stuff eventually.

Currently, it only sends the results to standard output. Next step is sending it to Jira 
with these [Jira API](https://developer.atlassian.com/server/jira/platform/rest-apis/), [Jira API Examples](https://developer.atlassian.com/server/jira/platform/jira-rest-api-examples/)

## How To Generate Gherkin User Cases 

We were told by management that we need to generate gherkin user cases for each user type with access to any 
given feature. We have 6 types of users:

0. Physician
1. Facility
2. Region
3. Division
4. Company
5. Admin

### Example

If I put in the following command.

`./gherkinUserCases -minLevel=2 -scenario="go to a party" -result="will have a great time."`

It will output the following:

```
Scenario 1:  User with REGION level access go to a party
  GIVEN I am a user with REGION access level
  WHEN I go to a party
  THEN I will have a great time.
Scenario 2:  User with DIVISION level access go to a party
  GIVEN I am a user with DIVISION access level
  WHEN I go to a party
  THEN I will have a great time.
Scenario 3:  User with COMPANY level access go to a party
  GIVEN I am a user with COMPANY access level
  WHEN I go to a party
  THEN I will have a great time.
Scenario 4:  User with ADMIN level access go to a party
  GIVEN I am a user with ADMIN access level
  WHEN I go to a party
  THEN I will have a great time.
```

### Flags

#### -minLevel
* The minimum level of access to see the feature.
* Receives an integer between 0 and 5.
* Defaults to 0 (will show all levels from Physician)

#### -scenario
* Fills in the SCENARIO and WHEN part.
* Receives a string.
* Defaults to "describe the scenario here"

#### -result
* Fills in the THEN part
* Receives a string
* Defaults to "write what should happen here"