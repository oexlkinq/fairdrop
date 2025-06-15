<script setup lang="ts">
import api from '@/api';
import CopyBox from '@/components/CopyBox.vue';
import HomeBtn from '@/components/HomeBtn.vue';
import { ref, useTemplateRef } from 'vue';

const maxFormSize = 33554432

const base = import.meta.env.BASE_URL

const fileInput = useTemplateRef('files')
const uploadState = ref({
  progress: 0,
  processing: false,
  success: false,
  password: '',
  link: '',
})

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

    const { data: folder } = await api.foldersPushPost(fileArray, {
      onUploadProgress(progressEvent) {
        uploadState.value.progress = progressEvent.progress ?? 0
      },
    })

    uploadState.value.password = folder.password
    uploadState.value.link = makeLink(folder.password)
    uploadState.value.success = true
  } catch (e) {
    alert('error while uploading')
  } finally {
    uploadState.value.processing = false
  }
}

function makeLink(password: string) {
  return `${location.origin}${base}open/${password}`
}
</script>

<template>
  <HomeBtn />

  <form class="grid-vstack" @submit.prevent="submit">
    <input type="hidden" name="MAX_FILE_SIZE" :value="maxFormSize">
    <input class="form-control" type="file" name="file" ref="files" multiple :disabled="uploadState.processing">

    <input class="btn btn-primary w-100" type="submit" value="Загрузить" :disabled="uploadState.processing">

    <div class="progress" :aria-disabled="!uploadState.processing">
      <div class="progress-bar" :style="{ width: `${uploadState.progress * 100}%` }"></div>
    </div>

    <div v-show="uploadState.success">
      <p>Файлы успешно загружены</p>
      <p>
        Пароль: <CopyBox>{{ uploadState.password }}</CopyBox><br>
        Ссылка: <CopyBox>{{ uploadState.link }}</CopyBox>
      </p>
    </div>
  </form>
</template>

<route lang="json">{
  "alias": [
    "/upload/"
  ]
}</route>
