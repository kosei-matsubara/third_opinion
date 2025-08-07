import { defineStore } from 'pinia'
import type { Todo, TodoCreateRequest, TodoUpdateRequest, ApiResponse } from '~/types/todo'

export const useTodosStore = defineStore('todos', () => {
  // State
  const todos = ref<Todo[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  // Runtime config for API base URL
  const config = useRuntimeConfig()

  // Actions
  const clearError = () => {
    error.value = null
  }

  const fetchTodos = async () => {
    isLoading.value = true
    error.value = null
    
    try {
      const response = await $fetch<ApiResponse<Todo[]>>(`${config.public.apiBase}/api/todos`)
      todos.value = response.data || []
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch todos'
      console.error('Error fetching todos:', err)
    } finally {
      isLoading.value = false
    }
  }

  const fetchTodoById = async (id: number): Promise<Todo | null> => {
    isLoading.value = true
    error.value = null

    try {
      const response = await $fetch<ApiResponse<Todo>>(`${config.public.apiBase}/api/todos/${id}`)
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch todo'
      console.error('Error fetching todo:', err)
      return null
    } finally {
      isLoading.value = false
    }
  }

  const createTodo = async (todoData: TodoCreateRequest): Promise<boolean> => {
    isLoading.value = true
    error.value = null

    try {
      const response = await $fetch<ApiResponse<Todo>>(`${config.public.apiBase}/api/todos`, {
        method: 'POST',
        body: todoData
      })
      
      if (response.data) {
        todos.value.push(response.data)
        return true
      }
      return false
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to create todo'
      console.error('Error creating todo:', err)
      return false
    } finally {
      isLoading.value = false
    }
  }

  const updateTodo = async (id: number, updateData: TodoUpdateRequest): Promise<boolean> => {
    isLoading.value = true
    error.value = null

    try {
      const response = await $fetch<ApiResponse<Todo>>(`${config.public.apiBase}/api/todos/${id}`, {
        method: 'PUT',
        body: updateData
      })

      if (response.data) {
        const index = todos.value.findIndex(todo => todo.id === id)
        if (index !== -1) {
          todos.value[index] = response.data
        }
        return true
      }
      return false
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to update todo'
      console.error('Error updating todo:', err)
      return false
    } finally {
      isLoading.value = false
    }
  }

  const deleteTodo = async (id: number): Promise<boolean> => {
    isLoading.value = true
    error.value = null

    try {
      await $fetch(`${config.public.apiBase}/api/todos/${id}`, {
        method: 'DELETE'
      })

      todos.value = todos.value.filter(todo => todo.id !== id)
      return true
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to delete todo'
      console.error('Error deleting todo:', err)
      return false
    } finally {
      isLoading.value = false
    }
  }

  const toggleComplete = async (id: number): Promise<boolean> => {
    isLoading.value = true
    error.value = null

    try {
      const response = await $fetch<ApiResponse<Todo>>(`${config.public.apiBase}/api/todos/${id}/toggle`, {
        method: 'PATCH'
      })

      if (response.data) {
        const index = todos.value.findIndex(todo => todo.id === id)
        if (index !== -1) {
          todos.value[index] = response.data
        }
        return true
      }
      return false
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to toggle todo'
      console.error('Error toggling todo:', err)
      return false
    } finally {
      isLoading.value = false
    }
  }

  // Getters
  const completedTodos = computed(() => todos.value.filter(todo => todo.completed))
  const pendingTodos = computed(() => todos.value.filter(todo => !todo.completed))
  const todoCount = computed(() => todos.value.length)
  const completedCount = computed(() => completedTodos.value.length)
  const pendingCount = computed(() => pendingTodos.value.length)

  return {
    // State
    todos: readonly(todos),
    isLoading: readonly(isLoading),
    error: readonly(error),
    
    // Actions
    clearError,
    fetchTodos,
    fetchTodoById,
    createTodo,
    updateTodo,
    deleteTodo,
    toggleComplete,
    
    // Getters
    completedTodos,
    pendingTodos,
    todoCount,
    completedCount,
    pendingCount
  }
})
