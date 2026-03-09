# 2. Descriptions and Metadata

This file demonstrates how to associate descriptions and arbitrary key-value metadata with a task.

## Descriptions
Any indented block of text immediately following a task is its description.

- [ ] Review quarterly report
  The report should cover Q1 and Q2 metrics.
  Make sure to include the new sales figures from the European division.
  
  The marketing team will need these numbers by Thursday.

- [x] Call the client
  Client wanted to discuss the new pricing model.

## Metadata
Metadata must follow a `key: value` format and be indented to the task level. If a description exists, the metadata must follow it.

- [ ] Prepare investor presentation
  Please make sure all the Q3 charts are finalized before the meeting.
  
  deadline: 2026-10-31
  scheduled: * * * * 1
  reminder: every tue @ 4pm
  assignee: @aricma

- [ ] Book flights for conference
  destination: London
  priority: high
  budget: $1200
