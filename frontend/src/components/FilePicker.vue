<script setup lang="ts">
const emit = defineEmits<{
  newFiles: [files: File[]],
}>()

function handleSelectedFiles(e: Event) {
  if (!('files' in e && e.files instanceof FileList)) {
    return
  }

  handleFiles(e.files)
}

function handleDroppedFiles(e: DragEvent) {
  if (!e.dataTransfer) {
    return
  }

  handleFiles(e.dataTransfer.files)
}

function handleFiles(files: FileList) {
  emit('newFiles', Array.from(files))
}

</script>

<template>
  <input type="file" id="filePicker" hidden @change="handleSelectedFiles($event)">
  <label for="filePicker" class="btn btn-outline-primary w-100 rounded-5 p-5" @drop.prevent="handleDroppedFiles($event)"
    @dragover.prevent>
    Добавить файлы
  </label>
</template>
