<script setup lang="ts">
import api from '@/api';
import { useRoute } from 'vue-router';

const route = useRoute('/list/[pw]')
const pw = route.params.pw

const folder = await api.foldersFolderGet(pw)
const files = folder.data

function makeFileLink(filename: string) {
  return `${import.meta.env.BASE_URL}/folders/${pw}/${filename}`
}
</script>

<template>
  <main class="island">
    <span v-if="files.length === 0">no files</span>
    <div v-for="file of files" :key="file"><a :href="makeFileLink(file)">{{ file }}</a></div>
  </main>
</template>

<style scoped></style>
