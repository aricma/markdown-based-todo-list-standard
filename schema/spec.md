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

To maximize adoption for simple plain-text users while supporting complex App migrations, MDTS defines two compliance levels:

**Level 1: Core (Strictly Plain-Text)**
- `[ ]` : Incomplete / Todo
- `[x]` or `[X]` : Complete / Done (Case-insensitive)

**Level 2: Extended (App-Specific Mappings)**
For developers building parsers, bridging proprietary data, or supporting complex project management workflows:
- `[/]` : In Progress / Partially Done
- `[?]` : On Hold / Needs Review / Question
- `[!]` : Canceled / Dropped

### Priority

Priority is defined independently of the completion state. You can set the priority of a task by adding exclamation marks `!` immediately following the bracketed state, separated by a space. The number of exclamation marks designates the priority level, up to a maximum of 3.

```markdown
- [ ] ! Buy groceries (Low Priority)
- [ ] !! Feed the cat (Medium Priority)
- [ ] !!! Read a book (High Priority)
```



## 2. Descriptions and Notes

Tasks can contain block elements like paragraphs. To associate content with a task, the content must be **indented** to align with the text of the task item (typically 2 or 4 spaces).

The indented block of text immediately following a task is considered its **Description**, which can span multiple lines and paragraphs until a new task or an indentation break occurs. If metadata is included, the description block must precede the metadata.

```markdown
- [ ] Write the quarterly report
  The report should cover Q1 and Q2 metrics. Make sure to include the new sales figures from the European division.
```

## 3. Task Metadata

To supercharge tasks for power users, MDTS supports arbitrary key-value metadata. Metadata must follow a strict `key: value` format.

**Rules for Metadata:**
1. It must be placed immediately following the task definition, OR immediately following the task's Description if one exists.
2. It must be indented exactly to the level of the task's text content.
3. Multiple metadata key-value pairs can be stacked sequentially.
4. Any metadata is allowed as long as it can be stringified and is readable in plain markdown (e.g. `deadline: 2026-10-31`, `priority: high`, `assignees: ["@aricma", "@john"]`).
5. If a blank line or a non-metadata block is encountered, the metadata block is considered closed.

```markdown
- [ ] Prepare investor presentation
  deadline: 2026-10-31
  reminder: every tue @ 4pm
  scheduled: 0 9 * * 1
  location: Conference Room A
  assignee: @aricma
  
  Please make sure all the Q3 charts are finalized before the meeting.
```




## 4. Projects and Complex Tasks

Tasks often require more context. MDTS allows tasks to act as containers for paragraphs, sub-tasks, and other block elements.

To associate content with a task, the content must be **indented** to align with the text of the task item (typically 2 or 4 spaces).

### 4.1 Sub-tasks (Projects)

To create a project or sub-task list, indent another list beneath the parent task:

```markdown
- [ ] Prepare for Launch
  - [x] Create marketing assets
  - [/] Deploy landing page
    Waiting on DNS propagation.
  - [ ] Send newsletter
```




## 5. Attachments

Attachments provide context to tasks. MDTS supports both local file paths and remote URLs as attachments. Attachments must be placed within the task's indented block.

The standard syntax for an attachment is a Markdown link or image, optionally prefixed with a designator if specific rendering is required.

### 5.1 Remote Attachments

A standard URL link.

```markdown
- [ ] Review design mockups
  [Figma Mockup](https://figma.com/file/...)
```

### 5.2 Local Attachments

A link to a local file relative to the Markdown file.

```markdown
- [ ] Read the updated guidelines
  [Q3 Guidelines](./docs/q3_guidelines.pdf)
```

### 5.3 Media Attachments

To embed media directly (images, videos), use the standard markdown image syntax:

```markdown
- [ ] Verify UI changes
  ![Screenshot Mobile](./assets/mobile_ui_bug.png)
```



## 6. Parser Guidelines

MDTS clients **must** implement a custom text parser from scratch rather than relying entirely on generic Markdown ASTs. This ensures strict compliance with the indentation and state-tracking rules unique to MDTS.

**Key Parsing Rules:**
1. **Task Identification**: Any line starting with `\s*[-*+]\s+\[[ x/]\]` is a task node.
2. **Block Continuation**: Any subsequent line with an indentation level greater than the parent task's list marker is considered part of the parent task's body.
3. **Termination**: The task body terminates upon encountering a blank line followed by a line that is indented *less than* the required block continuation margin, or the end of the file.

---
*Draft compiled for initial peer review. Subject to change.*


