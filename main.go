package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var x int
	planid, err := fmt.Sscanf(os.Getenv("PlanID"), "%d", &x)
	if err != nil {
		fmt.Println(err.Error())
	}

	suiteid, err := fmt.Sscanf(os.Getenv("SuiteID"), "%d", &x)
	if err != nil {
		fmt.Println(err.Error())
	}

	app := &App{
		Organization: os.Getenv("Organization"),
		Project:      os.Getenv("Project"),
		Pat:          os.Getenv("Pat"),
		PlanID:       planid,
		SuiteID:      suiteid,
		CsvFile:      "issues1.csv",
	}

	// Load Issues
	issues := app.ReadCsv()

	// Start Test Case
	var testCase TestCase
	var steps Steps
	for _, issue := range issues {

		if len(testCase.Fields.Title) > 0 && len(issue.Title) > 0 {
			app.SubmitWorkItem(steps, testCase)

			// Start a new testCase
			testCase = TestCase{}
			steps = Steps{}
		}

		if len(issue.Title) > 0 {
			fmt.Println(issue.Title)
			testCase.Fields = Fields{
				Title:         issue.Title,
				AreaPath:      issue.AreaPath,
				IterationPath: issue.AreaPath,
				State:         issue.State,
				AssignedTo:    issue.AssignedTo,
			}
		} else {
			step := Step{
				ID:   issue.TestStep,
				Type: "ActionStep",
				ParameterizedStrings: []ParameterizedString{
					{
						IsFormatted: "false",
						Text:        issue.StepAction,
					},
					{
						IsFormatted: "false",
						Text:        issue.StepExpected,
					},
				},
			}
			steps.Steps = append(steps.Steps, step)
		}
	}
	app.SubmitWorkItem(steps, testCase)
}
