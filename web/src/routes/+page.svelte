<script lang="ts">
  import CodePreview from '$lib/components/CodePreview.svelte';
  import SchemaTabs from '$lib/components/SchemaTabs.svelte';
</script>

<svelte:head>
  <title>Markdown Todo Standard</title>
  <meta name="description" content="The open standard for markdown-based todo lists, empowering users with simplicity and plain-text portability." />
</svelte:head>

<section class="hero container">
  <div class="hero-content">
    <div class="badge">Open Standard v0.1.0-draft</div>
    <h1 class="hero-title">
      Your tasks, tracked in <br />
      <span class="text-gradient">plain text</span>
    </h1>
    <p class="hero-description">
      The Markdown Todo Standard (MDTS) defines Extensible schemas for task management in pure Markdown. From simple lists to complex projects with file attachments.
    </p>
    
    <div class="hero-actions">
      <a href="#schema" class="btn primary">View the Spec</a>
      <a href="https://github.com/markdown-based-todo-standard" class="btn secondary" target="_blank">
        <svg viewBox="0 0 24 24" width="20" height="20" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round" class="css-i6dzq1"><path d="M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77 5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22"></path></svg>
        GitHub
      </a>
    </div>
  </div>
  
  <div class="hero-visual">
    <CodePreview />
  </div>
</section>

<section id="features" class="container section">
  <div class="text-center section-header">
    <h2 class="section-title">Built for Portability</h2>
    <p class="section-description">A standard language for any parser. Parse and display it anywhere.</p>
  </div>
  
  <div class="grid features-grid">
    <div class="card">
      <div class="icon-wrapper">📝</div>
      <h3>Simple Todo Lists</h3>
      <p>Supports robust task identification using Standard GFM with brackets.</p>
    </div>
    
    <div class="card">
      <div class="icon-wrapper">📦</div>
      <h3>Complex Projects</h3>
      <p>Indent content to associate paragraphs and nested lists to their parent task node.</p>
    </div>
    
    <div class="card">
      <div class="icon-wrapper">🔗</div>
      <h3>Attachments</h3>
      <p>Associate local/remote attachments using standard markdown syntax directly inline.</p>
    </div>
  </div>
</section>

<section id="schema" class="container section">
  <div class="text-center section-header">
    <h2 class="section-title">The Specification</h2>
    <p class="section-description">A glimpse into the open syntax rules.</p>
  </div>
  
  <div style="max-width: 800px; margin: 0 auto;">
    <SchemaTabs />
  </div>
</section>

<section id="apps" class="container section mt-ext">
  <div class="banner">
    <div class="banner-content">
      <h2>The Ecosystem</h2>
      <p>We are building official implementations to get you started with this new standard.</p>
      
      <div class="apps-grid">
        <div class="app-card terminal">
          <div class="app-header">
            <h4>TUI</h4>
            <span class="status">In Development</span>
          </div>
          <p>Keyboard-centric Terminal application built in Go.</p>
        </div>
        
        <div class="app-card gui">
           <div class="app-header">
            <h4>GUI Client</h4>
            <span class="status">Planned</span>
          </div>
          <p>Cross-platform application for Desktop and Mobile powered by Tauri & Svelte.</p>
        </div>
      </div>
    </div>
  </div>
</section>

<style>
  .hero {
    display: flex;
    align-items: center;
    gap: 4rem;
    padding: 6rem 2rem 4rem;
    min-height: calc(100vh - 80px); /* subtract header height */
  }

  .hero-content {
    flex: 1;
    max-width: 600px;
  }

  .hero-visual {
    flex: 1;
    display: flex;
    justify-content: center;
    perspective: 1000px;
  }

  .badge {
    display: inline-block;
    padding: 0.4rem 1rem;
    border-radius: 99px;
    background: rgba(234, 74, 170, 0.1);
    color: var(--accent-primary);
    font-size: 0.85rem;
    font-weight: 600;
    margin-bottom: 2rem;
    border: 1px solid rgba(234, 74, 170, 0.2);
  }

  .hero-title {
    font-size: 4.5rem;
    margin-bottom: 1.5rem;
  }

  .hero-description {
    font-size: 1.25rem;
    color: var(--text-secondary);
    margin-bottom: 3rem;
    max-width: 500px;
  }

  .hero-actions {
    display: flex;
    gap: 1rem;
  }

  .btn {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.8rem 1.8rem;
    border-radius: 8px;
    font-weight: 600;
    font-size: 1rem;
    cursor: pointer;
    transition: all 0.2s ease;
  }

  .btn.primary {
    background: var(--text-primary);
    color: var(--bg-color);
  }

  .btn.primary:hover {
    background: #fff;
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(255, 255, 255, 0.1);
  }

  .btn.secondary {
    background: var(--bg-elevated);
    color: var(--text-primary);
    border: 1px solid var(--border-color);
  }

  .btn.secondary:hover {
    background: var(--bg-elevated-hover);
    border-color: var(--text-secondary);
  }

  .section {
    padding: 6rem 2rem;
  }

  .section-header {
    margin-bottom: 4rem;
  }

  .section-title {
    font-size: 2.5rem;
    margin-bottom: 1rem;
  }

  .section-description {
    font-size: 1.1rem;
    color: var(--text-secondary);
  }

  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 2rem;
  }

  .card {
    background: var(--bg-elevated);
    border: 1px solid var(--border-color);
    padding: 2.5rem;
    border-radius: 16px;
    transition: transform 0.3s ease, border-color 0.3s ease;
  }

  .card:hover {
    transform: translateY(-5px);
    border-color: var(--accent-tertiary);
  }

  .icon-wrapper {
    font-size: 2rem;
    margin-bottom: 1.5rem;
    background: rgba(255, 255, 255, 0.05);
    width: 60px;
    height: 60px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 12px;
  }

  .card h3 {
    font-size: 1.25rem;
    margin-bottom: 1rem;
  }

  .card p {
    color: var(--text-secondary);
    font-size: 0.95rem;
  }

  .mt-ext {
    margin-top: 2rem;
  }

  .banner {
    background: linear-gradient(145deg, var(--bg-elevated), rgba(234, 74, 170, 0.05));
    border: 1px solid var(--border-color);
    border-radius: 24px;
    padding: 4rem;
    position: relative;
    overflow: hidden;
  }
  
  .banner::before {
    content: '';
    position: absolute;
    top: 0;
    right: 0;
    width: 400px;
    height: 400px;
    background: radial-gradient(circle, rgba(139, 92, 246, 0.1) 0%, transparent 70%);
    z-index: 0;
  }

  .banner-content {
    position: relative;
    z-index: 1;
  }

  .banner h2 {
    font-size: 2.5rem;
    margin-bottom: 1rem;
  }

  .banner p {
    color: var(--text-secondary);
    font-size: 1.1rem;
    max-width: 600px;
    margin-bottom: 3rem;
  }

  .apps-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
    gap: 1.5rem;
  }

  .app-card {
    background: rgba(13, 17, 23, 0.6);
    border: 1px solid var(--border-color);
    padding: 1.5rem;
    border-radius: 12px;
  }

  .app-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1rem;
  }

  .app-header h4 {
    font-size: 1.1rem;
  }

  .status {
    font-size: 0.75rem;
    padding: 0.2rem 0.6rem;
    border-radius: 99px;
    font-weight: 600;
  }

  .terminal .status {
    background: rgba(35, 134, 54, 0.15);
    color: #3fb950;
    border: 1px solid rgba(35, 134, 54, 0.4);
  }

  .gui .status {
    background: rgba(125, 133, 144, 0.15);
    color: var(--text-secondary);
    border: 1px solid rgba(125, 133, 144, 0.4);
  }

  @media (max-width: 900px) {
    .hero {
      flex-direction: column;
      text-align: center;
      padding-top: 4rem;
    }
    .hero-content {
      margin: 0 auto;
    }
    .hero-actions {
      justify-content: center;
    }
    .banner {
      padding: 2rem;
    }
  }
</style>
