export interface Task {
  id: number;
  task: string;
  user_id: number;
  is_done: boolean;
  created_at?: string;
  updated_at?: string;
  deleted_at?: string | null;
}

export interface CreateTaskPayload {
  task: string;
  user_id: number;
  is_done?: boolean;
}
