package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// CreateTestCase creates a test case in Azure DevOps
func (app App) CreateTestCase(testCase TestCase) { // ID of the Test Suite

	// Construct the API URL
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/wit/workitems/$%s?api-version=6.0", app.Organization, app.Project, "Test%20Case")

	patches := []JSONPatch{
		{Op: "add", Path: "/fields/System.Title", Value: testCase.Fields.Title},
		{Op: "add", Path: "/fields/System.Description", Value: testCase.Fields.Description},
		{Op: "add", Path: "/fields/Microsoft.VSTS.TCM.Steps", Value: testCase.Fields.Steps},
		{Op: "add", Path: "/fields/System.AreaPath", Value: testCase.Fields.AreaPath},
		{Op: "add", Path: "/fields/System.IterationPath", Value: testCase.Fields.IterationPath},
		{Op: "add", Path: "/fields/System.AssignedTo", Value: testCase.Fields.AssignedTo},
		{Op: "add", Path: "/fields/System.Tags", Value: "Jira XRay, Manual Testing"},
	}

	// Convert test case to JSON
	jsonData, err := json.Marshal(patches)
	if err != nil {
		fmt.Println("Error marshaling test case data:", err)
		os.Exit(1)
	}

	// Create request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	// Set authentication (Basic Auth) and headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Type", "application/json-patch+json")
	encodedPAT := base64.StdEncoding.EncodeToString([]byte(":" + app.Pat))
	req.Header.Set("Authorization", "Basic "+encodedPAT)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		var workitem WorkItemResponse
		err = json.Unmarshal(body, &workitem)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Test case created successfully with id:", workitem.ID)
		app.ActiveWorkItem = workitem.ID

	} else {
		fmt.Printf("Failed to create test case. Status: %d\n", resp.StatusCode)
		fmt.Printf("Body: %s\n", body)
	}
	fmt.Println(" ---------------------- ")
}

func (app App) AddToTestPlan() {
	// Set to Test Plan
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/test/Plans/%d/suites/%d/testcases/%d?api-version=7.1",
		app.Organization, app.Project, app.PlanID, app.SuiteID, app.ActiveWorkItem)

	// Create request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	// Set authentication (Basic Auth) and headers
	encodedPAT := base64.StdEncoding.EncodeToString([]byte(":" + app.Pat))
	req.Header.Set("Authorization", "Basic "+encodedPAT)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		fmt.Println("Test case added to the test plan successfully")
	} else {
		fmt.Printf("Failed to add test case to test plan. Status: %d\n", resp.StatusCode)
		fmt.Printf("Body: %s\n", body)
	}
}

func (app App) SubmitWorkItem(steps Steps, testCase TestCase) {
	if len(steps.Steps) > 0 {
		steps.ID = "0"
		steps.Last = strconv.Itoa(len(steps.Steps))

		xmlData, err := xml.MarshalIndent(steps, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling XML:", err)
			return
		}
		testCase.Fields.Steps = string(xmlData)
	}
	// Create test case
	app.CreateTestCase(testCase)
}
