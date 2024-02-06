export type Task = {
  id: string // UUID
  userID: string // UUID
  title: string
  description?: string
  priority: number
  completedAt?: string
  createdAt: string
  updatedAt: string
}

export type User = {
  id: string; // UUID
  email: string; //unique
  avatar?: string;
  username: string;
  // password: string;
  bio?: string;
  created_at: Date;
  updated_at?: Date;
}

export type ErrorResponse = {
  type: string,
  message: string
}

export type DefaultActionResponse<T = Record<string, unknown>> = {
  data?: T,
  error?: ErrorResponse
}