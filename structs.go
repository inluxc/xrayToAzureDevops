package main

import (
	"encoding/xml"
	"time"
)

type App struct {
	Organization   string
	Project        string
	Pat            string
	PlanID         int
	SuiteID        int
	ActiveWorkItem int
	CsvFile        string
}

// TestCase represents the structure of a test case
type TestCase struct {
	Fields Fields `json:"fields"`
}

// Fields holds test case fields like title and steps
type Fields struct {
	Title         string `json:"System.Title"`
	Steps         string `json:"Microsoft.VSTS.TCM.Steps"`
	AreaPath      string `json:"System.AreaPath"`
	IterationPath string `json:"System.IterationPath"`
	State         string `json:"System.State"`
	AssignedTo    string `json:"System.AssignedTo"`
	Description   string `json:"System.Description"`
}

type Issues struct {
	Id           string `csv:"Id"`
	WorkItemType string `csv:"Work Item Type"`
	Title        string `csv:"Title"`
	TestStep     string `csv:"TestStep"`
	StepAction   string `csv:"Step Action"`
	StepExpected string `csv:"Step Expected"`
	AreaPath     string `csv:"Area Path"`
	AssignedTo   string `csv:"Assigned To"`
	State        string `csv:"State"`
}

type JSONPatch struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

type WorkItemResponse struct {
	ID     int `json:"id"`
	Rev    int `json:"rev"`
	Fields struct {
		SystemAreaPath      string `json:"System.AreaPath"`
		SystemTeamProject   string `json:"System.TeamProject"`
		SystemIterationPath string `json:"System.IterationPath"`
		SystemWorkItemType  string `json:"System.WorkItemType"`
		SystemState         string `json:"System.State"`
		SystemReason        string `json:"System.Reason"`
		SystemAssignedTo    struct {
			DisplayName string `json:"displayName"`
			URL         string `json:"url"`
			Links       struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"_links"`
			ID         string `json:"id"`
			UniqueName string `json:"uniqueName"`
			ImageURL   string `json:"imageUrl"`
			Descriptor string `json:"descriptor"`
		} `json:"System.AssignedTo"`
		SystemCreatedDate time.Time `json:"System.CreatedDate"`
		SystemCreatedBy   struct {
			DisplayName string `json:"displayName"`
			URL         string `json:"url"`
			Links       struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"_links"`
			ID         string `json:"id"`
			UniqueName string `json:"uniqueName"`
			ImageURL   string `json:"imageUrl"`
			Descriptor string `json:"descriptor"`
		} `json:"System.CreatedBy"`
		SystemChangedDate time.Time `json:"System.ChangedDate"`
		SystemChangedBy   struct {
			DisplayName string `json:"displayName"`
			URL         string `json:"url"`
			Links       struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"_links"`
			ID         string `json:"id"`
			UniqueName string `json:"uniqueName"`
			ImageURL   string `json:"imageUrl"`
			Descriptor string `json:"descriptor"`
		} `json:"System.ChangedBy"`
		SystemCommentCount                 int       `json:"System.CommentCount"`
		SystemTitle                        string    `json:"System.Title"`
		MicrosoftVSTSCommonStateChangeDate time.Time `json:"Microsoft.VSTS.Common.StateChangeDate"`
		MicrosoftVSTSCommonActivatedDate   time.Time `json:"Microsoft.VSTS.Common.ActivatedDate"`
		MicrosoftVSTSCommonActivatedBy     struct {
			DisplayName string `json:"displayName"`
			URL         string `json:"url"`
			Links       struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"_links"`
			ID         string `json:"id"`
			UniqueName string `json:"uniqueName"`
			ImageURL   string `json:"imageUrl"`
			Descriptor string `json:"descriptor"`
		} `json:"Microsoft.VSTS.Common.ActivatedBy"`
		MicrosoftVSTSCommonPriority      int    `json:"Microsoft.VSTS.Common.Priority"`
		MicrosoftVSTSTCMAutomationStatus string `json:"Microsoft.VSTS.TCM.AutomationStatus"`
	} `json:"fields"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		WorkItemUpdates struct {
			Href string `json:"href"`
		} `json:"workItemUpdates"`
		WorkItemRevisions struct {
			Href string `json:"href"`
		} `json:"workItemRevisions"`
		WorkItemComments struct {
			Href string `json:"href"`
		} `json:"workItemComments"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		WorkItemType struct {
			Href string `json:"href"`
		} `json:"workItemType"`
		Fields struct {
			Href string `json:"href"`
		} `json:"fields"`
	} `json:"_links"`
	URL string `json:"url"`
}

// Struct to represent the root element <steps>
type Steps struct {
	XMLName xml.Name `xml:"steps"`
	ID      string   `xml:"id,attr"`
	Last    string   `xml:"last,attr"`
	Steps   []Step   `xml:"step"`
}

// Struct to represent each <step>
type Step struct {
	ID                   string                `xml:"id,attr"`
	Type                 string                `xml:"type,attr"`
	ParameterizedStrings []ParameterizedString `xml:"parameterizedString"`
}

// Struct to represent the <parameterizedString> elements
type ParameterizedString struct {
	IsFormatted string `xml:"isformatted,attr"`
	Text        string `xml:",cdata"` // CDATA section
}
