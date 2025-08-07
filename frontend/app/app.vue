<template>
  <div id="app">
    <NuxtRouteAnnouncer />
    
    <header class="app-header">
      <div class="container">
        <NuxtLink to="/" class="logo">
          📝 Todo App
        </NuxtLink>
        <nav class="nav">
          <NuxtLink to="/" class="nav-link">
            📋 All Todos
          </NuxtLink>
          <NuxtLink to="/todos/new" class="nav-link">
            ➕ New Todo
          </NuxtLink>
        </nav>
      </div>
    </header>

    <main class="main-content">
      <NuxtPage />
    </main>

    <footer class="app-footer">
      <div class="container">
        <p>© 2024 Todo App - Built with Nuxt.js & Go</p>
      </div>
    </footer>
  </div>
</template>

<script setup>
// Global error handling
useHead({
  title: 'Todo App',
  meta: [
    { name: 'description', content: 'A simple todo management app built with Nuxt.js and Go' },
    { name: 'viewport', content: 'width=device-width, initial-scale=1' }
  ]
})
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  line-height: 1.6;
  color: #2d3748;
  background-color: #f7fafc;
}

#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

.app-header {
  background: white;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  position: sticky;
  top: 0;
  z-index: 100;
}

.app-header .container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1rem;
}

.logo {
  font-size: 1.5rem;
  font-weight: bold;
  color: #3182ce;
  text-decoration: none;
  transition: color 0.2s;
}

.logo:hover {
  color: #2c5aa0;
}

.nav {
  display: flex;
  gap: 1rem;
}

.nav-link {
  padding: 0.5rem 1rem;
  color: #4a5568;
  text-decoration: none;
  border-radius: 0.5rem;
  transition: all 0.2s;
  font-weight: 500;
}

.nav-link:hover {
  background-color: #edf2f7;
  color: #2d3748;
}

.nav-link.router-link-active {
  background-color: #3182ce;
  color: white;
}

.main-content {
  flex: 1;
  padding: 2rem 0;
}

.app-footer {
  background: white;
  border-top: 1px solid #e2e8f0;
  padding: 1rem 0;
  margin-top: auto;
}

.app-footer p {
  text-align: center;
  color: #718096;
  font-size: 0.875rem;
}

@media (max-width: 640px) {
  .app-header .container {
    flex-direction: column;
    gap: 1rem;
  }
  
  .nav {
    width: 100%;
    justify-content: center;
  }
}
</style>
