<script setup lang="ts">
import api from '@/api';
import type { AxiosProgressEvent } from 'axios';
import { reactive, ref, useTemplateRef } from 'vue';

const maxFormSize = 33554432

const fileInput = useTemplateRef('files')
const uploadState = ref({
  progress: 0,
  processing: false,
})

const { data: folder } = await api.foldersCreatePost()

async function submit() {
  if (uploadState.value.processing) {
    return
  }

  try {
    uploadState.value.processing = true
    uploadState.value.progress = 0

    const files = fileInput.value?.files
    if (!files) {
      return alert('no files')
    }

    const fileArray = new Array<File>(files.length)
    let formSize = 0
    for (let i = 0; i < files.length; i++) {
      const file = files.item(i)
      if (!file) {
        return alert(`null item(${i}) in files`)
      }

      formSize += file.size

      fileArray[i] = file
    }

    if (formSize > maxFormSize) {
      return alert('too big files')
    }

    await api.foldersFolderPost(folder.password, fileArray, {
      onUploadProgress(progressEvent) {
        uploadState.value.progress = progressEvent.progress ?? 0
      },
    })
  } catch (e) {
    alert('error while uploading')
  } finally {
    uploadState.value.processing = false
  }
}
</script>

<template>
  <main class="island">
    <h3>{{ folder.password }}</h3>
    <input type="date" :value="folder.creation_date" disabled>
    <code>{{ folder.ip }}</code>
    <form @submit.prevent="submit">
      <input type="hidden" name="MAX_FILE_SIZE" :value="maxFormSize">
      <input class="form-control" type="file" name="file" ref="files" multiple :disabled="uploadState.processing">

      <input class="btn btn-primary w-100" type="submit" value="Загрузить" :disabled="uploadState.processing">

      <div class="progress" :aria-disabled="!uploadState.processing">
        <div class="progress-bar" :style="{ width: `${uploadState.progress * 100}%` }"></div>
      </div>
    </form>
  </main>
</template>

<style scoped>
form {
  display: grid;
  gap: 1rem;
}
</style>
