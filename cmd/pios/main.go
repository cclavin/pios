package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/cclavin/pios/templates"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
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
	case "next":
		cmdNext()
	case "mcp":
		cmdMcp()
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
	fmt.Println("  next        Archive completed tasks, snapshot the milestone, and reset the board")
	fmt.Println("  mcp         Start the JSON-RPC stdio Model Context Protocol server")
}

func cmdNext() {
	if err := SnapshotMilestone(); err != nil {
		fmt.Printf("Error capturing milestone snapshot: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✓ Milestone snapshot complete. Completed tasks archived. STATUS block reset.")
}

func cmdInit(args []string) {
	initFs := flag.NewFlagSet("init", flag.ExitOnError)
	ideFlag := initFs.String("ide", "", "IDE context to scaffold (cursor, windsurf, claude)")
	_ = initFs.Parse(args)

	fmt.Println("Initializing PIOS in the current directory...")

	if err := InitializeTemplates(*ideFlag); err != nil {
		fmt.Printf("Initialization failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("PIOS templates successfully initialized.")
	printBanner()
}

func InitializeTemplates(ide string) error {
	if err := os.MkdirAll("templates", 0755); err != nil {
		return fmt.Errorf("error creating templates directory: %v", err)
	}

	files, err := fs.ReadDir(templates.FS, ".")
	if err != nil {
		return fmt.Errorf("error reading embedded templates: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		data, err := templates.FS.ReadFile(fileName)
		if err != nil {
			continue
		}

		if fileName == "status-template.md" {
			if _, err := os.Stat("STATUS.md"); os.IsNotExist(err) {
				_ = os.WriteFile("STATUS.md", data, 0644)
			}
		}

		targetPath := filepath.Join("templates", fileName)
		_ = os.WriteFile(targetPath, data, 0644)
	}

	if ide != "" {
		writeIDEContext(ide)
	}
	return nil
}

func cmdStatus() {
	out, err := GetStatusData()
	if err != nil {
		fmt.Printf("{\"error\": \"%v\"}\n", err)
		os.Exit(1)
	}

	jsonBytes, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		fmt.Printf("{\"error\": \"Failed to encode JSON. Error: %v\"}\n", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonBytes))
}

func GetStatusData() (map[string]interface{}, error) {
	rootDir, err := findProjectRoot()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(filepath.Join(rootDir, "STATUS.md"))
	if err != nil {
		return nil, fmt.Errorf("Failed to read STATUS.md: %v", err)
	}

	status, err := parseStatusFrontmatter(string(data))
	if err != nil {
		return nil, err
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
	return out, nil
}

func cmdValidate() {
	if err := ValidateContract(); err != nil {
		fmt.Printf("Validation Failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Validation Passed: all task criteria are met.")
	os.Exit(0)
}

func cmdMcp() {
	s := server.NewMCPServer("pios-mcp", "1.0.0")

	// pios_status
	statusTool := mcp.NewTool("pios_status",
		mcp.WithDescription("Parses the project's STATUS.md and TASKS.md to return the current phase gate, active tasks, and total progress."),
	)
	s.AddTool(statusTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		out, err := GetStatusData()
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get status data: %v", err)), nil
		}
		jsonBytes, err := json.MarshalIndent(out, "", "  ")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to encode JSON: %v", err)), nil
		}
		return mcp.NewToolResultText(string(jsonBytes)), nil
	})

	// pios_validate
	validateTool := mcp.NewTool("pios_validate",
		mcp.WithDescription("Programmatically scans the current tasks contract and asserts that no pending checklists exist before allowing the agent to proceed to the next milestone. If pending items exist, it returns an error."),
	)
	s.AddTool(validateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		err := ValidateContract()
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Validation Failed: %v", err)), nil
		}
		return mcp.NewToolResultText("Validation Passed: all task criteria are met."), nil
	})

	// pios_init
	initTool := mcp.NewTool("pios_init",
		mcp.WithDescription("Initializes the PIOS templates and writes them to the current directory."),
		mcp.WithString("ide", mcp.Description("Optional. IDE context to scaffold (cursor, windsurf, claude)")),
	)
	s.AddTool(initTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var ide string
		if args, ok := request.Params.Arguments.(map[string]interface{}); ok {
			if val, ok := args["ide"].(string); ok {
				ide = val
			}
		}

		err := InitializeTemplates(ide)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Init failed: %v", err)), nil
		}
		return mcp.NewToolResultText("PIOS templates successfully initialized."), nil
	})

	// pios_next
	nextTool := mcp.NewTool("pios_next",
		mcp.WithDescription("Transitions to the next milestone by snapshotting the current STATUS.md and templates/tasks.md into an archive, erasing [x] completed tasks from the active board, and resetting STATUS to 'Next Milestone Planning'."),
	)
	s.AddTool(nextTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		err := SnapshotMilestone()
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Milestone snapshot failed: %v", err)), nil
		}
		return mcp.NewToolResultText("Snapshot complete. Completed tasks archived. STATUS block reset."), nil
	})

	// Start the stdio JSON-RPC loop
	stdioServer := server.NewStdioServer(s)
	if err := stdioServer.Listen(context.Background(), os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "MCP server error: %v\n", err)
		os.Exit(1)
	}
}

func ValidateContract() error {
	rootDir, err := findProjectRoot()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(filepath.Join(rootDir, "templates", "tasks.md"))
	if err != nil {
		return fmt.Errorf("templates/tasks.md not found")
	}

	version, err := parseTasksContractVersion(string(data))
	if err != nil {
		return err
	}
	if version != tasksContractVersion {
		return fmt.Errorf("unsupported tasks contract version '%s'. Expected '%s'", version, tasksContractVersion)
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
		return fmt.Errorf("malformed checkbox syntax at lines %v", malformedLines)
	}

	if unchecked > 0 {
		return fmt.Errorf("found %d unchecked or in-progress items in tasks", unchecked)
	}

	return nil
}

func SnapshotMilestone() error {
	rootDir, err := findProjectRoot()
	if err != nil {
		return err
	}

	tasksPath := filepath.Join(rootDir, "templates", "tasks.md")
	statusPath := filepath.Join(rootDir, "STATUS.md")

	tasksData, err := os.ReadFile(tasksPath)
	if err != nil {
		return fmt.Errorf("templates/tasks.md not found")
	}

	statusData, err := os.ReadFile(statusPath)
	if err != nil {
		return fmt.Errorf("STATUS.md not found")
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	archiveDir := filepath.Join(rootDir, "templates", "archive", timestamp)

	if err := os.MkdirAll(archiveDir, 0755); err != nil {
		return fmt.Errorf("failed to create archive directory: %v", err)
	}

	if err := os.WriteFile(filepath.Join(archiveDir, "tasks.md"), tasksData, 0644); err != nil {
		return fmt.Errorf("failed to snapshot tasks.md: %v", err)
	}
	if err := os.WriteFile(filepath.Join(archiveDir, "STATUS.md"), statusData, 0644); err != nil {
		return fmt.Errorf("failed to snapshot STATUS.md: %v", err)
	}

	// Filter completed tasks
	lines := strings.Split(string(tasksData), "\n")
	var newTasks []string
	for _, line := range lines {
		if !completedTaskRe.MatchString(line) {
			newTasks = append(newTasks, line)
		}
	}

	if err := os.WriteFile(tasksPath, []byte(strings.Join(newTasks, "\n")), 0644); err != nil {
		return fmt.Errorf("failed to rewrite tasks.md: %v", err)
	}

	// Reset STATUS.md
	statusStr := string(statusData)
	phaseRe := regexp.MustCompile(`(?m)^current_phase:\s*".*"`)
	statusRe := regexp.MustCompile(`(?m)^status:\s*".*"`)

	statusStr = phaseRe.ReplaceAllString(statusStr, `current_phase: "Next Milestone Planning"`)
	statusStr = statusRe.ReplaceAllString(statusStr, `status: "Not Started"`)

	if err := os.WriteFile(statusPath, []byte(statusStr), 0644); err != nil {
		return fmt.Errorf("failed to rewrite STATUS.md: %v", err)
	}

	return nil
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
