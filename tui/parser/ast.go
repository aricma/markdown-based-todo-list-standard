package parser

type State string

const (
	StateTodo     State = " "
	StateDone     State = "x"
	StateProgress State = "/"
	StateOnHold   State = "?"
	StateCanceled State = "!"
)

// Node represents any element in the MDTS DOM
type Node interface {
	IsNode()
}

// Paragraph represents a block of text associated with a task
type Paragraph struct {
	Content string
}
func (Paragraph) IsNode() {}

// Attachment represents a local or remote link/image
type Attachment struct {
	Label string
	URL   string
}
func (Attachment) IsNode() {}

// Task represents a single MDTS task item
type Task struct {
	State       State             // Complete, Todo, InProgress, etc.
	Priority    int               // 0 (none), 1 (!), 2 (!!), 3 (!!!)
	Title       string            // The inline text of the task
	Metadata    map[string]string // Key/Value pairs from indented metadata blocks
	IndentLevel int               // The indentation level of this task

	Children []Node // Can be *Task (sub-projects), *Paragraph, or *Attachment
}
func (*Task) IsNode() {}

// Document is the root node of an MDTS file
type Document struct {
	Nodes []Node
}
func (*Document) IsNode() {}
