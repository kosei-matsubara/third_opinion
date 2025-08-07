<template>
  <div class="todo-list">
    <div class="todo-list-header">
      <h2 class="todo-list-title">
        📝 Todo List
        <span class="todo-count">({{ todoCount }})</span>
      </h2>
      
      <!-- Filter tabs -->
      <div class="filter-tabs">
        <button 
          @click="currentFilter = 'all'"
          :class="['filter-tab', { active: currentFilter === 'all' }]"
        >
          All ({{ todoCount }})
        </button>
        <button 
          @click="currentFilter = 'pending'"
          :class="['filter-tab', { active: currentFilter === 'pending' }]"
        >
          Pending ({{ pendingCount }})
        </button>
        <button 
          @click="currentFilter = 'completed'"
          :class="['filter-tab', { active: currentFilter === 'completed' }]"
        >
          Completed ({{ completedCount }})
        </button>
      </div>
    </div>

    <!-- Loading state -->
    <div v-if="isLoading" class="loading">
      <div class="spinner"></div>
      <p>Loading todos...</p>
    </div>

    <!-- Error state -->
    <div v-else-if="error" class="error-message">
      <p>❌ {{ error }}</p>
      <button @click="retryFetch" class="retry-btn">
        🔄 Retry
      </button>
    </div>

    <!-- Empty state -->
    <div v-else-if="filteredTodos.length === 0" class="empty-state">
      <div class="empty-icon">📝</div>
      <h3>{{ getEmptyMessage() }}</h3>
      <p>{{ getEmptyDescription() }}</p>
      <NuxtLink to="/todos/new" class="create-first-btn">
        ➕ Create your first todo
      </NuxtLink>
    </div>

    <!-- Todo items -->
    <div v-else class="todo-items">
      <TodoItem 
        v-for="todo in filteredTodos" 
        :key="todo.id"
        :todo="todo"
        @toggle="handleToggle"
        @delete="handleDelete"
        @edit="handleEdit"
      />
    </div>

    <!-- Add new todo button -->
    <div v-if="!isLoading && !error" class="add-todo-section">
      <NuxtLink to="/todos/new" class="add-todo-btn">
        ➕ Add New Todo
      </NuxtLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Todo } from '~/types/todo'

// Store
const todosStore = useTodosStore()
const { 
  todos, 
  isLoading, 
  error, 
  todoCount, 
  pendingCount, 
  completedCount,
  pendingTodos,
  completedTodos 
} = storeToRefs(todosStore)

// Local state
const currentFilter = ref<'all' | 'pending' | 'completed'>('all')

// Computed
const filteredTodos = computed(() => {
  switch (currentFilter.value) {
    case 'pending':
      return pendingTodos.value
    case 'completed':
      return completedTodos.value
    default:
      return todos.value
  }
})

// Methods
const retryFetch = () => {
  todosStore.clearError()
  todosStore.fetchTodos()
}

const handleToggle = async (todo: Todo) => {
  await todosStore.toggleComplete(todo.id)
}

const handleDelete = async (todo: Todo) => {
  if (confirm(`Are you sure you want to delete "${todo.title}"?`)) {
    await todosStore.deleteTodo(todo.id)
  }
}

const handleEdit = (todo: Todo) => {
  navigateTo(`/todos/${todo.id}`)
}

const getEmptyMessage = () => {
  switch (currentFilter.value) {
    case 'pending':
      return 'No pending todos! 🎉'
    case 'completed':
      return 'No completed todos yet'
    default:
      return 'No todos yet'
  }
}

const getEmptyDescription = () => {
  switch (currentFilter.value) {
    case 'pending':
      return 'All your todos are completed. Great job!'
    case 'completed':
      return 'Complete some todos to see them here.'
    default:
      return 'Start by creating your first todo item.'
  }
}

// Initialize
onMounted(() => {
  todosStore.fetchTodos()
})
</script>

<style scoped>
.todo-list {
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem;
}

.todo-list-header {
  margin-bottom: 2rem;
}

.todo-list-title {
  font-size: 2rem;
  color: #2d3748;
  margin-bottom: 1rem;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.todo-count {
  font-size: 1rem;
  color: #718096;
  font-weight: normal;
}

.filter-tabs {
  display: flex;
  gap: 0.5rem;
  margin-bottom: 1rem;
}

.filter-tab {
  padding: 0.5rem 1rem;
  border: 2px solid #e2e8f0;
  background: white;
  color: #4a5568;
  border-radius: 0.5rem;
  cursor: pointer;
  transition: all 0.2s;
  font-weight: 500;
}

.filter-tab:hover {
  border-color: #3182ce;
  color: #3182ce;
}

.filter-tab.active {
  background: #3182ce;
  color: white;
  border-color: #3182ce;
}

.loading {
  text-align: center;
  padding: 3rem;
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

.error-message {
  text-align: center;
  padding: 2rem;
  background: #fed7d7;
  border: 1px solid #feb2b2;
  border-radius: 0.5rem;
  color: #c53030;
}

.retry-btn {
  margin-top: 1rem;
  padding: 0.5rem 1rem;
  background: #c53030;
  color: white;
  border: none;
  border-radius: 0.25rem;
  cursor: pointer;
}

.empty-state {
  text-align: center;
  padding: 4rem 2rem;
  color: #718096;
}

.empty-icon {
  font-size: 4rem;
  margin-bottom: 1rem;
}

.empty-state h3 {
  font-size: 1.5rem;
  margin-bottom: 0.5rem;
  color: #4a5568;
}

.create-first-btn {
  display: inline-block;
  margin-top: 1.5rem;
  padding: 0.75rem 1.5rem;
  background: #3182ce;
  color: white;
  text-decoration: none;
  border-radius: 0.5rem;
  font-weight: 500;
  transition: background 0.2s;
}

.create-first-btn:hover {
  background: #2c5aa0;
}

.todo-items {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.add-todo-section {
  margin-top: 2rem;
  text-align: center;
}

.add-todo-btn {
  display: inline-block;
  padding: 0.75rem 1.5rem;
  background: #48bb78;
  color: white;
  text-decoration: none;
  border-radius: 0.5rem;
  font-weight: 500;
  transition: background 0.2s;
}

.add-todo-btn:hover {
  background: #38a169;
}

@media (max-width: 640px) {
  .todo-list {
    padding: 1rem;
  }
  
  .todo-list-title {
    font-size: 1.5rem;
  }
  
  .filter-tabs {
    flex-wrap: wrap;
  }
  
  .filter-tab {
    flex: 1;
    min-width: 100px;
  }
}
</style>
