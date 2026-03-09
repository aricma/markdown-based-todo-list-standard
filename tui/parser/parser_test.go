package parser

import (
	"strings"
	"testing"
)

func TestParse_SimpleList(t *testing.T) {
	md := `
- [ ] Incomplete
- [x] Complete
- [X] Case Insensitive
- [/] In Progress
- [?] On Hold
- [!] Canceled
`
	doc, err := Parse(strings.NewReader(strings.TrimSpace(md)))
	if err != nil {
		t.Fatalf("unexpected error parsing simple list: %v", err)
	}
	if len(doc.Nodes) != 6 {
		t.Errorf("expected 6 nodes, got %d", len(doc.Nodes))
	}

	task0 := doc.Nodes[0].(*Task)
	if task0.State != StateTodo || task0.Title != "Incomplete" {
		t.Errorf("failed parsing task0. Got State: %q, Title: %q", task0.State, task0.Title)
	}

	task1 := doc.Nodes[1].(*Task)
	if task1.State != StateDone {
		t.Errorf("failed parsing task1 (Done). Got State: %q", task1.State)
	}

	task2 := doc.Nodes[2].(*Task)
	if task2.State != StateDone {
		t.Errorf("failed parsing task2 (Done case insensitive). Got State: %q", task2.State)
	}

	task4 := doc.Nodes[4].(*Task)
	if task4.State != StateOnHold {
		t.Errorf("failed parsing task4 (OnHold). Got State: %q", task4.State)
	}
}

func TestParse_PriorityAndMetadata(t *testing.T) {
	md := `
- [ ] !!! Urgent Launch Task
  deadline: 2026-10-31
  assignee: @aricma
`
	doc, err := Parse(strings.NewReader(strings.TrimSpace(md)))
	if err != nil {
		t.Fatalf("unexpected error parsing priority list: %v", err)
	}

	task := doc.Nodes[0].(*Task)
	if task.Priority != 3 {
		t.Errorf("expected priority 3, got %d", task.Priority)
	}
	if task.Title != "Urgent Launch Task" {
		t.Errorf("expected title 'Urgent Launch Task', got %q", task.Title)
	}

	if val, ok := task.Metadata["deadline"]; !ok || val != "2026-10-31" {
		t.Errorf("metadata map missing deadline or wrong value: %v", task.Metadata)
	}
	if val, ok := task.Metadata["assignee"]; !ok || val != "@aricma" {
		t.Errorf("metadata map missing assignee or wrong value: %v", task.Metadata)
	}
}
