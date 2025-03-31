<script setup lang="ts">
import api from '@/api';

const { data: folder } = await api.foldersCreatePost()
</script>

<template>
  <main class="island">
    <h3>{{ folder.password }}</h3>
    <input type="date" :value="folder.creation_date" disabled>
    <code>{{ folder.ip }}</code>
  </main>
</template>
