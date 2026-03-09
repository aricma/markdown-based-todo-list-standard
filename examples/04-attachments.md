# 4. Attachments

Attachments are treated as Metadata and are defined using the `attachment:` key followed by a YAML-style list of links or images.

## Remote Links
Any standard Markdown URL within the array.

- [ ] Review design mockups
  attachment:
    - [Figma Mockup](https://figma.com/file/...)
    - [Google Doc Notes](https://docs.google.com/...)

## Local Files
Links resolving to local relative paths.

- [ ] Read the updated guidelines
  Be sure to read chapter 4 specifically.
  
  attachment:
    - [Q3 Guidelines](./docs/q3_guidelines.pdf)

## Media
Images and visual embeds.

- [ ] Verify UI changes
  attachment:
    - ![Screenshot Mobile](./assets/mobile_ui_bug.png)
    - ![Screenshot Desktop](./assets/desktop_ui_bug.png)
