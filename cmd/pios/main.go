package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/cclavin/pios/templates"
	"gopkg.in/yaml.v3"
)

type StatusFrontmatter struct {
	PiosVersion  string `yaml:"pios_version"`
	CurrentPhase string `yaml:"current_phase"`
	CurrentGate  string `yaml:"current_gate"`
	Status       string `yaml:"status"`
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "init":
		cmdInit()
	case "status":
		cmdStatus()
	case "validate":
		cmdValidate()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("PIOS - Project Intelligence Operating System CLI")
	fmt.Println("\nUsage:")
	fmt.Println("  pios <command>")
	fmt.Println("\nCommands:")
	fmt.Println("  init        Initialize PIOS templates in the current directory")
	fmt.Println("  status      Parse STATUS.md and output a JSON summary")
	fmt.Println("  validate    Scan Tasks to validate if phase exit criteria are met")
}

func cmdInit() {
	fmt.Println("Initializing PIOS in the current directory...")

	if err := os.MkdirAll("templates", 0755); err != nil {
		fmt.Printf("Error creating templates directory: %v\n", err)
		os.Exit(1)
	}

	// Read from the embedded filesystem
	files, err := fs.ReadDir(templates.FS, ".")
	if err != nil {
		fmt.Printf("Error reading embedded templates: %v\n", err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		
		fileName := file.Name()
		data, err := templates.FS.ReadFile(fileName)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", fileName, err)
			continue
		}

		// Pull the status template to the root folder out of convenience
		if fileName == "status-template.md" {
			if _, err := os.Stat("STATUS.md"); os.IsNotExist(err) {
				_ = os.WriteFile("STATUS.md", data, 0644)
			}
		}

		targetPath := filepath.Join("templates", fileName)
		if err := os.WriteFile(targetPath, data, 0644); err != nil {
			fmt.Printf("Error writing %s: %v\n", targetPath, err)
		}
	}

	fmt.Println("✓ PIOS templates successfully initialized.")
}

func cmdStatus() {
	data, err := os.ReadFile("STATUS.md")
	if err != nil {
		fmt.Printf("{\"error\": \"Failed to read STATUS.md: %v\"}\n", err)
		os.Exit(1)
	}

	frontmatterRaw := extractFrontmatter(string(data))

	var status StatusFrontmatter
	if err := yaml.Unmarshal([]byte(frontmatterRaw), &status); err != nil {
		fmt.Printf("{\"error\": \"Failed to parse YAML frontmatter: %v\"}\n", err)
		os.Exit(1)
	}

	pending, inProg, done := countTasks("templates/tasks.md")

	out := map[string]interface{}{
		"current_gate": status.CurrentGate,
		"tasks": map[string]int{
			"pending":     pending,
			"in_progress": inProg,
			"completed":   done,
		},
	}

	jsonBytes, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		fmt.Printf("{\"error\": \"Failed to encode JSON. Error: %v\"}\n", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonBytes))
}

func cmdValidate() {
	data, err := os.ReadFile("templates/tasks.md")
	if err != nil {
		fmt.Println("Error: templates/tasks.md not found.")
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")
	unchecked := 0

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		
		// Catch any empty checkboxes in the tasks checklist
		if strings.HasPrefix(trimmed, "- [ ]") || strings.HasPrefix(trimmed, "### [ ]") || strings.HasPrefix(trimmed, "### [/]") {
			unchecked++
		}
	}

	if unchecked > 0 {
		fmt.Printf("Validation Failed: Found %d unchecked or in-progress items in tasks.\n", unchecked)
		os.Exit(1)
	}

	fmt.Println("Validation Passed: All task criteria are met.")
	os.Exit(0)
}

// extractFrontmatter isolates the metadata block between two '---' lines
func extractFrontmatter(content string) string {
	lines := strings.Split(content, "\n")
	var fm []string
	inFm := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "---" {
			if !inFm {
				inFm = true
				continue
			} else {
				break
			}
		}
		if inFm {
			fm = append(fm, line)
		}
	}
	return strings.Join(fm, "\n")
}

// countTasks tallies checkboxes representing PIOS task states
func countTasks(filepath string) (pending, inProgress, completed int) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return // return 0,0,0 if missing
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "### [ ]") {
			pending++
		} else if strings.HasPrefix(trimmed, "### [/]") {
			inProgress++
		} else if strings.HasPrefix(trimmed, "### [x]") {
			completed++
		}
	}
	return
}
