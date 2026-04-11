import { http } from '@/shared/api/http';
import type { CreateTaskPayload, Task } from '@/entities/task/model/types';

const mockTasks: Task[] = [
  { id: 1, task: 'Сверстать экран задач', user_id: 1001, is_done: false },
  { id: 2, task: 'Подключить API', user_id: 1001, is_done: true },
];

const useMock = import.meta.env.VITE_USE_MOCK === 'true';

export const taskApi = {
  list: async (): Promise<Task[]> => {
    if (useMock) {
      return mockTasks;
    }

    try {
      const { data } = await http.get<Task[]>('/tasks');
      return data;
    } catch {
      return mockTasks;
    }
  },
  create: async (payload: CreateTaskPayload): Promise<Task> => {
    if (useMock) {
      return {
        id: Date.now(),
        task: payload.task,
        user_id: payload.user_id,
        is_done: payload.is_done ?? false,
      };
    }

    const { data } = await http.post<Task>('/tasks', payload);
    return data;
  },
};
