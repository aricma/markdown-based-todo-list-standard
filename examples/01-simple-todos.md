# 1. Simple Todo Lists

This file demonstrates the Level 1 (Core) and Level 2 (Extended) syntax for simple todo lists in the Markdown Todo Standard (MDTS).

## Core States (Level 1)
These are standard GitHub Flavored Markdown (GFM) tasks.

- [ ] A pending task
- [x] A completed task
- [X] A completed task (case-insensitive)
* [ ] A pending task using an asterisk marker
+ [ ] A pending task using a plus marker

## Extended States (Level 2)
These states are for power apps to map custom workflows. Any compliant MDTS parser should recognize these states.

- [/] An in-progress task
- [?] A task on hold or needing review
- [!] A canceled task

## Priority
Tasks can have an independent priority using exclamation marks separated by a space.

- [ ] ! Low priority task
- [x] !! Medium priority completed task
- [/] !!! High priority in-progress task
