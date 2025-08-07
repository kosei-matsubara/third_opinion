export interface Todo {
  id: number
  title: string
  description: string
  completed: boolean
  created_at: string
  updated_at: string
}

export interface TodoCreateRequest {
  title: string
  description: string
}

export interface TodoUpdateRequest {
  title?: string
  description?: string
  completed?: boolean
}

export interface ApiResponse<T> {
  status: string
  data: T
}

export interface ApiErrorResponse {
  error: string
  message: string
} 