<script lang="ts">
  import '../styles/global.css';
  let { children } = $props();
</script>

<div class="app">
  <nav class="navbar">
    <div class="container nav-content">
      <div class="logo">
        <span class="logo-text">MDTS</span>
      </div>
      <div class="nav-links">
        <a href="#features">Features</a>
        <a href="#schema">Schema</a>
        <a href="#apps">Apps</a>
        <a href="https://github.com/markdown-based-todo-standard" target="_blank" class="github-link">
          GitHub
        </a>
      </div>
    </div>
  </nav>

  <main>
    {@render children()}
  </main>

  <footer class="footer">
    <div class="container text-center">
      <p>Markdown Todo Standard — Open Ecosystem</p>
    </div>
  </footer>
</div>

<style>
  .app {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
  }

  main {
    flex: 1;
  }

  .navbar {
    position: sticky;
    top: 0;
    z-index: 100;
    backdrop-filter: blur(12px);
    -webkit-backdrop-filter: blur(12px);
    background: rgba(13, 17, 23, 0.8);
    border-bottom: 1px solid var(--border-color);
    padding: 1rem 0;
  }

  .nav-content {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .logo-text {
    font-family: var(--font-heading);
    font-size: 1.5rem;
    font-weight: 800;
    letter-spacing: -0.05em;
    background: linear-gradient(135deg, var(--text-primary), var(--text-secondary));
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .nav-links {
    display: flex;
    gap: 2rem;
    align-items: center;
  }

  .nav-links a {
    color: var(--text-secondary);
    font-size: 0.9rem;
    font-weight: 500;
  }

  .nav-links a:hover {
    color: var(--text-primary);
    text-decoration: none;
  }

  .github-link {
    padding: 0.4rem 1rem;
    border-radius: 99px;
    background: var(--bg-elevated);
    border: 1px solid var(--border-color);
    transition: all 0.2s ease;
  }

  .github-link:hover {
    background: var(--text-primary);
    color: var(--bg-color) !important;
    border-color: var(--text-primary);
  }

  .footer {
    padding: 3rem 0;
    margin-top: 4rem;
    border-top: 1px solid var(--border-color);
    color: var(--text-secondary);
    font-size: 0.9rem;
  }
</style>
