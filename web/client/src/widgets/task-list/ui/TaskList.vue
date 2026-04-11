<template>
  <ul v-if="tasks.length">
    <TaskItem v-for="task in tasks" :key="task.id" :task="task" />
  </ul>
  <p v-else>Список задач пуст.</p>
</template>

<script setup lang="ts">
import TaskItem from '@/entities/task/ui/TaskItem.vue';
import type { Task } from '@/entities/task/model/types';

defineProps<{
  tasks: Task[];
}>();
</script>
