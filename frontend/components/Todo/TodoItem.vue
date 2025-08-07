<template>
  <div :class="['todo-item', { completed: todo.completed }]">
    <div class="todo-content">
      <!-- Checkbox -->
      <button 
        @click="$emit('toggle', todo)"
        :class="['todo-checkbox', { checked: todo.completed }]"
        :aria-label="todo.completed ? 'Mark as incomplete' : 'Mark as complete'"
      >
        <span v-if="todo.completed" class="checkmark">✓</span>
      </button>

      <!-- Todo info -->
      <div class="todo-info">
        <h3 :class="['todo-title', { completed: todo.completed }]">
          {{ todo.title }}
        </h3>
        <p v-if="todo.description" class="todo-description">
          {{ todo.description }}
        </p>
        <div class="todo-meta">
          <span class="todo-date">
            📅 Created: {{ formatDate(todo.created_at) }}
          </span>
          <span v-if="todo.updated_at !== todo.created_at" class="todo-date">
            ✏️ Updated: {{ formatDate(todo.updated_at) }}
          </span>
        </div>
      </div>
    </div>

    <!-- Actions -->
    <div class="todo-actions">
      <button 
        @click="$emit('edit', todo)"
        class="action-btn edit-btn"
        title="Edit todo"
      >
        ✏️
      </button>
      <button 
        @click="$emit('delete', todo)"
        class="action-btn delete-btn"
        title="Delete todo"
      >
        🗑️
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Todo } from '~/types/todo'

// Props
interface Props {
  todo: Todo
}
defineProps<Props>()

// Emits
defineEmits<{
  toggle: [todo: Todo]
  edit: [todo: Todo]
  delete: [todo: Todo]
}>()

// Methods
const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}
</script>

<style scoped>
.todo-item {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  padding: 1.5rem;
  background: white;
  border: 2px solid #e2e8f0;
  border-radius: 0.75rem;
  transition: all 0.2s ease;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.todo-item:hover {
  border-color: #cbd5e0;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.todo-item.completed {
  background: #f7fafc;
  border-color: #9ae6b4;
}

.todo-content {
  display: flex;
  align-items: flex-start;
  gap: 1rem;
  flex: 1;
}

.todo-checkbox {
  width: 24px;
  height: 24px;
  border: 2px solid #cbd5e0;
  border-radius: 0.375rem;
  background: white;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
  flex-shrink: 0;
  margin-top: 0.125rem;
}

.todo-checkbox:hover {
  border-color: #4299e1;
}

.todo-checkbox.checked {
  background: #48bb78;
  border-color: #48bb78;
  color: white;
}

.checkmark {
  font-size: 14px;
  font-weight: bold;
}

.todo-info {
  flex: 1;
  min-width: 0;
}

.todo-title {
  font-size: 1.125rem;
  font-weight: 600;
  color: #2d3748;
  margin-bottom: 0.5rem;
  line-height: 1.4;
  word-wrap: break-word;
}

.todo-title.completed {
  text-decoration: line-through;
  color: #718096;
}

.todo-description {
  color: #4a5568;
  margin-bottom: 0.75rem;
  line-height: 1.5;
  word-wrap: break-word;
}

.todo-meta {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
  font-size: 0.875rem;
  color: #718096;
}

.todo-date {
  display: flex;
  align-items: center;
  gap: 0.25rem;
}

.todo-actions {
  display: flex;
  gap: 0.5rem;
  margin-left: 1rem;
}

.action-btn {
  width: 36px;
  height: 36px;
  border: none;
  border-radius: 0.5rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  transition: all 0.2s;
  flex-shrink: 0;
}

.edit-btn {
  background: #bee3f8;
  color: #2b6cb0;
}

.edit-btn:hover {
  background: #90cdf4;
}

.delete-btn {
  background: #fed7d7;
  color: #c53030;
}

.delete-btn:hover {
  background: #feb2b2;
}

@media (max-width: 640px) {
  .todo-item {
    padding: 1rem;
    flex-direction: column;
    gap: 1rem;
  }

  .todo-content {
    width: 100%;
  }

  .todo-actions {
    margin-left: 0;
    align-self: flex-end;
  }

  .todo-meta {
    flex-direction: column;
  }
}
</style>
