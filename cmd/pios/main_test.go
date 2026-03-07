package main

import (
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
pios_version: "0.4.0"
current_phase: "v0.4.0"
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
pios_version: "0.4.0"
current_phase: "v0.4.0"
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
pios_version: "0.4.0"
current_phase: "v0.4.0"
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
pios_contract_version: "0.4"
---
# TASKS`

	version, err := parseTasksContractVersion(content)
	if err != nil {
		t.Fatalf("expected valid tasks frontmatter, got error: %v", err)
	}
	if version != "0.4" {
		t.Fatalf("expected version 0.4, got %s", version)
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
