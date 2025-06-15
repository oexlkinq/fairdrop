<script setup lang="ts">
import { useDark } from '@vueuse/core';
import { RouterLink, RouterView } from 'vue-router'

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
      <uv class="nav nav-pills nav-justified mb-4">
        <li class="nav-item">
          <RouterLink class="nav-link" active-class="active" to="/upload">Создать</RouterLink>
        </li>
        <li class="nav-item">
          <RouterLink class="nav-link" active-class="active" to="/open">Открыть</RouterLink>
        </li>
      </uv>

      <Suspense>
        <RouterView />
      </Suspense>
    </div>
  </main>
</template>

<style scoped>
.earth {
  width: 25rem;
  min-height: 15rem;
  align-content: start;
}
</style>
