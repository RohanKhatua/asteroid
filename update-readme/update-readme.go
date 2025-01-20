// NOTE - This has nothing to do with the asteroid programming language.
// This simply shows the LOC in the readme.md file using a github action.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type CodeStats struct {
	Language    string `json:"language"`
	Files       int    `json:"files"`
	Lines       int    `json:"lines"`
	Blanks      int    `json:"blanks"`
	Comments    int    `json:"comments"`
	LinesOfCode int    `json:"linesOfCode"`
}

func main() {
	// API URL
	apiURL := "https://api.codetabs.com/v1/loc/?github=rohankhatua/asteroid&branch=main" // Replace with your API URL

	// Fetch data from API
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	// Parse JSON data
	var stats []CodeStats
	if err := json.Unmarshal(body, &stats); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// Build the markdown table
	table := "| Language    | Files | Lines | Blanks | Comments | Lines of Code |\n"
	table += "|-------------|-------|-------|--------|----------|---------------|\n"

	for _, stat := range stats {
		table += fmt.Sprintf(
			"| %-11s | %-5d | %-5d | %-6d | %-8d | %-13d |\n",
			stat.Language, stat.Files, stat.Lines, stat.Blanks, stat.Comments, stat.LinesOfCode,
		)
	}

	// Read the existing README.md file
	readmePath := "README.md"
	readmeContent, err := ioutil.ReadFile(readmePath)
	if err != nil {
		fmt.Printf("Error reading README.md: %v\n", err)
		return
	}

	// Locate the markers and replace the content between them
	startMarker := "<!---start--->"
	endMarker := "<!---end--->"

	markerRegex := regexp.MustCompile(fmt.Sprintf("(?s)(%s)(.*?)(%s)", regexp.QuoteMeta(startMarker), regexp.QuoteMeta(endMarker)))
	newContent := markerRegex.ReplaceAll(readmeContent, []byte(fmt.Sprintf("$1\n\n%s\n\n$3", table)))

	// Write the updated content back to README.md
	if err := ioutil.WriteFile(readmePath, newContent, 0644); err != nil {
		fmt.Printf("Error writing to README.md: %v\n", err)
		return
	}

	fmt.Println("README.md updated successfully!")
}
