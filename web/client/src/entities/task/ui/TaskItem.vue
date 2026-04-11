<template>
  <li class="task-item">
    <span>{{ task.task }}</span>
    <small>#{{ task.id }} • user: {{ task.user_id }}</small>
  </li>
</template>

<script setup lang="ts">
import type { Task } from '@/entities/task/model/types';

defineProps<{
  task: Task;
}>();
</script>
