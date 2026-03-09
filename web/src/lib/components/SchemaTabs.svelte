<script lang="ts">
  let openTab = $state('simple');
</script>

<div class="schema-tabs">
  <div class="tabs-header">
    <button class:active={openTab === 'simple'} onclick={() => openTab = 'simple'}>Simple Lists</button>
    <button class:active={openTab === 'projects'} onclick={() => openTab = 'projects'}>Projects</button>
    <button class:active={openTab === 'states'} onclick={() => openTab = 'states'}>Extended States</button>
    <button class:active={openTab === 'priority'} onclick={() => openTab = 'priority'}>Priority</button>
    <button class:active={openTab === 'metadata'} onclick={() => openTab = 'metadata'}>Metadata</button>
    <button class:active={openTab === 'attachments'} onclick={() => openTab = 'attachments'}>Attachments</button>
  </div>
  
  <div class="tabs-content">
    {#if openTab === 'simple'}
      <div class="code-block">
<pre><code>- [ ] Buy groceries
- [x] Feed the cat
* [ ] Read a book</code></pre>
      </div>
      <p class="tab-desc">The foundation of MDTS is the standard GitHub Flavored Markdown (GFM) task list feature.</p>
    {/if}

    {#if openTab === 'projects'}
      <div class="code-block">
<pre><code>- [ ] Prepare for Launch
  - [x] Create marketing assets
  - [/] Deploy landing page
    Waiting on DNS propagation.
  - [ ] Send newsletter</code></pre>
      </div>
      <p class="tab-desc">Tasks can act as containers. Indent content to align with the text of the task item to nest paragraphs and sub-tasks.</p>
    {/if}

     {#if openTab === 'states'}
      <div class="code-block">
<pre><code>- [/] In Progress / Partially Done
- [?] On Hold / Needs Review / Question
- [!] Canceled / Dropped</code></pre>
      </div>
      <p class="tab-desc">Additional state brackets for complex workflows. Standard parsers degrade gracefully to simple text.</p>
    {/if}

    {#if openTab === 'priority'}
      <div class="code-block">
<pre><code>- [ ] ! Buy groceries (Low Priority)
- [ ] !! Feed the cat (Medium Priority)
- [ ] !!! Read a book (High Priority)</code></pre>
      </div>
      <p class="tab-desc">Priority is defined independently of the completion state using exclamation marks.</p>
    {/if}

    {#if openTab === 'metadata'}
      <div class="code-block">
<pre><code>- [ ] Prepare investor presentation
  deadline: 2026-10-31
  scheduled: * * * * 1
  reminder: every tue @ 4pm
  assignee: @aricma</code></pre>
      </div>
      <p class="tab-desc">Supercharge tasks for power users. Support arbitrary key-value metadata indented immediately below a task.</p>
    {/if}

    {#if openTab === 'attachments'}
      <div class="code-block">
<pre><code>- [ ] Review design mockups
  [Figma Mockup](https://figma.com/file/...)
- [ ] Verify UI changes
  ![Screenshot Mobile](./assets/mobile_ui_bug.png)</code></pre>
      </div>
      <p class="tab-desc">Associate context with tasks. MDTS supports both remote URLs and local file attachments within the indented block.</p>
    {/if}
  </div>
</div>

<style>
  .schema-tabs {
    background: var(--bg-elevated);
    border: 1px solid var(--border-color);
    border-radius: 12px;
    overflow: hidden;
  }

  .tabs-header {
    display: flex;
    border-bottom: 1px solid var(--border-color);
    background: rgba(0, 0, 0, 0.2);
  }

  .tabs-header button {
    flex: 1;
    background: transparent;
    border: none;
    color: var(--text-secondary);
    padding: 1rem;
    font-family: var(--font-body);
    font-size: 0.95rem;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s ease;
    border-bottom: 2px solid transparent;
  }

  .tabs-header button:hover {
    color: var(--text-primary);
    background: rgba(255, 255, 255, 0.02);
  }

  .tabs-header button.active {
    color: var(--accent-primary);
    border-bottom-color: var(--accent-primary);
    background: transparent;
  }

  .tabs-content {
    padding: 2rem;
  }

  .code-block {
    background: #0d1117;
    border: 1px solid var(--border-color);
    border-radius: 8px;
    padding: 1.5rem;
    margin-bottom: 1.5rem;
    overflow-x: auto;
  }

  pre {
    margin: 0;
    font-family: var(--font-mono);
    color: #e6edf3;
    font-size: 0.95rem;
    line-height: 1.6;
  }

  .tab-desc {
    color: var(--text-secondary);
    font-size: 1rem;
    line-height: 1.6;
  }
</style>
