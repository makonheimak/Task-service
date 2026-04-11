<template>
  <section>
    <h1>Tasks</h1>
    <TaskList :tasks="tasks" />
  </section>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import TaskList from '@/widgets/task-list/ui/TaskList.vue';
import { taskApi } from '@/entities/task/api/taskApi';
import type { Task } from '@/entities/task/model/types';

const tasks = ref<Task[]>([]);

onMounted(async () => {
  tasks.value = await taskApi.list();
});
</script>
