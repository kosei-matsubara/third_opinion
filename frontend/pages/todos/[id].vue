<template>
  <div class="edit-todo-page">
    <div class="container">
      <!-- Loading state -->
      <div v-if="isLoading" class="loading-container">
        <div class="spinner"></div>
        <p>Loading todo...</p>
      </div>

      <!-- Error state -->
      <div v-else-if="error" class="error-container">
        <h2>❌ Error</h2>
        <p>{{ error }}</p>
        <button @click="goBack" class="back-btn">
          ← Back to Todo List
        </button>
      </div>

      <!-- Edit form -->
      <div v-else-if="todo">
        <div class="page-header">
          <button @click="goBack" class="back-btn">
            ← Back to Todo List
          </button>
          
          <h1 class="page-title">
            ✏️ Edit Todo
          </h1>
          <p class="page-description">
            Update your todo item.
          </p>
        </div>

        <TodoForm 
          :todo="todo"
          :is-edit="true"
          @success="handleSuccess"
          @cancel="handleCancel"
        />
      </div>

      <!-- Not found state -->
      <div v-else class="not-found-container">
        <h2>📝 Todo Not Found</h2>
        <p>The todo you're looking for doesn't exist.</p>
        <button @click="goBack" class="back-btn">
          ← Back to Todo List
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import type { Todo } from '~/types/todo'

// Route
const route = useRoute()
const todoId = computed(() => parseInt(route.params.id as string))

// Store
const todosStore = useTodosStore()

// State
const todo = ref<Todo | null>(null)
const isLoading = ref(true)
const error = ref<string | null>(null)

// SEO
useHead(() => ({
  title: todo.value ? `Edit: ${todo.value.title} - Todo App` : 'Edit Todo - Todo App',
  meta: [
    { name: 'description', content: 'Edit your todo item.' }
  ]
}))

// Methods
const loadTodo = async () => {
  if (!todoId.value || isNaN(todoId.value)) {
    error.value = 'Invalid todo ID'
    isLoading.value = false
    return
  }

  try {
    isLoading.value = true
    error.value = null
    
    const fetchedTodo = await todosStore.fetchTodoById(todoId.value)
    if (fetchedTodo) {
      todo.value = fetchedTodo
    } else {
      error.value = 'Todo not found'
    }
  } catch (err) {
    console.error('Error loading todo:', err)
    error.value = 'Failed to load todo'
  } finally {
    isLoading.value = false
  }
}

const goBack = () => {
  navigateTo('/')
}

const handleSuccess = () => {
  // Form component handles navigation
}

const handleCancel = () => {
  navigateTo('/')
}

// Initialize
onMounted(() => {
  loadTodo()
})

// Watch for route changes
watch(() => todoId.value, () => {
  loadTodo()
})
</script>

<style scoped>
.edit-todo-page {
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

.loading-container {
  text-align: center;
  padding: 4rem 2rem;
  color: #718096;
}

.spinner {
  width: 40px;
  height: 40px;
  border: 4px solid #e2e8f0;
  border-top: 4px solid #3182ce;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 1rem;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.error-container,
.not-found-container {
  text-align: center;
  padding: 4rem 2rem;
}

.error-container h2,
.not-found-container h2 {
  font-size: 2rem;
  margin-bottom: 1rem;
  color: #2d3748;
}

.error-container p,
.not-found-container p {
  color: #718096;
  margin-bottom: 2rem;
  font-size: 1.125rem;
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
