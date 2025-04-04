<script setup lang="ts">
import api from '@/api';
import { ref, useTemplateRef } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const maxFormSize = 33554432

const fileInput = useTemplateRef('files')
const uploadState = ref({
  progress: 0,
  processing: false,
})

const route = useRoute('/upload/[[pw]]')
let pw = route.params.pw ?? ''
if (pw === '') {
  const { data: folder } = await api.foldersCreatePost()

  pw = folder.password
  useRouter().replace('/upload/' + pw)
}

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

    await api.foldersFolderPost(pw, fileArray, {
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
  <h3>{{ pw }}</h3>

  <form class="grid-vstack" @submit.prevent="submit">
    <input type="hidden" name="MAX_FILE_SIZE" :value="maxFormSize">
    <input class="form-control" type="file" name="file" ref="files" multiple :disabled="uploadState.processing">

    <input class="btn btn-primary w-100" type="submit" value="Загрузить" :disabled="uploadState.processing">

    <div class="progress" :aria-disabled="!uploadState.processing">
      <div class="progress-bar" :style="{ width: `${uploadState.progress * 100}%` }"></div>
    </div>
  </form>
</template>

<route lang="json">{
  "alias": [
    "/upload/"
  ]
}</route>
