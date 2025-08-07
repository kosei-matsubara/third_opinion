<template>
  <div class="home-page">
    <div class="container">
      <!-- Page header -->
      <div class="page-header">
        <h1 class="page-title">
          📝 My Todos
        </h1>
        <p class="page-description">
          Manage your tasks efficiently with this simple todo application.
        </p>
      </div>

      <!-- Todo list component -->
      <TodoList />
    </div>
  </div>
</template>

<script setup>
// SEO
useHead({
  title: 'Todo List - Todo App',
  meta: [
    { name: 'description', content: 'View and manage all your todo items in one place.' }
  ]
})
</script>

<style scoped>
.home-page {
  min-height: calc(100vh - 200px);
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

.page-header {
  text-align: center;
  margin-bottom: 3rem;
}

.page-title {
  font-size: 3rem;
  color: #2d3748;
  margin-bottom: 1rem;
  font-weight: 800;
}

.page-description {
  font-size: 1.125rem;
  color: #718096;
  max-width: 600px;
  margin: 0 auto;
}

@media (max-width: 640px) {
  .page-title {
    font-size: 2rem;
  }
  
  .page-description {
    font-size: 1rem;
  }
}
</style>
