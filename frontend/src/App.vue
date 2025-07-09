<script setup lang="ts">
import { useDark } from '@vueuse/core';
import { RouterView } from 'vue-router'

useDark({
  selector: 'html',
  attribute: 'data-bs-theme',
  valueDark: 'dark',
  valueLight: 'light',
})
</script>

<template>
  <main class="island">
    <div class="grid-vstack earth">
      <Suspense>
        <RouterView />
      </Suspense>
    </div>
  </main>
</template>

<style scoped>
.earth {
  width: 27rem;
  min-height: 20rem;
  align-content: start;
}
</style>
