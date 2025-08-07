<template>
  <div class="new-todo-page">
    <div class="container">
      <!-- Back button -->
      <div class="page-header">
        <button @click="goBack" class="back-btn">
          ← Back to Todo List
        </button>
        
        <h1 class="page-title">
          ➕ Create New Todo
        </h1>
        <p class="page-description">
          Add a new task to your todo list.
        </p>
      </div>

      <!-- Todo form -->
      <TodoForm 
        :is-edit="false"
        @success="handleSuccess"
        @cancel="handleCancel"
      />
    </div>
  </div>
</template>

<script setup>
// SEO
useHead({
  title: 'Create New Todo - Todo App',
  meta: [
    { name: 'description', content: 'Create a new todo item for your task list.' }
  ]
})

// Methods
const goBack = () => {
  navigateTo('/')
}

const handleSuccess = () => {
  // Form component handles navigation
}

const handleCancel = () => {
  navigateTo('/')
}
</script>

<style scoped>
.new-todo-page {
  min-height: calc(100vh - 200px);
}

.container {
  max-width: 800px;
  margin: 0 auto;
  padding: 0 1rem;
}

.page-header {
  text-align: center;
  margin-bottom: 3rem;
}

.back-btn {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  background: #e2e8f0;
  color: #4a5568;
  border: none;
  border-radius: 0.5rem;
  cursor: pointer;
  font-weight: 500;
  transition: background 0.2s;
  margin-bottom: 2rem;
}

.back-btn:hover {
  background: #cbd5e0;
}

.page-title {
  font-size: 2.5rem;
  color: #2d3748;
  margin-bottom: 1rem;
  font-weight: 700;
}

.page-description {
  font-size: 1.125rem;
  color: #718096;
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
