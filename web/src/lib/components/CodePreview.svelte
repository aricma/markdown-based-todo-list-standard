<script lang="ts">
  let mockLines = [
    { type: 'item', text: "- [x] Define MDTS Parser", indent: 0 },
    { type: 'item', text: "- [/] Initialize SvelteKit Web Hub", indent: 0, active: true },
    { type: 'item', text: "  Deploy to Vercel/Netlify for CI/CD.", indent: 1, active: true },
    { type: 'item', text: "  - [ ] Add dark mode toggles", indent: 1, active: true },
    { type: 'item', text: "- [ ] ! Write Tauri App Specs", indent: 0 },
    { type: 'item', text: "  [Figma Design mockups](https://...)", indent: 1 }
  ];
</script>

<div class="code-window">
  <div class="window-header">
    <div class="controls">
      <span class="control red"></span>
      <span class="control yellow"></span>
      <span class="control green"></span>
    </div>
    <div class="title">todos.md</div>
  </div>
  
  <div class="code-content">
    <div class="line-numbers">
      {#each mockLines as _, i}
        <div class="number">{i + 1}</div>
      {/each}
    </div>
    
    <div class="code-lines">
      {#each mockLines as line}
        <div class="line" style="padding-left: {line.indent * 2}rem" class:active={line.active}>
          {#if line.text.includes('- [x]')}
            <span class="syntax-list">- </span><span class="syntax-bracket">[</span><span class="syntax-x">x</span><span class="syntax-bracket">]</span>
            <span class="syntax-text strike">{line.text.replace('- [x] ', '')}</span>
          {:else if line.text.includes('- [/]')}
            <span class="syntax-list">- </span><span class="syntax-bracket">[</span><span class="syntax-progress">/</span><span class="syntax-bracket">]</span>
            <span class="syntax-text highlight">{line.text.replace('- [/] ', '')}</span>
          {:else if line.text.includes('- [ ] !')}
            <span class="syntax-list">- </span><span class="syntax-bracket">[ ]</span><span class="syntax-priority"> !</span>
            <span class="syntax-text">{line.text.replace('- [ ] ! ', '')}</span>
          {:else if line.text.includes('- [ ]')}
            <span class="syntax-list">- </span><span class="syntax-bracket">[ ]</span>
            <span class="syntax-text">{line.text.replace('- [ ] ', '')}</span>
          {:else if line.text.includes('[Figma')}
            <span class="syntax-link-bracket">[</span><span class="syntax-link-text">Figma Design mockups</span><span class="syntax-link-bracket">]</span><span class="syntax-url-bracket">(</span><span class="syntax-url">https://...</span><span class="syntax-url-bracket">)</span>
          {:else}
             <span class="syntax-text dim">{line.text}</span>
          {/if}
        </div>
      {/each}
      <div class="cursor"></div>
    </div>
  </div>
</div>

<style>
  .code-window {
    background: #0d1117;
    border: 1px solid var(--border-color);
    border-radius: 12px;
    width: 100%;
    max-width: 500px;
    box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5), 0 0 0 1px rgba(255, 255, 255, 0.05);
    overflow: hidden;
    transform: rotateY(-5deg) rotateX(5deg);
    transition: transform 0.5s ease;
  }

  .code-window:hover {
    transform: rotateY(0deg) rotateX(0deg) translateY(-10px);
  }

  .window-header {
    background: #161b22;
    padding: 0.75rem 1rem;
    border-bottom: 1px solid var(--border-color);
    display: flex;
    align-items: center;
  }

  .controls {
    display: flex;
    gap: 0.5rem;
  }

  .control {
    width: 12px;
    height: 12px;
    border-radius: 50%;
  }

  .red { background: #ff5f56; }
  .yellow { background: #ffbd2e; }
  .green { background: #27c93f; }

  .title {
    flex: 1;
    text-align: center;
    color: var(--text-secondary);
    font-size: 0.85rem;
    font-family: var(--font-mono);
    margin-right: 48px; /* Offset controls width for true center */
  }

  .code-content {
    padding: 1.5rem 0;
    display: flex;
    font-family: var(--font-mono);
    font-size: 0.9rem;
    line-height: 1.7;
    background: #0d1117;
  }

  .line-numbers {
    padding: 0 1rem;
    color: #484f58;
    text-align: right;
    user-select: none;
    border-right: 1px solid var(--border-color);
  }

  .code-lines {
    padding-left: 1rem;
    flex: 1;
    position: relative;
  }

  .line {
    white-space: pre;
    color: #c9d1d9;
  }

  .line.active {
    background: rgba(56, 139, 253, 0.1);
    border-left: 2px solid #58a6ff;
    margin-left: -1rem;
    padding-left: calc(1rem - 2px); /* adjust for border */
  }

  .syntax-list { color: #8b949e; }
  .syntax-bracket { color: #8b949e; }
  .syntax-x { color: #3fb950; }
  .syntax-progress { color: #d29922; }
  .syntax-priority { color: #f85149; font-weight: bold; }
  
  .syntax-text.strike { 
    text-decoration: line-through; 
    color: #8b949e; 
  }
  .syntax-text.highlight {
    color: #e6edf3;
  }
  .syntax-text.dim {
    color: #8b949e;
  }

  .syntax-link-bracket { color: #8b949e; }
  .syntax-link-text { color: #58a6ff; }
  .syntax-url-bracket { color: #8b949e; }
  .syntax-url { color: #a5d6ff; text-decoration: underline; }

  .cursor {
    position: absolute;
    width: 8px;
    height: 1.1rem;
    background: #58a6ff;
    opacity: 0.8;
    top: calc(1.7rem * 2 + 0.3rem); /* rough calc targeting 3rd line end */
    left: 21rem;
    animation: blink 1s step-end infinite;
  }

  @keyframes blink {
    0%, 100% { opacity: 1; }
    50% { opacity: 0; }
  }
</style>
