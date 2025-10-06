<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div id="app">
    <header class="app-header">
      <h1 class="app-title">Definitely Not Wordle</h1>
      <nav class="app-nav">
        <RouterLink to="/" class="nav-link">Home</RouterLink>
      </nav>
    </header>

    <main class="app-main">
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
#app {
  min-height: 100vh;
  background-color: #f8f9fa;
}

.app-header {
  background-color: #fff;
  border-bottom: 1px solid #e9ecef;
  padding: 1rem 0;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.app-title {
  text-align: center;
  margin: 0 0 1rem 0;
  font-size: 2.5rem;
  font-weight: bold;
  color: #212529;
  letter-spacing: 0.1em;
}

.app-nav {
  display: flex;
  justify-content: center;
  gap: 2rem;
}

.nav-link {
  text-decoration: none;
  color: #6c757d;
  font-weight: 500;
  padding: 0.5rem 1rem;
  border-radius: 0.375rem;
  transition: all 0.2s ease;
}

.nav-link:hover {
  color: #212529;
  background-color: #f8f9fa;
}

.nav-link.router-link-active {
  color: #212529;
  background-color: #e9ecef;
}

.app-main {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem 1rem;
}

@media (min-width: 768px) {
  .app-header {
    padding: 1.5rem 0;
  }
  
  .app-title {
    font-size: 3rem;
    margin-bottom: 1.5rem;
  }
  
  .app-nav {
    gap: 3rem;
  }
}
</style>
