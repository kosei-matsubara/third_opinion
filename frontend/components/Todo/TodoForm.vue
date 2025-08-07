<template>
  <form @submit.prevent="handleSubmit" class="todo-form">
    <div class="form-header">
      <h2 class="form-title">
        {{ isEdit ? '✏️ Edit Todo' : '➕ Create New Todo' }}
      </h2>
    </div>

    <!-- Title field -->
    <div class="form-group">
      <label for="title" class="form-label">
        Title <span class="required">*</span>
      </label>
      <input
        id="title"
        v-model="formData.title"
        type="text"
        class="form-input"
        placeholder="Enter todo title..."
        required
        :disabled="isSubmitting"
      />
      <p v-if="errors.title" class="error-text">{{ errors.title }}</p>
    </div>

    <!-- Description field -->
    <div class="form-group">
      <label for="description" class="form-label">
        Description
      </label>
      <textarea
        id="description"
        v-model="formData.description"
        class="form-textarea"
        placeholder="Enter todo description (optional)..."
        rows="4"
        :disabled="isSubmitting"
      ></textarea>
      <p v-if="errors.description" class="error-text">{{ errors.description }}</p>
    </div>

    <!-- Completed checkbox (only for edit mode) -->
    <div v-if="isEdit" class="form-group">
      <label class="checkbox-label">
        <input
          v-model="formData.completed"
          type="checkbox"
          class="form-checkbox"
          :disabled="isSubmitting"
        />
        <span class="checkbox-text">Mark as completed</span>
      </label>
    </div>

    <!-- Error message -->
    <div v-if="submitError" class="error-message">
      <p>❌ {{ submitError }}</p>
    </div>

    <!-- Form actions -->
    <div class="form-actions">
      <button
        type="button"
        @click="handleCancel"
        class="btn btn-secondary"
        :disabled="isSubmitting"
      >
        Cancel
      </button>
      <button
        type="submit"
        class="btn btn-primary"
        :disabled="isSubmitting || !isFormValid"
      >
        <span v-if="isSubmitting" class="loading-spinner"></span>
        {{ isSubmitting ? 'Saving...' : (isEdit ? 'Update Todo' : 'Create Todo') }}
      </button>
    </div>
  </form>
</template>

<script setup lang="ts">
import type { Todo, TodoCreateRequest, TodoUpdateRequest } from '~/types/todo'

// Props
interface Props {
  todo?: Todo
  isEdit?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  isEdit: false
})

// Emits
defineEmits<{
  success: [todo: Todo]
  cancel: []
}>()

// Store
const todosStore = useTodosStore()

// Form data
const formData = reactive({
  title: props.todo?.title || '',
  description: props.todo?.description || '',
  completed: props.todo?.completed || false
})

// Form state
const isSubmitting = ref(false)
const submitError = ref<string | null>(null)
const errors = reactive({
  title: '',
  description: ''
})

// Computed
const isFormValid = computed(() => {
  return formData.title.trim().length > 0 && !errors.title && !errors.description
})

// Methods
const validateForm = () => {
  // Reset errors
  errors.title = ''
  errors.description = ''

  // Validate title
  if (!formData.title.trim()) {
    errors.title = 'Title is required'
  } else if (formData.title.trim().length < 3) {
    errors.title = 'Title must be at least 3 characters long'
  } else if (formData.title.trim().length > 100) {
    errors.title = 'Title must be less than 100 characters'
  }

  // Validate description
  if (formData.description.length > 500) {
    errors.description = 'Description must be less than 500 characters'
  }

  return !errors.title && !errors.description
}

const handleSubmit = async () => {
  if (!validateForm()) {
    return
  }

  isSubmitting.value = true
  submitError.value = null

  try {
    let success = false

    if (props.isEdit && props.todo) {
      // Update existing todo
      const updateData: TodoUpdateRequest = {
        title: formData.title.trim(),
        description: formData.description.trim(),
        completed: formData.completed
      }
      success = await todosStore.updateTodo(props.todo.id, updateData)
    } else {
      // Create new todo
      const createData: TodoCreateRequest = {
        title: formData.title.trim(),
        description: formData.description.trim()
      }
      success = await todosStore.createTodo(createData)
    }

    if (success) {
      // Navigate back to todo list
      await navigateTo('/')
    } else {
      submitError.value = 'Failed to save todo. Please try again.'
    }
  } catch (error) {
    console.error('Form submission error:', error)
    submitError.value = 'An unexpected error occurred. Please try again.'
  } finally {
    isSubmitting.value = false
  }
}

const handleCancel = () => {
  navigateTo('/')
}

// Watch for title changes to clear errors
watch(() => formData.title, () => {
  if (errors.title) {
    errors.title = ''
  }
})

watch(() => formData.description, () => {
  if (errors.description) {
    errors.description = ''
  }
})
</script>

<style scoped>
.todo-form {
  max-width: 600px;
  margin: 0 auto;
  padding: 2rem;
  background: white;
  border-radius: 1rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.form-header {
  margin-bottom: 2rem;
  text-align: center;
}

.form-title {
  font-size: 1.875rem;
  color: #2d3748;
  margin: 0;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-label {
  display: block;
  font-weight: 600;
  color: #374151;
  margin-bottom: 0.5rem;
  font-size: 0.875rem;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.required {
  color: #e53e3e;
}

.form-input,
.form-textarea {
  width: 100%;
  padding: 0.75rem;
  border: 2px solid #e2e8f0;
  border-radius: 0.5rem;
  font-size: 1rem;
  transition: border-color 0.2s;
  font-family: inherit;
}

.form-input:focus,
.form-textarea:focus {
  outline: none;
  border-color: #3182ce;
  box-shadow: 0 0 0 3px rgba(49, 130, 206, 0.1);
}

.form-input:disabled,
.form-textarea:disabled {
  background-color: #f7fafc;
  cursor: not-allowed;
}

.form-textarea {
  resize: vertical;
  min-height: 100px;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  cursor: pointer;
  padding: 0.5rem;
  border-radius: 0.5rem;
  transition: background-color 0.2s;
}

.checkbox-label:hover {
  background-color: #f7fafc;
}

.form-checkbox {
  width: 18px;
  height: 18px;
  cursor: pointer;
}

.checkbox-text {
  font-weight: 500;
  color: #374151;
}

.error-text {
  color: #e53e3e;
  font-size: 0.875rem;
  margin-top: 0.25rem;
}

.error-message {
  padding: 1rem;
  background: #fed7d7;
  border: 1px solid #feb2b2;
  border-radius: 0.5rem;
  margin-bottom: 1.5rem;
  color: #c53030;
}

.form-actions {
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
  margin-top: 2rem;
  padding-top: 1.5rem;
  border-top: 1px solid #e2e8f0;
}

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 0.5rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.875rem;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.btn:disabled {
  cursor: not-allowed;
  opacity: 0.6;
}

.btn-secondary {
  background: #e2e8f0;
  color: #4a5568;
}

.btn-secondary:hover:not(:disabled) {
  background: #cbd5e0;
}

.btn-primary {
  background: #3182ce;
  color: white;
}

.btn-primary:hover:not(:disabled) {
  background: #2c5aa0;
}

.loading-spinner {
  width: 16px;
  height: 16px;
  border: 2px solid transparent;
  border-top: 2px solid currentColor;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

@media (max-width: 640px) {
  .todo-form {
    padding: 1.5rem;
    margin: 1rem;
  }
  
  .form-actions {
    flex-direction: column;
  }
  
  .btn {
    justify-content: center;
  }
}
</style>
