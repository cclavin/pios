package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExtractFrontmatter(t *testing.T) {
	content := `---
foo: bar
---
# Content
`
	fm, ok := extractFrontmatter(content)
	if !ok {
		t.Fatalf("expected frontmatter to be detected")
	}
	if fm != "foo: bar" {
		t.Fatalf("expected 'foo: bar', got '%s'", fm)
	}
}

func TestExtractFrontmatterMissing(t *testing.T) {
	content := `# Content without frontmatter`
	_, ok := extractFrontmatter(content)
	if ok {
		t.Fatalf("expected missing frontmatter")
	}
}

func TestRegexMatches(t *testing.T) {
	tests := []struct {
		line      string
		pending   bool
		inProg    bool
		completed bool
	}{
		{"### [ ] Task 1", true, false, false},
		{"### [/] Task 2", false, true, false},
		{"### [x] Task 3", false, false, true},
		{"- [ ] List task", true, false, false},
		{"- [/] List task", false, true, false},
		{"- [X] List task", false, false, true},
		{"  - [ ] Indented task", true, false, false},
		{"Some random text", false, false, false},
		{"-[ ] bad syntax", false, false, false},
		{"###[] bad syntax", false, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			if pendingTaskRe.MatchString(tt.line) != tt.pending {
				t.Errorf("expected pending %v for '%s'", tt.pending, tt.line)
			}
			if inProgressTaskRe.MatchString(tt.line) != tt.inProg {
				t.Errorf("expected inProg %v for '%s'", tt.inProg, tt.line)
			}
			if completedTaskRe.MatchString(tt.line) != tt.completed {
				t.Errorf("expected completed %v for '%s'", tt.completed, tt.line)
			}
		})
	}
}

func TestParseStatusFrontmatterValid(t *testing.T) {
	content := `---
pios_version: "1.0.0"
current_phase: "v1.0.0"
current_gate: "Plan Lock"
status: "In Progress"
---
# STATUS`

	status, err := parseStatusFrontmatter(content)
	if err != nil {
		t.Fatalf("expected valid status frontmatter, got error: %v", err)
	}
	if status.Status != "In Progress" {
		t.Fatalf("expected status In Progress, got %s", status.Status)
	}
}

func TestParseStatusFrontmatterMissingKey(t *testing.T) {
	content := `---
pios_version: "1.0.0"
current_phase: "v1.0.0"
status: "In Progress"
---`

	_, err := parseStatusFrontmatter(content)
	if err == nil {
		t.Fatalf("expected error for missing key")
	}
	if !strings.Contains(err.Error(), "current_gate") {
		t.Fatalf("expected missing current_gate error, got: %v", err)
	}
}

func TestParseStatusFrontmatterInvalidStatus(t *testing.T) {
	content := `---
pios_version: "1.0.0"
current_phase: "v1.0.0"
current_gate: "Plan Lock"
status: "Active"
---`

	_, err := parseStatusFrontmatter(content)
	if err == nil {
		t.Fatalf("expected error for invalid status")
	}
	if !strings.Contains(err.Error(), "unsupported status") {
		t.Fatalf("expected unsupported status error, got: %v", err)
	}
}

func TestParseTasksContractVersion(t *testing.T) {
	content := `---
pios_contract_version: "1.0"
---
# TASKS`

	version, err := parseTasksContractVersion(content)
	if err != nil {
		t.Fatalf("expected valid tasks frontmatter, got error: %v", err)
	}
	if version != "1.0" {
		t.Fatalf("expected version 1.0, got %s", version)
	}
}

func TestParseTasksContractVersionMissing(t *testing.T) {
	content := `# TASKS`

	_, err := parseTasksContractVersion(content)
	if err == nil {
		t.Fatalf("expected error for missing tasks frontmatter")
	}
}

func TestMalformedCheckboxDetectionRegex(t *testing.T) {
	badLines := []string{
		"-[ ] TASK-001: Missing space",
		"###[] TASK-001: Missing space",
	}
	for _, line := range badLines {
		if !checkboxLikeRe.MatchString(line) {
			t.Fatalf("expected checkboxLikeRe to match malformed line: %s", line)
		}
		if pendingTaskRe.MatchString(line) || inProgressTaskRe.MatchString(line) || completedTaskRe.MatchString(line) {
			t.Fatalf("expected malformed line to fail strict checkbox regex: %s", line)
		}
	}
}

func TestSnapshotMilestone(t *testing.T) {
	tempDir := t.TempDir()

	// Create mock templates directory
	templatesDir := filepath.Join(tempDir, "templates")
	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		t.Fatalf("failed to create templates dir: %v", err)
	}

	mockStatus := `---
pios_version: "1.0.0"
current_phase: "v1.0.0"
current_gate: "Plan Lock"
status: "Done"
---
# STATUS`

	mockTasks := `---
pios_contract_version: "1.0"
---
### Phase 1
- [x] Task 1
- [x] Task 2
### Phase 2
- [ ] Task 3
- [/] Task 4`

	if err := os.WriteFile(filepath.Join(tempDir, "STATUS.md"), []byte(mockStatus), 0644); err != nil {
		t.Fatalf("failed to write mock STATUS.md")
	}
	if err := os.WriteFile(filepath.Join(templatesDir, "tasks.md"), []byte(mockTasks), 0644); err != nil {
		t.Fatalf("failed to write mock tasks.md")
	}

	// Run the snapshot
	if err := snapshotMilestone(tempDir); err != nil {
		t.Fatalf("snapshotMilestone failed: %v", err)
	}

	// Verify STATUS.md was reset
	newStatusBytes, _ := os.ReadFile(filepath.Join(tempDir, "STATUS.md"))
	newStatusStr := string(newStatusBytes)
	if !strings.Contains(newStatusStr, `status: "Not Started"`) || !strings.Contains(newStatusStr, `current_phase: "Next Milestone Planning"`) {
		t.Fatalf("STATUS.md was not properly reset. Got: %s", newStatusStr)
	}

	// Verify tasks.md was cleaned of [x] but kept [ ] / [/]
	newTasksBytes, _ := os.ReadFile(filepath.Join(templatesDir, "tasks.md"))
	newTasksStr := string(newTasksBytes)
	if strings.Contains(newTasksStr, "[x] Task 1") || strings.Contains(newTasksStr, "[x] Task 2") {
		t.Fatalf("tasks.md still contains completed tasks!")
	}
	if !strings.Contains(newTasksStr, "[ ] Task 3") || !strings.Contains(newTasksStr, "[/] Task 4") {
		t.Fatalf("tasks.md lost its pending or in-progress tasks!")
	}

	// Verify archive exists
	archiveParent := filepath.Join(templatesDir, "archive")
	entries, err := os.ReadDir(archiveParent)
	if err != nil || len(entries) == 0 {
		t.Fatalf("archive directory was not created: %v", err)
	}
}
