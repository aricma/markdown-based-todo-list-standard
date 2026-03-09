package parser

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

var (
	// Matches: `- [ ] Title`, `* [x] Title`, etc.
	// Groups: 1=indent, 2=marker, 3=state, 4=priority_spaces, 5=priority_marks, 6=title
	taskRegex = regexp.MustCompile(`^(\s*)([-*+])\s+\[([ xX/\-?>!])\](\s*)(!{0,3})\s*(.*)$`)

	// Matches: `  key: value`
	metadataRegex = regexp.MustCompile(`^(\s+)([a-zA-Z0-9_-]+):\s*(.+)$`)
)

// Parse reads an MDTS document and returns the root Document node
func Parse(r io.Reader) (*Document, error) {
	scanner := bufio.NewScanner(r)
	doc := &Document{Nodes: []Node{}}

	var currentTask *Task
	var expectedMetaIndent int

	for scanner.Scan() {
		line := scanner.Text()

		// 1. Try matching a Task item
		if matches := taskRegex.FindStringSubmatch(line); matches != nil {
			indentStr := matches[1]
			stateStr := matches[3]
			priorityMarks := matches[5]
			title := matches[6]

			// Normalize case-insensitive complete
			if stateStr == "X" {
				stateStr = "x"
			}

			task := &Task{
				State:       State(stateStr),
				Priority:    len(priorityMarks),
				Title:       strings.TrimSpace(title),
				Metadata:    make(map[string]string),
				IndentLevel: len(indentStr),
			}

			// In a real robust parser, we'd use a stack to handle children/sub-lists based on IndentLevel.
			// For v0.1.0 simplicity in this draft, we append all top-level tasks to the document root
			// and treat nested lines as children of the LAST top-level task encountered.
			if currentTask == nil || len(indentStr) <= currentTask.IndentLevel {
				doc.Nodes = append(doc.Nodes, task)
				currentTask = task
				expectedMetaIndent = len(indentStr) + 2 // Typical GFM block indent
			} else {
				currentTask.Children = append(currentTask.Children, task)
			}
			continue
		}

		// If no active task context, skip (it's just regular markdown outside a list)
		if currentTask == nil {
			continue
		}

		// 2. Try matching Metadata Block
		if matches := metadataRegex.FindStringSubmatch(line); matches != nil {
			indentStr := matches[1]
			key := matches[2]
			val := matches[3]

			// Metadata must be indented greater than the parent task marker
			if len(indentStr) >= expectedMetaIndent {
				currentTask.Metadata[key] = strings.TrimSpace(val)
				continue
			}
		}

		// 3. Fallback: Parse Paragraphs / Text block context
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			// Blank line closes context natively unless part of a multiline block
			continue
		}

		// Calculate leading whitespace to determine if it's a child block
		leadingSpaces := len(line) - len(strings.TrimLeft(line, " "))
		if leadingSpaces > currentTask.IndentLevel {
			// Basic text associated with a task
			currentTask.Children = append(currentTask.Children, &Paragraph{Content: trimmedLine})
		} else {
			// We returned to a lower indentation level that isn't a task—reset context
			currentTask = nil
		}
	}

	return doc, scanner.Err()
}
