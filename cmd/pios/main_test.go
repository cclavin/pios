package main

import (
	"testing"
)

func TestExtractFrontmatter(t *testing.T) {
	content := `---
foo: bar
---
# Content
`
	fm := extractFrontmatter(content)
	if fm != "foo: bar" {
		t.Errorf("Expected 'foo: bar', got '%s'", fm)
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
