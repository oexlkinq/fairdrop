<script setup lang="ts">
import { computed, ref } from 'vue';
import { RouterLink, useRouter } from 'vue-router';

const router = useRouter()

const pw = ref('')
const fUri = computed(() => import.meta.env.BASE_URL + 'list/' + pw.value)


// function
function openFolderPage() {
  router.push(fUri.value)
}

function testFolder() {

}
</script>

<template>
  <main>
    <RouterLink class="btn btn-primary w-100" to="/upload">создать</RouterLink>

    <div class="input-group">
      <input type="text" class="form-control" placeholder="пароль папки" v-model="pw">
      <RouterLink class="btn btn-outline-primary" :to="fUri">открыть</RouterLink>
    </div>
  </main>
</template>

<style scoped>
main {
  display: grid;
  gap: 1rem;
  place-content: center;
  height: 100%;
}
</style>
