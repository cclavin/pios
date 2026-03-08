package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/cclavin/pios/templates"
	"gopkg.in/yaml.v3"
)

const tasksContractVersion = "0.4"

const ideContextPrompt = "You are operating under the PIOS execution contract. Read AGENTS.md. You must validate your work against the phase gates. Update STATUS.md and check off items in TASKS.md. Run `pios validate` frequently to ensure you are passing the gates."

const piosAscii = `
    ____  ____ ____  ____
   / __ \/  _//__  \/ ___/ 
  / /_/ // // / / /\__ \  
 / ____// // /_/ /___/ /  
/_/   /___/\____//____/    
`

const catAscii = `
   |\__/,|   (` + "`" + `\
 _.|o o  |_   ) )
-(((---(((--------
`

var pendingTaskRe = regexp.MustCompile(`^(?i)\s*(?:###\s+|-\s+)\[\s\]`)
var inProgressTaskRe = regexp.MustCompile(`^(?i)\s*(?:###\s+|-\s+)\[/\]`)
var completedTaskRe = regexp.MustCompile(`^(?i)\s*(?:###\s+|-\s+)\[[xX]\]`)
var checkboxLikeRe = regexp.MustCompile(`^(?i)\s*(?:###|-)\s*\[[^\]]*\]`)

var allowedStatusValues = map[string]struct{}{
	"Not Started": {},
	"In Progress": {},
	"Blocked":     {},
	"Done":        {},
}

type StatusFrontmatter struct {
	PiosVersion  string `yaml:"pios_version"`
	CurrentPhase string `yaml:"current_phase"`
	CurrentGate  string `yaml:"current_gate"`
	Status       string `yaml:"status"`
}

type TasksFrontmatter struct {
	PiosContractVersion string `yaml:"pios_contract_version"`
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "init":
		cmdInit(os.Args[2:])
	case "status":
		cmdStatus()
	case "validate":
		cmdValidate()
	case "cat", "meow":
		cmdCat()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("PIOS - AI Project Execution Contract CLI")
	fmt.Println("\nUsage:")
	fmt.Println("  pios <command>")
	fmt.Println("\nCommands:")
	fmt.Println("  init        Initialize PIOS templates in the current directory")
	fmt.Println("              Flags: --ide=<cursor|windsurf|claude>")
	fmt.Println("  status      Parse STATUS.md and output a JSON summary")
	fmt.Println("  validate    Validate tasks contract and phase gate completion")
}

func cmdInit(args []string) {
	initFs := flag.NewFlagSet("init", flag.ExitOnError)
	ideFlag := initFs.String("ide", "", "IDE context to scaffold (cursor, windsurf, claude)")
	_ = initFs.Parse(args)

	fmt.Println("Initializing PIOS in the current directory...")

	if err := os.MkdirAll("templates", 0755); err != nil {
		fmt.Printf("Error creating templates directory: %v\n", err)
		os.Exit(1)
	}

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

	if *ideFlag != "" {
		writeIDEContext(*ideFlag)
	}

	fmt.Println("PIOS templates successfully initialized.")
	printBanner()
}

func cmdStatus() {
	rootDir, err := findProjectRoot()
	if err != nil {
		fmt.Printf("{\"error\": \"%v\"}\n", err)
		os.Exit(1)
	}

	data, err := os.ReadFile(filepath.Join(rootDir, "STATUS.md"))
	if err != nil {
		fmt.Printf("{\"error\": \"Failed to read STATUS.md: %v\"}\n", err)
		os.Exit(1)
	}

	status, err := parseStatusFrontmatter(string(data))
	if err != nil {
		fmt.Printf("{\"error\": \"%v\"}\n", err)
		os.Exit(1)
	}

	pending, inProg, done := countTasks(filepath.Join(rootDir, "templates", "tasks.md"))

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
	rootDir, err := findProjectRoot()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	data, err := os.ReadFile(filepath.Join(rootDir, "templates", "tasks.md"))
	if err != nil {
		fmt.Println("Error: templates/tasks.md not found.")
		os.Exit(1)
	}

	version, err := parseTasksContractVersion(string(data))
	if err != nil {
		fmt.Printf("Validation Failed: %v\n", err)
		os.Exit(1)
	}
	if version != tasksContractVersion {
		fmt.Printf("Validation Failed: unsupported tasks contract version '%s'. Expected '%s'.\n", version, tasksContractVersion)
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")
	unchecked := 0
	malformedLines := make([]int, 0)

	for i, line := range lines {
		if checkboxLikeRe.MatchString(line) && !pendingTaskRe.MatchString(line) && !inProgressTaskRe.MatchString(line) && !completedTaskRe.MatchString(line) {
			malformedLines = append(malformedLines, i+1)
			continue
		}
		if pendingTaskRe.MatchString(line) || inProgressTaskRe.MatchString(line) {
			unchecked++
		}
	}

	if len(malformedLines) > 0 {
		fmt.Printf("Validation Failed: malformed checkbox syntax at lines %v.\n", malformedLines)
		os.Exit(1)
	}

	if unchecked > 0 {
		fmt.Printf("Validation Failed: found %d unchecked or in-progress items in tasks.\n", unchecked)
		os.Exit(1)
	}

	fmt.Println("Validation Passed: all task criteria are met.")
	os.Exit(0)
}

func writeIDEContext(ide string) {
	promptData := []byte(ideContextPrompt + "\n")
	switch strings.ToLower(ide) {
	case "cursor":
		_ = os.WriteFile(".cursorrules", promptData, 0644)
		fmt.Println("Generated .cursorrules")
	case "windsurf":
		_ = os.WriteFile(".windsurfrules", promptData, 0644)
		fmt.Println("Generated .windsurfrules")
	case "claude":
		_ = os.WriteFile("CLAUDE.md", promptData, 0644)
		fmt.Println("Generated CLAUDE.md")
	default:
		fmt.Printf("Unknown IDE flag value '%s'. Supported: cursor, windsurf, claude\n", ide)
	}
}

func printBanner() {
	// ANSI Color Codes
	// Orange: \033[38;5;208m
	// Green: \033[32m
	// Reset: \033[0m
	lines := strings.Split(strings.Trim(piosAscii, "\n"), "\n")
	for i, line := range lines {
		if i < len(lines)/2 {
			fmt.Printf("\033[38;5;208m%s\033[0m\n", line)
		} else {
			fmt.Printf("\033[32m%s\033[0m\n", line)
		}
	}
	fmt.Println()
}

func cmdCat() {
	fmt.Println("\033[38;5;208m" + strings.Trim(catAscii, "\n") + "\033[0m")
	os.Exit(0)
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir, nil
		}
		if _, err := os.Stat(filepath.Join(dir, "STATUS.md")); err == nil {
			return dir, nil
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}
	return "", fmt.Errorf("could not find project root containing .git or STATUS.md")
}

func parseStatusFrontmatter(content string) (StatusFrontmatter, error) {
	var status StatusFrontmatter

	frontmatterRaw, ok := extractFrontmatter(content)
	if !ok {
		return status, fmt.Errorf("failed to parse STATUS.md: missing YAML frontmatter")
	}

	if err := yaml.Unmarshal([]byte(frontmatterRaw), &status); err != nil {
		return status, fmt.Errorf("failed to parse YAML frontmatter: %v", err)
	}

	status.PiosVersion = strings.TrimSpace(status.PiosVersion)
	status.CurrentPhase = strings.TrimSpace(status.CurrentPhase)
	status.CurrentGate = strings.TrimSpace(status.CurrentGate)
	status.Status = strings.TrimSpace(status.Status)

	if status.PiosVersion == "" {
		return status, fmt.Errorf("failed to parse STATUS.md: missing required frontmatter key 'pios_version'")
	}
	if status.CurrentPhase == "" {
		return status, fmt.Errorf("failed to parse STATUS.md: missing required frontmatter key 'current_phase'")
	}
	if status.CurrentGate == "" {
		return status, fmt.Errorf("failed to parse STATUS.md: missing required frontmatter key 'current_gate'")
	}
	if status.Status == "" {
		return status, fmt.Errorf("failed to parse STATUS.md: missing required frontmatter key 'status'")
	}
	if _, ok := allowedStatusValues[status.Status]; !ok {
		return status, fmt.Errorf("failed to parse STATUS.md: unsupported status '%s'", status.Status)
	}

	return status, nil
}

func parseTasksContractVersion(content string) (string, error) {
	frontmatterRaw, ok := extractFrontmatter(content)
	if !ok {
		return "", fmt.Errorf("missing YAML frontmatter in templates/tasks.md")
	}

	var fm TasksFrontmatter
	if err := yaml.Unmarshal([]byte(frontmatterRaw), &fm); err != nil {
		return "", fmt.Errorf("failed to parse tasks YAML frontmatter: %v", err)
	}

	version := strings.TrimSpace(fm.PiosContractVersion)
	if version == "" {
		return "", fmt.Errorf("missing required frontmatter key 'pios_contract_version' in templates/tasks.md")
	}

	return version, nil
}

// extractFrontmatter isolates the metadata block between top-of-file '---' lines.
func extractFrontmatter(content string) (string, bool) {
	lines := strings.Split(content, "\n")
	if len(lines) == 0 {
		return "", false
	}

	firstLine := strings.TrimPrefix(lines[0], "\uFEFF")
	if strings.TrimSpace(firstLine) != "---" {
		return "", false
	}

	var fm []string
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "---" {
			return strings.Join(fm, "\n"), true
		}
		fm = append(fm, line)
	}

	return "", false
}

// countTasks tallies checkboxes representing PIOS task states.
func countTasks(path string) (pending, inProgress, completed int) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if pendingTaskRe.MatchString(line) {
			pending++
		} else if inProgressTaskRe.MatchString(line) {
			inProgress++
		} else if completedTaskRe.MatchString(line) {
			completed++
		}
	}
	return
}
