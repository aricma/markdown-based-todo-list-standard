# Markdown Todo Standard (MDTS) Specification

Version: 0.1.0-draft

The Markdown Todo Standard (MDTS) defines a clear, extensible schema for managing tasks, projects, and notes entirely in plain text Markdown. MDTS files are designed to be easily read by humans and robustly parsed by custom parsers.

## 1. Simple Todo Lists

The foundation of MDTS is the standard GitHub Flavored Markdown (GFM) task list feature.

A task item is an unordered list item starting with a bracketed space `[ ]` (incomplete) or bracketed `x` `[x]` (complete).

```markdown
- [ ] Buy groceries
- [x] Feed the cat
* [ ] Read a book
```

**State Definitions:**
- `[ ]` : Incomplete / Todo (Core)
- `[x]` : Complete / Done (Core)
- `[/]` : In Progress / Partially Done (Extended)
- `[-]` : Canceled / Dropped (Extended)
- `[>]` : Deferred / Forwarded (Extended)
- `[?]` : Question / Needs Review (Extended)
- `[!]` : Alert / Notification Trigger (Extended)

*Note: Environments that only support core GFM will render extended states as simple text `[-]` within a list item.*

### Priority

Priority is defined independently of the completion state. You can set the priority of a task by adding exclamation marks `!` immediately following the bracketed state, separated by a space. The number of exclamation marks designates the priority level, up to a maximum of 3.

```markdown
- [ ] ! Buy groceries (Low Priority)
- [ ] !! Feed the cat (Medium Priority)
- [ ] !!! Read a book (High Priority)
```

## 2. Projects and Complex Tasks

Tasks often require more context. MDTS allows tasks to act as containers for paragraphs, sub-tasks, and other block elements.

To associate content with a task, the content must be **indented** to align with the text of the task item (typically 2 or 4 spaces).

### 2.1 Paragraphs

```markdown
- [ ] Write the quarterly report
  The report should cover Q1 and Q2 metrics.
  Make sure to include the new sales figures from the European division.
```

### 2.2 Sub-tasks (Projects)

To create a project or sub-task list, indent another list beneath the parent task:

```markdown
- [ ] Prepare for Launch
  - [x] Create marketing assets
  - [/] Deploy landing page
    Waiting on DNS propagation.
  - [ ] Send newsletter
```

## 3. Attachments

Attachments provide context to tasks. MDTS supports both local file paths and remote URLs as attachments. Attachments must be placed within the task's indented block.

The standard syntax for an attachment is a Markdown link or image, optionally prefixed with a designator if specific rendering is required.

### 3.1 Remote Attachments

A standard URL link.

```markdown
- [ ] Review design mockups
  [Figma Mockup](https://figma.com/file/...)
```

### 3.2 Local Attachments

A link to a local file relative to the Markdown file.

```markdown
- [ ] Read the updated guidelines
  [Q3 Guidelines](./docs/q3_guidelines.pdf)
```

### 3.3 Media Attachments

To embed media directly (images, videos), use the standard markdown image syntax:

```markdown
- [ ] Verify UI changes
  ![Screenshot Mobile](./assets/mobile_ui_bug.png)
```

## 4. Parser Guidelines

MDTS clients **must** implement a custom text parser from scratch rather than relying entirely on generic Markdown ASTs. This ensures strict compliance with the indentation and state-tracking rules unique to MDTS.

**Key Parsing Rules:**
1. **Task Identification**: Any line starting with `\s*[-*+]\s+\[[ x/]\]` is a task node.
2. **Block Continuation**: Any subsequent line with an indentation level greater than the parent task's list marker is considered part of the parent task's body.
3. **Termination**: The task body terminates upon encountering a blank line followed by a line that is indented *less than* the required block continuation margin, or the end of the file.

---
*Draft compiled for initial peer review. Subject to change.*
