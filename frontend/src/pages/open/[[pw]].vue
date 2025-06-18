<script setup lang="ts">
import api from '@/api';
import HomeBtn from '@/components/HomeBtn.vue';
import { ref } from 'vue';
import { useRoute } from 'vue-router';

const base = import.meta.env.BASE_URL

const route = useRoute('/[[pw]]')
const pw = ref(route.params.pw ?? '')

enum states {
  basic,
  emptyPassword,
  processing,
  emptyFolder,
  error,
  success,
}

let state = states.basic

const files = ref<string[]>([])
await fetchFolder()

async function fetchFolder() {
  try {
    if (pw.value === '') {
      state = states.emptyPassword
      return
    }
    state = states.processing

    const res = await api.foldersFolderGet(pw.value)
    files.value = res.data

    if (files.value.length === 0) {
      state = states.emptyFolder
      return
    }

    state = states.success
  } catch (e) {
    alert('произошла ошибка')
    console.error(e)
    state = states.error
  }
}

function onInput(event: KeyboardEvent) {
  if (event.key !== 'Enter') {
    return
  }

  fetchFolder().catch(e => { throw e })
}

function makeLink(parts: string[]) {
  return parts.map(part => part.replace(/\/+$/, '')).join('/')
}

</script>

<template>
  <HomeBtn />

  <div class="input-group">
    <input type="text" class="form-control" placeholder="Пароль от папки" v-model="pw" @keypress.passive="onInput">
    <button class="btn btn-primary" @click="fetchFolder" :disabled="state === states.processing">Открыть</button>
  </div>

  <p class="text-secondary-emphasis" v-show="state === states.emptyPassword">
    Пароль пуст
  </p>
  <p class="text-secondary-emphasis" v-show="state === states.emptyFolder">
    Папка пуста
  </p>

  <ul class="list-group list-group-flush">
    <li class="list-group-item" v-for="file of files" :key="file">
      <a :href="makeLink([base, 'folders', pw, file])" download>{{ file }}</a>
    </li>
  </ul>
</template>
